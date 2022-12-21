package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCashFromFakeSavingsAccount(t *testing.T) {

	fpa := &FakePrimaryAccount{balance : 500}
	fsa := &FakeSecondaryAccount{ balance : 0}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fpa.GetIdentifier(), accType)
}

func TestGetCashFromFakeSecondaryAccount(t *testing.T) {

	fpa := &FakePrimaryAccountWithZeroBalance{}
	fsa := &FakeSecondaryAccount{}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fsa.GetIdentifier(), accType)
}

func TestGetCashFromSavingsAccount(t *testing.T) {

	sa := &SavingsAccount{balance:500}
	ca := &CreditCardAccount{balance:0}

	qc := QuickCash{
		sa,
		ca,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, sa.GetIdentifier(), accType)
}

func TestGetCashFromCreditCardAccount(t *testing.T) {

	sa := &SavingsAccount{balance:0}
	ca := &CreditCardAccount{balance:500}

	qc := QuickCash{
		sa,
		ca,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, ca.GetIdentifier(), accType)
}