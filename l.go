package sampling

import "math"

type L struct {
	reservoir []int64
	counter   int64
	next      int64
	w         float64
}

func NewL(n int) *L {
	l := &L{
		reservoir: make([]int64, n),
		next:      int64(n),
		w:         math.Exp(math.Log(rng.Float64()) / float64(n)),
	}
	return l
}

func (l *L) Offer(value int64) {
	if int(l.counter) < cap(l.reservoir) {
		l.reservoir[l.counter] = value
		return
	} else {
		l.next += int64(math.Log(rng.Float64())/math.Log(1-l.w)) + 1
		if l.counter == l.next {
			l.reservoir[int(rng.Int63n(int64(cap(l.reservoir))))] = value
			l.w *= math.Exp(math.Log(rng.Float64()) / float64(cap(l.reservoir)))
		}
	}
	l.counter++
}
