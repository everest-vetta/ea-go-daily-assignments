package main

import "fmt"

type Card struct {
	cardType      string // credit or debit
	cardNo        int
	validityYear  int
	validityMonth int
}

type BankAccount struct {
	card       *Card
	identifier string
	balance    float64
}

func (pa *BankAccount) CanWithdraw(amount float64) (bool, WithdrawError) {
	if amount <= pa.balance {
		return true, WithdrawError{}
	} else {
		return false, WithdrawError{pa.identifier, pa.balance, "Insufficient balance/limit", amount}
	}
}

func (pa *BankAccount) Withdraw(amount float64) float64 {
	//fmt.Println("Initial Balance", pa.balance)
	pa.balance -= amount
	return pa.balance
}

func (pa *BankAccount) GetBalanceOrAvailableLimit() float64 {
	return pa.balance
}

func (pa *BankAccount) GetIdentifier() string {
	return pa.identifier
}

type Wallet struct {
	phoneNo    int
	identifier string
	balance    float64
}

func (w *Wallet) CanWithdraw(amount float64) (bool, WithdrawError) {
	if amount <= w.balance {
		return true, WithdrawError{}
	} else {
		return false, WithdrawError{w.identifier, w.balance, "Insufficient balance/limit", amount}
	}
}

func (w *Wallet) Withdraw(amount float64) float64 {
	//fmt.Println("Initial Balance", pa.balance)
	w.balance -= amount
	return w.balance
}

func (w *Wallet) GetBalanceOrAvailableLimit() float64 {
	return w.balance
}

func (w *Wallet) GetIdentifier() string {
	return fmt.Sprintf("%v %v", w.identifier, w.phoneNo)
}
