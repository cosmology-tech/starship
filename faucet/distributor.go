package main

import (
	"fmt"
	"os/exec"
	"sync"
)

// Distributor holds all functions for performing various actions
type Distributor struct {
	Genesis *Account
	Addrs   []*Account
}

func NewDistributor() (*Distributor, error) {
	return nil, nil
}

type Balance struct {
	Denom  string `name:"denom" json:"denom,omitempty" yaml:"denom"`
	Amount string `name:"amount" json:"amount,omitempty" yaml:"amount"`
}

type Account struct {
	mu sync.Mutex

	Name     string
	Address  string
	Balances []Balance
}

// NewAccount function creates the account keys based on name and mnemonic provided
func NewAccount(config Config, name string, mnemonic string) (*Account, error) {
	// todo: add keys in keyring. Has to use Chain bin
	exec.Command(config.ChainBinary, "keys", "add", "...")
	return nil, nil
}

func (a *Account) String() string {
	return fmt.Sprintf("name: %s, addr: %s", a.Name, a.Address)
}

// SendTokens performs chain binary send txn from account
func (a *Account) SendTokens(address string, denom string, amount string) error {
	ok := a.mu.TryLock()
	if !ok {
		return fmt.Errorf("account %s busy: %w", a, ErrResourceInUse)
	}
	defer a.mu.Unlock()
	// todo: cmd command to send tokens
	return nil
}

// FetchBalance queries for balance of account and populates Balances
// since this is query, no need for blocking
func (a *Account) FetchBalance() error {
	return nil
}

func (a *Account) GetBalanceByDenom(denom string) string {
	for _, b := range a.Balances {
		if b.Denom == denom {
			return b.Amount
		}
	}
	return ""
}
