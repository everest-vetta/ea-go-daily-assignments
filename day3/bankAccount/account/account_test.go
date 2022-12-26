package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	balance, _ := acc.Withdraw(200)

	assert.Equal(t, float64(300), balance)
}

func TestUnsucessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}
	balance, withdrawError := acc.Withdraw(1000)

	assert.Equal(t, "Insufficient Balance", withdrawError.reason)
	assert.Equal(t, float64(500), balance)
	assert.Equal(t, "Error Occured \n Reason : Insufficient Balance \nCurrent Balance in your account is 500 your account is lacking 500 amount\n", withdrawError.Error())
}
