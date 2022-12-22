package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	err := acc.Withdraw(200)
	if err != nil {
		t.Errorf("Withdrawal not possible")
	}

	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestWithdrawalInsufficientBalance(t *testing.T) {
	acc := Account{balance: 500}

	err := acc.Withdraw(600)
	withdrawError := err.(WithdrawError)

	assert.Equal(t, "Withdraw is not possible with current balance 500.000000. "+
		"Account is lacking balance of 100.000000", withdrawError.Error())
}
