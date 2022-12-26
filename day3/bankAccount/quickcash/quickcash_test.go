package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromSavingsAccount(t *testing.T) {

	fpa := &FakePrimaryAccount{}
	fsa := &FakeSecondaryAccount{}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.GetCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fpa.GetIdentifier(), accType)
}

func TestGetCashFromSecondaryAccount(t *testing.T) {

	fpa := &FakePrimaryAccountWithZeroBalance{}
	fsa := &FakeSecondaryAccount{}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.GetCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fsa.GetIdentifier(), accType)
}
