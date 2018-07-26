package opentdb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const userAgent = "packagebot opentdb client (github.com/bsdlp/packagebot)"

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

// Question is a trivia question
type Question struct {
	Category         string
	Type             string
	Difficulty       string
	Question         string
	CorrectAnswer    string
	IncorrectAnswers []string
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

func buildURL(opt *QuestionParameters) string {
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

	kv.Set("amount", "1")
	kv.Set("difficulty", strings.ToLower(opt.Difficulty))

	u.RawQuery = kv.Encode()
	return u.String()
}

// QuestionParameters contains the parameters for the question to retrieve
type QuestionParameters struct {
	Difficulty string
}

// GetQuestion retrieves a trivia question
func GetQuestion(ctx context.Context, opts *QuestionParameters) (question *Question, err error) {
	req, err := http.NewRequest(http.MethodGet, buildURL(opts), http.NoBody)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
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

	question = &Question{
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
