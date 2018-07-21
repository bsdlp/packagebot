package trivia

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

// AnswerQuestion answers the question
func (svc *Service) AnswerQuestion(ctx context.Context, opt *AnswerQuestionInput) (*empty.Empty, error) {
	// TODO: calculate question hash, look up in dynamodb and compare answers
	return nil, nil
}
