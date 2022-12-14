package bank

import "fmt"

type Customer struct {
	clientAccount *Account
	name          string
}

func (c *Customer) SetName(name string) {
	c.name = name
}

func (c *Customer) SetAccount(account *Account) {
	c.clientAccount = account
}

func (c Customer) String() string {
	return fmt.Sprintf("Name: %v \nAccount: %v", c.name, c.clientAccount)
}

// function to withdraw money from an account
func (c *Customer) Withdraw(amt float64) error {
	return c.clientAccount.Withdraw(amt)
}

// function to deposit money to an account
func (c *Customer) Deposit(amt float64) error {
	return c.clientAccount.Deposit(amt)
}
