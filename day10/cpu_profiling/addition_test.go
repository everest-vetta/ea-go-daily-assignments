package cpu_profiling

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAddition(t *testing.T) {
	sum := Add(5.0, 2.0)
	assert.Equal(t, 7.0, sum)
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(5.0, 2.0)
	}
}
