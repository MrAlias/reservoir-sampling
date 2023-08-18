package sampling

import "math"

type L struct {
	reservoir []float64
	counter   int64
	next      int64
	w         float64
}

func NewL(n int) *L {
	l := &L{
		reservoir: make([]float64, n),
		next:      int64(n),
		w:         math.Exp(math.Log(rng.Float64()) / float64(n)),
	}
	l.advance()
	return l
}

func (l *L) Offer(value float64) {
	if int(l.counter) < cap(l.reservoir) {
		l.reservoir[l.counter] = value
	} else {
		if l.counter == l.next {
			l.reservoir[int(rng.Int63n(int64(cap(l.reservoir))))] = value
			l.advance()
		}
	}
	l.counter++
}

func (l *L) advance() {
	l.w *= math.Exp(math.Log(rng.Float64()) / float64(cap(l.reservoir)))
	l.next += int64(math.Log(rng.Float64())/math.Log(1-l.w)) + 1
}

func (l *L) mean() float64 {
	var sum float64
	for _, v := range l.reservoir {
		sum += v
	}
	return sum / float64(cap(l.reservoir))
}
