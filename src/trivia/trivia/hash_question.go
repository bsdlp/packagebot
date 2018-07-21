package trivia

import (
	"hash/fnv"
	"sort"

	"github.com/bsdlp/packagebot/src/trivia/internal/opentdb"
)

func hashQuestion(question *opentdb.Question) (key string, err error) {
	hasher := fnv.New128()
	_, err = hasher.Write([]byte(question.Category + "\n"))
	if err != nil {
		return
	}
	_, err = hasher.Write([]byte(question.Type + "\n"))
	if err != nil {
		return
	}
	_, err = hasher.Write([]byte(question.Difficulty + "\n"))
	if err != nil {
		return
	}
	_, err = hasher.Write([]byte(question.Question + "\n"))
	if err != nil {
		return
	}
	_, err = hasher.Write([]byte(question.CorrectAnswer + "\n"))

	sort.Strings(question.IncorrectAnswers)
	for _, ans := range question.IncorrectAnswers {
		_, err = hasher.Write([]byte(ans + "\n"))
		if err != nil {
			return
		}
	}
	key = string(hasher.Sum(nil))
}
