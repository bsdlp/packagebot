package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/bsdlp/packagebot/src/trivia/trivia"
)

func main() {
	svc := trivia.NewTriviaServer(&trivia.Service{
		Rand: trivia.NewCryptoRand(),
	}, nil)
	algnhsa.ListenAndServe(svc, &algnhsa.Options{
		BinaryContentTypes: []string{"application/protobuf"},
		UseProxyPath:       true,
	})
}