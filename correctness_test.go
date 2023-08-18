package sampling

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgorithms(t *testing.T) {
	var (
		data       []float64
		sampleSize int
		intensity  float64
	)

	reset := func() {
		intensity = 0.1
		sampleSize = 1000

		data = make([]float64, sampleSize*1000)
		for i := range data {
			data[i] = (-1.0 / intensity) * math.Log(rng.Float64())
		}
		// Sort to avoid position bias.
		sort.Float64s(data)
	}

	t.Run("L", func(t *testing.T) {
		reset()

		l := NewL(sampleSize)
		for _, value := range data {
			l.Offer(value)
		}

		assert.InDelta(t, 1/l.mean(), intensity, 0.01)
	})

	t.Run("R", func(t *testing.T) {
		reset()

		r := NewR(sampleSize)
		for _, value := range data {
			r.Offer(value)
		}

		assert.InDelta(t, 1/r.mean(), intensity, 0.01)
	})

	t.Run("X", func(t *testing.T) {
		reset()

		x := NewX(sampleSize)
		for _, value := range data {
			x.Offer(value)
		}

		assert.InDelta(t, 1/x.mean(), intensity, 0.01)
	})

	t.Run("Z", func(t *testing.T) {
		reset()

		x := NewZ(sampleSize, 40)
		for _, value := range data {
			x.Offer(value)
		}

		assert.InDelta(t, 1/x.mean(), intensity, 0.01)
	})
}
