package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWithdrawWithSufficientBalance(t *testing.T) {
	ac := &BankAccount{&Card{"Debit-test", 1234, 2024, 12}, "Test Saving Account", 100}
	result, _ := ac.CanWithdraw(50)
	assert.Equal(t, true, result)
}

func TestCanWithdrawWithInsufficientBalance(t *testing.T) {
	ac := &BankAccount{&Card{"Debit-test", 1234, 2024, 12}, "Test Saving Account", 100}
	result, err := ac.CanWithdraw(150)
	assert.Equal(t, false, result)
	assert.Equal(t, "Insufficient balance/limit", err.reason)
}

func TestWithdraw(t *testing.T) {
	ac := &BankAccount{&Card{"Debit-test", 1234, 2024, 12}, "Test Saving Account", 100}
	balance := ac.Withdraw(50)
	assert.Equal(t, 50.0, balance)
}

func TestGetBalanceOrAvailableLimit(t *testing.T) {
	ac := &BankAccount{&Card{"Debit-test", 1234, 2024, 12}, "Test Saving Account", 100}
	assert.Equal(t, 100.0, ac.GetBalanceOrAvailableLimit())
}

func TestGetIdentifier(t *testing.T) {
	ac := &BankAccount{&Card{"Debit-test", 1234, 2024, 12}, "Test Saving Account", 100}
	assert.Equal(t, "Test Saving Account", ac.GetIdentifier())
}

func TestWalletCanWithdrawWithSufficientBalance(t *testing.T) {
	ac := &Wallet{99, "Test Wallet", 100}
	result, _ := ac.CanWithdraw(50)
	assert.Equal(t, true, result)
}

func TestWalletCanWithdrawWithInsufficientBalance(t *testing.T) {
	ac := &Wallet{99, "Test Wallet", 100}
	result, err := ac.CanWithdraw(150)
	assert.Equal(t, false, result)
	assert.Equal(t, "Insufficient balance/limit", err.reason)
}

func TestWalletWithdraw(t *testing.T) {
	ac := &Wallet{99, "Test Wallet", 100}
	balance := ac.Withdraw(50)
	assert.Equal(t, 50.0, balance)
}

func TestWalletGetBalanceOrAvailableLimit(t *testing.T) {
	ac := &Wallet{99, "Test Wallet", 100}
	assert.Equal(t, 100.0, ac.GetBalanceOrAvailableLimit())
}

func TestWalletGetIdentifier(t *testing.T) {
	ac := &Wallet{99, "Test Wallet", 100}
	assert.Equal(t, "Test Wallet 99", ac.GetIdentifier())
}
