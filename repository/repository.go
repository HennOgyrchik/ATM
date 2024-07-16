package repository

import (
	"atm/models"
	"github.com/google/uuid"
)

type Repository interface {
	AddAccount(account models.BankAccount) error
	GetAccount(id uuid.UUID) (models.BankAccount, bool)
}

func New() (Repository, error) {
	return NewStorageInMemory(), nil
}
