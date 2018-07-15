package opentdb

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalJSON(t *testing.T) {
	testData := `{"response_code":0,"results":[{"category":"RW50ZXJ0YWlubWVudDogSmFwYW5lc2UgQW5pbWUgJiBNYW5nYQ==","type":"bXVsdGlwbGU=","difficulty":"bWVkaXVt","question":"QW9pIE1peWFtb3JpIGlzIHRoZSBwcm9kdWN0aW9uIG1hbmFnZXIgb2Ygd2hhdCBhbmltZSBpbiAiU2hpcm9iYWtvIj8=","correct_answer":"VGhlIFRoaXJkIEFlcmlhbCBHaXJscyBTcXVhZA==","incorrect_answers":["RXhvZHVzIQ==","QW5kZXMgQ2h1Y2t5","QW5nZWwgQmVhdHMh"]}]}`
	expected := question{
		Category:      "Entertainment: Japanese Anime & Manga",
		Type:          "multiple",
		Difficulty:    "medium",
		Question:      `Aoi Miyamori is the production manager of what anime in "Shirobako"?`,
		CorrectAnswer: "The Third Aerial Girls Squad",
		IncorrectAnswers: []base64EncodedString{
			"Exodus!",
			"Andes Chucky",
			"Angel Beats!",
		},
	}

	var resp apiResponse
	err := json.Unmarshal([]byte(testData), &resp)
	require.NoError(t, err)

	assert.Equal(t, expected, resp.Results[0])
}

func TestGetQuestion(t *testing.T) {
	question, err := GetQuestion(context.TODO())
	require.NoError(t, err)

	assert.NotEmpty(t, question.Category)
	assert.NotEmpty(t, question.Type)
	assert.NotEmpty(t, question.Difficulty)
	assert.NotEmpty(t, question.Question)
	assert.NotEmpty(t, question.CorrectAnswer)
	assert.NotEmpty(t, question.IncorrectAnswers)
}
