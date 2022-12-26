package main

import "fmt"

type WithdrawError struct {
	account             string
	balance             float64
	reason              string
	amountToBeWithdrawn float64
}

func (werr WithdrawError) Error() string {
	return fmt.Sprintf("Your %v has %v balance/limit, we cannot withdraw %v amount due to %v", werr.account, werr.balance, werr.amountToBeWithdrawn, werr.reason)
}

type Withdrawable interface {
	CanWithdraw(float64) (bool, WithdrawError)
	Withdraw(float64) float64
	GetBalanceOrAvailableLimit() float64
	GetIdentifier() string
}
