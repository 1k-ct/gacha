package dogacha

import (
	"errors"
	"math/rand"
	"time"
)

type Pro struct {
}

// New type gacha struct return *gacha
func New() *Pro {
	g := &Pro{}
	return g
}

// Gacha do gacha
// if Gacha(100, 1) is 1%, Gacha(1000, 2) is 0.2%
// 1% chance of true else return false (99%)
func Gacha(all, expectProbability int64) (bool, error) {
	if all < expectProbability {
		return false, errors.New("all < expectProbability")
	}
	rand.Seed(time.Now().UnixNano())
	p := rand.Int63n(all)
	if 0 <= p && p <= expectProbability-1 {
		return true, nil
	}
	return false, nil
}

// Percentages all:100,n:1 => 0.01(1%)
func Percentages(all, n int64) float64 {
	p := float64(n) / float64(all)
	return p
}
