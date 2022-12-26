package main

import (
	"fmt"
)

type QuickCash struct {
	accounts []Withdrawable
}

func (qc *QuickCash) InitializeAccounts() {
	sa := &BankAccount{&Card{"Debit", 1234, 2024, 12}, "Saving Account", 1000}
	ca := &BankAccount{&Card{"Credit", 2345, 2024, 12}, "Current Account", 500}
	w := &Wallet{9999, "GPay", 1000}
	qc.accounts = append(qc.accounts, sa, ca, w)
}

func (qc *QuickCash) GetCash(amount []float64) []string {
	result := []string{}
	qc.InitializeAccounts()
	for i := range amount {
		for n := range qc.accounts {
			if canwithdraw, err := qc.accounts[n].CanWithdraw(amount[i]); canwithdraw {
				fmt.Printf("Withdrawing %v from %v\n", amount[i], qc.accounts[n].GetIdentifier())
				qc.accounts[n].Withdraw(amount[i])
				result = append(result, qc.accounts[n].GetIdentifier())
				break
			} else {
				fmt.Println(err.Error())
			}
		}
	}
	return result
}
