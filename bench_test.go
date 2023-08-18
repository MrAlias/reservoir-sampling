package sampling

import "testing"

var result int64

func BenchmarkSampler(b *testing.B) {
	const n = 1024
	const value = 1

	b.Run("BaseRandomInt64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = rng.Int63()
		}
	})

	b.Run("R", func(b *testing.B) {
		sampler := NewR(n)

		// Measure random insert, not initial loading.
		for i := 0; i < n; i++ {
			sampler.Offer(value)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sampler.Offer(value)
		}
	})

	b.Run("L", func(b *testing.B) {
		sampler := NewL(n)

		// Measure random insert, not initial loading.
		for i := 0; i < n; i++ {
			sampler.Offer(value)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sampler.Offer(value)
		}
	})

	b.Run("X", func(b *testing.B) {
		sampler := NewX(n)

		// Measure random insert, not initial loading.
		for i := 0; i < n; i++ {
			sampler.Offer(value)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sampler.Offer(value)
		}
	})

	b.Run("Z", func(b *testing.B) {
		sampler := NewZ(n, 40)

		// Measure random insert, not initial loading.
		for i := 0; i < n; i++ {
			sampler.Offer(value)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sampler.Offer(value)
		}
	})
}
