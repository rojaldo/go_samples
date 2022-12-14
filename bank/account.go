package bank

import (
	"errors"
	"fmt"
)

type Account struct {
	name        string
	amount      float64
	is_customer bool
}

func (a *Account) SetName(name string) {
	a.name = name
}

func (a *Account) SetAmount(amt float64) {
	a.amount = amt
}

func (a *Account) SetIs_customer(is_customer bool) {
	a.is_customer = is_customer
}

func (a *Account) Withdraw(amt float64) error {

	if !a.is_customer {
		return errors.New("Not a customer")
	}
	if amt > a.amount {
		return errors.New("Not enough money")
	}
	a.amount = a.amount - amt
	return nil

}

func (a *Account) Deposit(amt float64) error {
	if !a.is_customer {
		return errors.New("Not a customer")
	}
	a.amount = a.amount + amt
	return nil

}

func (a Account) String() string {
	data := "Name: " + a.name + "\nAmount: " + fmt.Sprintf("%f", a.amount)
	return (data)

}
