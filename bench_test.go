package sampling

import "testing"

func BenchmarkSampler(b *testing.B) {
	const n = 1024
	const value = 1

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
}
