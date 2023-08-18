package sampling

type X struct {
	reservoir []float64
	counter   int64
	next      int64
}

func NewX(n int) *X {
	return &X{
		reservoir: make([]float64, n),
		next:      int64(n),
	}
}

func (x *X) Offer(value float64) {
	if int(x.counter) < cap(x.reservoir) {
		x.reservoir[x.counter] = value
	} else if x.counter == x.next {
		x.reservoir[int(rng.Int63n(int64(cap(x.reservoir))))] = value
		x.next += linearSearch(x.counter, int64(cap(x.reservoir)))
	}
	x.counter++
}

func (x *X) mean() float64 {
	var sum float64
	for _, v := range x.reservoir {
		sum += v
	}
	return sum / float64(cap(x.reservoir))
}

func linearSearch(counter, n int64) int64 {
	var s int64
	u := rng.Float64()
	quotient := float64(counter+1-n) / float64(counter+1)

	var i int64 = 1
	quotient *= float64(counter+1+i-n) / float64(counter+i+1)
	s++
	i++

	for quotient > u {
		quotient *= float64(counter+1+i-n) / float64(counter+i+1)
		s++
		i++
	}

	return s
}
