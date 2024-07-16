package models

import "github.com/google/uuid"

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
	GetID() uuid.UUID
}

func NewAccount(id uuid.UUID) BankAccount {
	return &Account{id: id, balance: 0.0}
}
