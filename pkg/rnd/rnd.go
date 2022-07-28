package rnd

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Generator interface {
	RandomInt(length int) string
	RandomStr(length int) (string, error)
}

type GeneratorRand struct {
	rand *rand.Rand
}

func NewGeneratorRand() *GeneratorRand {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &GeneratorRand{
		rand: rand,
	}
}

//Returns value in the min to max range
func (g *GeneratorRand) RandomInt(length int) string {
	if length < 1 || length >= 10 {
		return ""
	}
	min := int(math.Pow10(length - 1))
	max := min*10 - 1
	return strconv.Itoa(g.rand.Intn(max-min+1) + min)
}

//Returns random string by length
func (g *GeneratorRand) RandomStr(length int) (string, error) {
	b := make([]byte, length)
	_, err := g.rand.Read(b)
	if err != nil {
		return "", nil
	}
	token := fmt.Sprintf("%x", b)[:length]
	return token, nil
}
