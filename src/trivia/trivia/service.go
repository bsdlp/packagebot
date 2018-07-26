package trivia

import (
	"context"
	crand "crypto/rand"
	"math/big"
	"math/rand"

	"github.com/bsdlp/packagebot/src/outbound/outbound"
	"github.com/bsdlp/packagebot/src/trivia/internal/opentdb"
	"github.com/go-redis/redis"
)

// RedisClient describes a redis client
type RedisClient interface {
	redis.Cmdable
	WithContext(context.Context) RedisClient
}

// Service implements Trivia
type Service struct {
	Outbound    outbound.Outbound
	QuestionsDB RedisClient
	Rand        *rand.Rand
	GetQuestion func(context.Context, *opentdb.QuestionParameters) (*opentdb.Question, error)
}

// NewCryptoRand creates a crypto/rand backed *rand.Rand
func NewCryptoRand() *rand.Rand {
	return rand.New(cryptoRandSource{})
}

type cryptoRandSource struct{}

func (crs cryptoRandSource) Int63() int64 {
	nBig, err := crand.Int(crand.Reader, big.NewInt(27))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

func (crs cryptoRandSource) Seed(seed int64) {}
