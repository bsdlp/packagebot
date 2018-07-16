package main

import (
	"context"

	"github.com/akrylysov/algnhsa"
	"github.com/bsdlp/packagebot/src/trivia/internal/opentdb"
	"github.com/bsdlp/packagebot/src/trivia/internal/trivia"
)

type triviasvc struct{}

// GetQuestion implements trivia.Trivia
func (tc *triviasvc) GetQuestion(ctx context.Context, t *trivia.QuestionType) (*trivia.Question, error) {
	return opentdb.GetQuestion(ctx, t)
}

func main() {
	svc := trivia.NewTriviaServer(&triviasvc{}, nil)
	algnhsa.ListenAndServe(svc, &algnhsa.Options{
		BinaryContentTypes: []string{"application/protobuf"},
		UseProxyPath:       true,
	})
}
