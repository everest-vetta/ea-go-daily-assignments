package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinCosZero(t *testing.T) {
	sin, cos := sincos(0)
	assert.Equal(t, 0.0, sin)
	assert.Equal(t, 1.0, cos)
}

func TestTanZero(t *testing.T) {
	assert.Equal(t, 0.0, tan(0))
}
