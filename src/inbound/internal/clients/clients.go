package clients

import (
	"net/http"

	"github.com/bsdlp/packagebot/src/trivia/trivia"
)

// Clients contains clients
type Clients struct {
	Trivia trivia.Trivia
}

// New instantiates clients
func New() *Clients {
	httpClient := &http.Client{}
	return &Clients{
		Trivia: trivia.NewTriviaProtobufClient(serviceAddress("trivia"), httpClient),
	}
}
