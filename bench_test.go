package sampling

import "testing"

var (
	randInt64   int64
	randFloat64 float64
)

func BenchmarkRNG(b *testing.B) {
	b.Run("Int63", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			randInt64 = rng.Int63()
		}
	})

	b.Run("Float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			randFloat64 = rng.Float64()
		}
	})
}

func BenchmarkSampler(b *testing.B) {
	b.Run("R", benchSampler(func(n int) reservoir { return NewR(n) }))
	b.Run("L", benchSampler(func(n int) reservoir { return NewL(n) }))
	b.Run("X", benchSampler(func(n int) reservoir { return NewX(n) }))
	b.Run("Z", benchSampler(func(n int) reservoir { return NewZ(n, 40) }))
}

type reservoir interface {
	Offer(float64)
	mean() float64
}

func benchSampler(r func(int) reservoir) func(b *testing.B) {
	const n = 1024
	const value = 1
	sampler := r(n)

	// Measure random insert, not initial loading.
	for i := 0; i < n; i++ {
		sampler.Offer(value)
	}

	return func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sampler.Offer(value)
		}
	}
}
