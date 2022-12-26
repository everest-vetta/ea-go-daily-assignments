package account

import "fmt"

// TDD Bank Account app

type Account struct {
	balance float64
}

type WithdrawError struct {
	balance             float64
	amountToBeWithdrawn float64
	reason              string
}

func (err WithdrawError) Error() string {
	return fmt.Sprintf("Error Occured \n Reason : %v \nCurrent Balance in your account is %v your account is lacking %v amount\n", err.reason, err.balance, (err.amountToBeWithdrawn - err.balance))
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) (float64, WithdrawError) {
	if acc.balance < amount {
		err := WithdrawError{acc.balance, amount, "Insufficient Balance"}
		return acc.GetBalance(), err
	}
	acc.balance -= amount
	return acc.GetBalance(), WithdrawError{}
}
