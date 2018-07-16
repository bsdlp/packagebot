package opentdb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/bsdlp/packagebot/src/trivia/internal/trivia"
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

func buildURL(opt *trivia.QuestionType) string {
	if opt == nil {
		return "https://opentdb.com/api.php?amount=1&type=multiple&encode=base64"
	}

	u := &url.URL{
		Scheme: "https",
		Host:   "opentdb.com",
		Path:   "api.php",
	}

	kv := u.Query()
	kv.Set("type", "multiple")
	kv.Set("encode", "base64")

	if opt.GetCount() == 0 {
		kv.Set("amount", "1")
	} else {
		kv.Set("amount", fmt.Sprint(opt.GetCount()))
	}
	kv.Set("difficulty", strings.ToLower(opt.GetDifficulty().String()))

	u.RawQuery = kv.Encode()
	return u.String()
}

// GetQuestion retrieves a trivia question
func GetQuestion(ctx context.Context, questionType *trivia.QuestionType) (question *trivia.Question, err error) {
	resp, err := http.Get(buildURL(questionType))
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

	question = &trivia.Question{
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
