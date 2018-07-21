package trivia

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"

	"github.com/bsdlp/packagebot/src/outbound/outbound"
	"github.com/go-redis/redis"
)

// Service implements Trivia
type Service struct {
	Outbound    outbound.Outbound
	QuestionsDB redis.Cmdable
	Rand        *rand.Rand
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
	n := nBig.Int64()
}

func (crs cryptoRandSource) Seed(seed int64) {}
