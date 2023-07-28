package sampling

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type R struct {
	reservoir []int64
	counter   int64
}

func NewR(n int) *R {
	return &R{reservoir: make([]int64, n)}
}

func (r *R) Offer(value int64) {
	r.counter++
	if int(r.counter) <= cap(r.reservoir) {
		r.reservoir[r.counter-1] = value
		return
	}

	j := int(rng.Int63n(int64(r.counter)))
	if j < cap(r.reservoir) {
		r.reservoir[j] = value
	}
}
