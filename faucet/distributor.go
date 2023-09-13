package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"sync"

	"go.uber.org/zap"
)

// Distributor holds all functions for performing various actions
type Distributor struct {
	Config      *Config
	Logger      *zap.Logger
	CreditCoins Coins

	Holder *Account
	Addrs  []*Account
}

// NewDistributor returns a new Distributor struct pointer, initilazies all the addresses
func NewDistributor(config *Config, logger *zap.Logger) (*Distributor, error) {
	coins := Coins{}
	for _, coinStr := range strings.Split(config.CreditCoins, ",") {
		matches := reCoins.FindStringSubmatch(coinStr)
		if len(matches) < 2 {
			return nil, fmt.Errorf("validation error: coin expected to be <amount><denom>, found: %s", coinStr)
		}
		coins = append(coins, Coin{Denom: matches[1], Amount: matches[0]})
	}

	holder, err := NewAccount(config, logger, "holder", config.Mnemonic, 0)
	if err != nil {
		return nil, err
	}

	distributor := &Distributor{
		Config:      config,
		Logger:      logger,
		CreditCoins: coins,
		Holder:      holder,
	}

	if config.Concurrency <= 1 {
		return distributor, nil
	}

	var addrs []*Account
	for i := 0; i < config.Concurrency; i++ {
		addr, err := NewAccount(config, logger, fmt.Sprintf("distributor-%d", i), config.Mnemonic, i+1)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	distributor.Addrs = addrs

	return distributor, nil
}

// Refill transfers tokens from genesis to all distributor address with balance bellow threshold
func (d *Distributor) Refill() error {
	return nil
}

// Status returns a map of address and balance of the addresses in distributors
func (d *Distributor) Status() (map[string]Coins, error) {
	return nil, nil
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
	mu          sync.Mutex
	chainBinary string
	chainHome   string
	gas         string

	Logger  *zap.Logger
	Name    string
	Index   int
	Address string
}

// NewAccount function creates the account keys based on name and mnemonic provided
func NewAccount(config *Config, logger *zap.Logger, name string, mnemonic string, index int) (*Account, error) {
	account := &Account{
		chainBinary: config.ChainBinary,
		chainHome:   config.ChainHome,
		gas:         config.DefaultGas,
		Index:       index,
		Name:        name,
	}
	// add key to the keyring
	address, err := account.addKey(name, mnemonic, index)
	if err != nil {
		return nil, err
	}
	account.Address = address
	return nil, nil
}

// addKey adds key to the keyring, returns address
func (a *Account) addKey(name, mnemonic string, index int) (string, error) {
	ok := a.mu.TryLock()
	if !ok {
		return "", fmt.Errorf("account %s busy: %w", a, ErrResourceInUse)
	}
	defer a.mu.Unlock()

	cmdStr := fmt.Sprintf("echo \"%s\" | %s keys add %s --output json --index %d --recover --keyring-backend=\"test\"", mnemonic, a.chainBinary, name, index)
	a.Logger.Debug(fmt.Sprintf("running command to add key: %s", cmdStr))
	out, err := runCommand(cmdStr)
	if err != nil {
		return "", err
	}
	a.Logger.Debug(fmt.Sprintf("key added with output: %s", string(out)))

	keyMap := map[string]interface{}{}
	err = json.Unmarshal(out, &keyMap)
	if err != nil {
		return "", err
	}

	return keyMap["address"].(string), nil
}

func (a *Account) deleteKey(name string) error {
	ok := a.mu.TryLock()
	if !ok {
		return fmt.Errorf("account %s busy: %w", a, ErrResourceInUse)
	}
	defer a.mu.Unlock()

	cmdStr := fmt.Sprintf("%s keys delete %s --force --yes --keyring-backend=\"test\"", a.chainBinary, name)
	a.Logger.Debug(fmt.Sprintf("running command to delete key: %s", cmdStr))
	_, err := runCommand(cmdStr)
	return err
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

	cmdStr := fmt.Sprintf("%s tx bank send %s %s %s%s --keyring-backend=test --gas=auto --gas-adjustment=1.5 --yes", a.chainBinary, a.Address, address, amount, denom)
	cmd := exec.Command("bash", "-c", cmdStr)
	a.Logger.Debug(fmt.Sprintf("running cmd to send tokens: %s", cmd))
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// GetBalance queries for balance of account and populates Balances
// since this is query, no need for blocking. Make http request to rest/lcd endpoint
func (a *Account) GetBalance() (Coins, error) {
	return nil, nil
}

// GetBalanceByDenom queries for balance based on the denom
func (a *Account) GetBalanceByDenom(denom string) (Coin, error) {
	return Coin{}, nil
}
