package sampling

import "math"

type Z struct {
	reservoir []float64
	counter   int64
	next      int64
	threshold int
}

func NewZ(n, threshold int) *Z {
	z := &Z{
		reservoir: make([]float64, n),
		next:      int64(n),
		threshold: threshold,
	}
	z.advance()
	return z
}

func (z *Z) Offer(value float64) {
	if int(z.counter) < cap(z.reservoir) {
		z.reservoir[z.counter] = value
	} else if z.counter == z.next {
		z.reservoir[int(rng.Int63n(int64(cap(z.reservoir))))] = value
		z.advance()
	}
	z.counter++
}

func (z *Z) advance() {
	if z.counter <= int64(z.threshold*cap(z.reservoir)) {
		z.next += linearSearch(z.counter, int64(cap(z.reservoir)))
		return
	}

	c := float64(z.counter+1) / float64(z.counter-int64(cap(z.reservoir))+1)
	w := math.Exp(-math.Log(rng.Float64()) / float64(cap(z.reservoir)))

	var s int64

	for {
		u := rng.Float64()
		x := float64(z.counter) * (w - 1)
		s = int64(x)
		g := float64(cap(z.reservoir)) / (float64(z.counter) + x)
		p := math.Pow(
			float64(z.counter-int64(cap(z.reservoir))+1)/float64(z.counter+s-int64(cap(z.reservoir))+1),
			float64(cap(z.reservoir)+1),
		)
		h := float64(cap(z.reservoir)) / (float64(z.counter) + 1) * p

		if u <= (c*g)/h {
			break
		}

		// Slow path, need to check f.
		var f float64 = 1
		for i := int64(0); i <= s; i++ {
			f *= float64(z.counter-int64(cap(z.reservoir))+i) / float64(z.counter+1+i)
		}
		f *= float64(cap(z.reservoir))
		f /= float64(z.counter - int64(cap(z.reservoir)))
		if u <= (c*g)/f {
			break
		}
		w = math.Exp(-math.Log(rng.Float64()) / float64(cap(z.reservoir)))
	}

	z.next += s + 1
}

func (z *Z) mean() float64 {
	var sum float64
	for _, v := range z.reservoir {
		sum += v
	}
	return sum / float64(cap(z.reservoir))
}
