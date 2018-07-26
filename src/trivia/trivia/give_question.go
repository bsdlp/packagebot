package trivia

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/bsdlp/packagebot/src/outbound/outbound"
	"github.com/bsdlp/packagebot/src/trivia/internal/opentdb"
	"github.com/golang/protobuf/ptypes/empty"
)

var errQuestionAlreadyGiven = errors.New("question already given")

// only commits if question does not exist in redis. if it exists, return
// errQuestionAlreadyGiven
func (svc *Service) commitQuestion(ctx context.Context, question *opentdb.Question) error {
	key, err := hashQuestion(question)
	if err != nil {
		return err
	}

	questionJSON, err := json.Marshal(question)
	if err != nil {
		return err
	}

	var expiry = 60 * time.Minute
	set, err := svc.QuestionsDB.WithContext(ctx).SetNX(key, string(questionJSON), expiry).Result()
	if err != nil {
		return err
	}
	if !set {
		return errQuestionAlreadyGiven
	}
	return nil
}

func (svc *Service) shuffleChoices(choices []string) {
	svc.Rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})
}

// GiveQuestion implements Trivia
func (svc *Service) GiveQuestion(ctx context.Context, opt *GiveQuestionInput) (*empty.Empty, error) {
	// retrieve question from api
	question, err := svc.GetQuestion(ctx, &opentdb.QuestionParameters{
		Difficulty: opt.GetDifficulty().String(),
	})
	if err != nil {
		return nil, err
	}

	// commit trivia question to redis in order to look up correct answer on
	// reaction
	err = svc.commitQuestion(ctx, question)
	if err != nil {
		// if question has already been committed to redis, try again
		if err == errQuestionAlreadyGiven {
			return svc.GiveQuestion(ctx, opt)
		}
		return nil, err
	}

	// shuffle correct answer's placement in choices so it's a fair question
	choices := append(question.IncorrectAnswers, question.CorrectAnswer)
	svc.shuffleChoices(choices)

	// send trivia to channel
	params := &outbound.SendChannelTriviaInput{
		ChannelID:        opt.ChannelID,
		RequestingUserID: opt.RequestingUserID,
		Category:         question.Category,
		Difficulty:       question.Difficulty,
		Choices:          choices,
	}
	return svc.Outbound.SendChannelTrivia(ctx, params)
}
