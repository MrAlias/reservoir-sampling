package sampling

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgorithms(t *testing.T) {
	t.Run("R", testReservoir(func(n int) reservoir { return NewR(n) }))
	t.Run("L", testReservoir(func(n int) reservoir { return NewL(n) }))
	t.Run("X", testReservoir(func(n int) reservoir { return NewX(n) }))
	t.Run("Z", testReservoir(func(n int) reservoir { return NewZ(n, 40) }))
}

func testReservoir(newRes func(int) reservoir) func(t *testing.T) {
	intensity := 0.1
	sampleSize := 1000

	data := make([]float64, sampleSize*1000)
	for i := range data {
		data[i] = (-1.0 / intensity) * math.Log(rng.Float64())
	}
	// Sort to avoid position bias.
	sort.Float64s(data)

	return func(t *testing.T) {
		t.Helper()

		r := newRes(sampleSize)
		for _, value := range data {
			r.Offer(value)
		}

		assert.InDelta(t, 1/r.mean(), intensity, 0.01)
	}
}
