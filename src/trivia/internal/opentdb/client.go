package opentdb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type apiResponse struct {
	ResponseCode int        `json:"response_code"`
	Results      []question `json:"results"`
}

type question struct {
	Category         base64EncodedString   `json:"category"`
	Type             base64EncodedString   `json:"type"`
	Difficulty       base64EncodedString   `json:"difficulty"`
	Question         base64EncodedString   `json:"question"`
	CorrectAnswer    base64EncodedString   `json:"correct_answer"`
	IncorrectAnswers []base64EncodedString `json:"incorrect_answers"`
}

type base64EncodedString string

func (s *base64EncodedString) UnmarshalJSON(bs []byte) error {
	var bes string
	err := json.Unmarshal(bs, &bes)
	if err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(bes)
	if err != nil {
		return err
	}

	*s = base64EncodedString(decoded)
	return nil
}

// Question describes a trivia question
type Question struct {
	Category         string
	Type             string
	Difficulty       string
	Question         string
	CorrectAnswer    string
	IncorrectAnswers []string
}

// GetQuestion retrieves a trivia question
func GetQuestion(ctx context.Context) (question Question, err error) {
	resp, err := http.Get("https://opentdb.com/api.php?amount=1&type=multiple&encode=base64")
	if err != nil {
		return
	}

	defer func() {
		closeErr := resp.Body.Close()
		if err == nil {
			err = closeErr
		}
	}()

	var r apiResponse
	err = json.NewDecoder(&io.LimitedReader{
		R: resp.Body,
		N: 1 << 20, // 1 MB
	}).Decode(&r)
	if len(r.Results) == 0 {
		err = errors.New("unexpected empty trivia question response")
		return
	}

	question = Question{
		Category:         string(r.Results[0].Category),
		Type:             string(r.Results[0].Type),
		Difficulty:       string(r.Results[0].Difficulty),
		Question:         string(r.Results[0].Question),
		CorrectAnswer:    string(r.Results[0].CorrectAnswer),
		IncorrectAnswers: make([]string, len(r.Results[0].IncorrectAnswers)),
	}
	for i, v := range r.Results[0].IncorrectAnswers {
		question.IncorrectAnswers[i] = string(v)
	}
	return
}
