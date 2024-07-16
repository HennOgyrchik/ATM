package repository

import (
	"atm/models"
	"github.com/google/uuid"
	"sync"
)

type Accounts struct {
	store map[uuid.UUID]models.BankAccount
	rw    sync.RWMutex
}

func NewStorageInMemory() *Accounts {
	return &Accounts{store: make(map[uuid.UUID]models.BankAccount), rw: sync.RWMutex{}}
}

func (m *Accounts) AddAccount(account models.BankAccount) error {
	m.rw.Lock()
	m.store[account.GetID()] = account
	m.rw.Unlock()
	return nil
}

func (m *Accounts) GetAccount(id uuid.UUID) (models.BankAccount, bool) {
	m.rw.RLock()
	val, ok := m.store[id]
	m.rw.RUnlock()
	return val, ok
}
