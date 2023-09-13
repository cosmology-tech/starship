package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

var reCoins = regexp.MustCompile("([0-9]+)([a-zA-Z]+)")

type Coins []Coin

func (c Coins) GetDenomAmount(denom string) string {
	for _, coin := range c {
		if coin.Denom == denom {
			return coin.Amount
		}
	}
	return ""
}

func (c Coins) GetDenoms() []string {
	denoms := []string{}
	for _, coin := range c {
		denoms = append(denoms, coin.Denom)
	}
	return denoms
}

// IsDenom returns true if denom found in coins else false
func (c Coins) IsDenom(denom string) bool {
	for _, coin := range c {
		if coin.Denom == denom {
			return true
		}
	}
	return false
}

type Coin struct {
	Denom  string `name:"denom" json:"denom,omitempty" yaml:"denom"`
	Amount string `name:"amount" json:"amount,omitempty" yaml:"amount"`
}

// Distributor holds all functions for performing various actions
type Distributor struct {
	Config      *Config
	CreditCoins Coins

	Genesis *Account
	Addrs   []*Account
}

// NewDistributor returns a new Distributor struct pointer, initilazies all the addresses
func NewDistributor(config *Config) (*Distributor, error) {
	coins := Coins{}
	for _, coinStr := range strings.Split(config.CreditCoins, ",") {
		matches := reCoins.FindStringSubmatch(coinStr)
		if len(matches) < 2 {
			return nil, fmt.Errorf("validation error: coin expected to be <amount><denom>, found: %s", coinStr)
		}
		coins = append(coins, Coin{Denom: matches[1], Amount: matches[0]})
	}

	genesis, err := NewAccount(config, "genesis", config.Mnemonic)
	if err != nil {
		return nil, err
	}
	if config.Concurrency <= 1 {
		return &Distributor{
			Config:      config,
			CreditCoins: coins,
			Genesis:     genesis,
		}, nil
	}
	addrs := []*Account{}
	for i := 0; i < config.Concurrency; i++ {
		addr, err := NewAccount(config, fmt.Sprintf("distributor-%d", i), "")
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	return &Distributor{
		Config:      config,
		CreditCoins: coins,
		Genesis:     genesis,
		Addrs:       addrs,
	}, nil
}

// Refill transfers tokens from genesis to all distributor address with balance bellow threshold
func (d *Distributor) Refill() error {
	return nil
}

// SendTokens will transfer tokens to the given address and denom from one of distributor addresses
func (d *Distributor) SendTokens(address string, denom string) error {
	randIndex := rand.Intn(len(d.Addrs))
	amount := d.CreditCoins.GetDenomAmount(denom)
	if amount == "" {
		return fmt.Errorf("invalid denom: %s, expected denoms: %s", denom, d.CreditCoins.GetDenoms())
	}
	return d.Addrs[randIndex].SendTokens(address, denom, amount)
}

type Account struct {
	mu sync.Mutex

	Name    string
	Address string
}

// NewAccount function creates the account keys based on name and mnemonic provided
func NewAccount(config *Config, name string, mnemonic string) (*Account, error) {
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

// GetBalance queries for balance of account and populates Balances
// since this is query, no need for blocking
func (a *Account) GetBalance() (Coins, error) {
	return nil, nil
}

// GetBalanceByDenom queries for balance based on the denom
func (a *Account) GetBalanceByDenom(denom string) (Coin, error) {
	return Coin{}, nil
}
