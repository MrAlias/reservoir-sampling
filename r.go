package sampling

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type R struct {
	reservoir []float64
	counter   int64
}

func NewR(n int) *R {
	return &R{reservoir: make([]float64, n)}
}

func (r *R) Offer(value float64) {
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

func (r *R) mean() float64 {
	var sum float64
	for _, v := range r.reservoir {
		sum += v
	}
	return sum / float64(cap(r.reservoir))
}
