package trivia

import (
	"context"

	"github.com/bsdlp/packagebot/src/outbound/outbound"
	"github.com/bsdlp/packagebot/src/trivia/internal/opentdb"
	"github.com/golang/protobuf/ptypes/empty"
)

func marshalQuestion(question *opentdb.Question) map[string]interface{} {
	hash := map[string]interface{}{}
	hash["Category"] = question.Category
	hash["Type"] = question.Type
	hash["Difficulty"] = question.Difficulty
	hash["Question"] = question.Question
	hash["CorrectAnswer"] = question.CorrectAnswer
	hash["IncorrectAnswers"] = question.IncorrectAnswers
}

func (svc *Service) commitQuestion(ctx context.Context, question *opentdb.Question) error {
	key, err := hashQuestion(question)
	if err != nil {
		return err
	}

	return svc.QuestionsDB.WithContext(ctx).HMSet(key, marshalQuestion(question)).Err()
}

func (svc *Service) shuffleChoices(choices []string) {
	svc.Rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})
}

// GiveQuestion implements Trivia
func (svc *Service) GiveQuestion(ctx context.Context, opt *GiveQuestionInput) (*empty.Empty, error) {
	// retrieve question from api
	question, err := opentdb.GetQuestion(ctx, &opentdb.QuestionParameters{
		Difficulty: opt.GetDifficulty().String(),
	})
	if err != nil {
		return nil, err
	}

	// commit trivia question to redis in order to look up correct answer on
	// reaction
	err = svc.commitQuestion(ctx, question)
	if err != nil {
		return nil, err
	}

	// shuffle correct answer's placement in choices so it's a fair question
	choices = append(question.IncorrectAnswers, question.CorrectAnswer)
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
