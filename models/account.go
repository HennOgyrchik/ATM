package models

import (
	"errors"
	"github.com/google/uuid"
)

type Account struct {
	id      uuid.UUID
	balance float64
}

func (a *Account) Deposit(amount float64) error {
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount

	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) GetID() uuid.UUID {
	return a.id
}
