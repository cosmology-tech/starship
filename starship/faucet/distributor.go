package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"go.uber.org/zap"

	pb "github.com/hyperweb-io/starship/faucet/faucet"
)

// Distributor holds all functions for performing various actions
type Distributor struct {
	config *Config
	logger *zap.Logger

	CreditCoins Coins

	Holder *Account
	Addrs  []*Account
}

// NewDistributor returns a new Distributor struct pointer, initilazies all the addresses
func NewDistributor(config *Config, logger *zap.Logger) (*Distributor, error) {
	coins, err := NewCoinFromStr(config.CreditCoins)
	if err != nil {
		return nil, err
	}

	holder, err := NewAccount(config, logger, "holder", config.Mnemonic, 0)
	if err != nil {
		return nil, err
	}

	distributor := &Distributor{
		config:      config,
		logger:      logger,
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

	// log initial status
	status, err := distributor.Status()
	if err != nil {
		return nil, err
	}
	distributor.logger.Info("status of distributors", zap.Any("status", status))

	// initialze the distributor addresses
	err = distributor.Refill()
	if err != nil {
		return nil, err
	}

	return distributor, nil
}

// requireRefill will return true if the balance is such that it requires a refill
func (d *Distributor) requireRefill(amount string, denom string) bool {
	if amount == "" {
		return true
	}

	bigAmt, _ := new(big.Int).SetString(amount, 0)
	creditAmt := d.CreditCoins.GetDenomAmount(denom)
	bigCreditAmt, _ := new(big.Int).SetString(creditAmt, 0)
	bigFactor := new(big.Int).SetInt64(int64(d.config.RefillThreshold))

	if bigAmt.Cmp(new(big.Int).Mul(bigCreditAmt, bigFactor)) <= 0 {
		return true
	}

	return false
}

// refillAmount will return the ammount that needs to be credited for the denom
func (d *Distributor) refillAmount(denom string) string {
	creditAmt := d.CreditCoins.GetDenomAmount(denom)
	if creditAmt == "" {
		d.logger.Error("credit amount for denom seems to be empty",
			zap.Any("credit tokens", d.CreditCoins))
		return ""
	}

	bigCreditAmt, _ := new(big.Int).SetString(creditAmt, 0)
	d.logger.Debug("credit amount",
		zap.String("big credit amt", bigCreditAmt.String()),
		zap.String("credit amt", creditAmt),
		zap.Any("credit tokens", d.CreditCoins))
	bigFactor := new(big.Int).SetInt64(int64(d.config.RefillFactor))

	refillAmt := new(big.Int).Mul(bigCreditAmt, bigFactor)

	return fmt.Sprintf("%v", refillAmt)
}

// Refill transfers tokens from genesis to all distributor address with balance bellow threshold
func (d *Distributor) Refill() error {
	for _, account := range d.Addrs {
		balances, err := account.GetBalance()
		if err != nil {
			return nil
		}
		for _, creditCoin := range d.CreditCoins {
			balanceAmt := balances.GetDenomAmount(creditCoin.Denom)
			if !d.requireRefill(balanceAmt, creditCoin.Denom) {
				d.logger.Debug("skipping refill for address", zap.String("distributor_address", account.Address))
				continue
			}
			d.logger.Info("refilling address from main holder", zap.String("distributor_address", account.Address))
			err = d.Holder.SendTokens(account.Address, creditCoin.Denom, d.refillAmount(creditCoin.Denom))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Status returns a map of address and balance of the addresses in distributors
func (d *Distributor) Status() ([]AccountBalances, error) {
	holderCoins, err := d.Holder.GetBalance()
	if err != nil {
		return nil, err
	}
	accountBalances := []AccountBalances{
		{
			Account:  d.Holder,
			Balances: holderCoins,
		},
	}

	for _, account := range d.Addrs {
		coins, err := account.GetBalance()
		if err != nil {
			return nil, err
		}
		accountBalances = append(accountBalances, AccountBalances{Account: account, Balances: coins})
	}

	return accountBalances, nil
}

// SendTokens will transfer tokens to the given address and denom from one of distributor addresses
func (d *Distributor) SendTokens(address string, denom string) error {
	amount := d.CreditCoins.GetDenomAmount(denom)

	if d.Addrs == nil {
		return d.Holder.SendTokens(address, denom, amount)
	}

	randIndex := rand.Intn(len(d.Addrs))
	if amount == "" {
		return fmt.Errorf("invalid denom: %s, expected denoms: %s", denom, d.CreditCoins.GetDenoms())
	}

	return d.Addrs[randIndex].SendTokens(address, denom, amount)
}

type AccountBalances struct {
	Account  *Account
	Balances Coins
}

func (ab AccountBalances) String() string {
	return fmt.Sprintf("address: %s, coins: %s", ab.Account.Address, ab.Balances)
}

func (ab AccountBalances) ToProto() *pb.AddressBalance {
	var balances []*pb.Coin
	for _, coin := range ab.Balances {
		balances = append(balances, &pb.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount,
		})
	}
	proto := &pb.AddressBalance{
		Address: ab.Account.Address,
		Balance: balances,
	}

	return proto
}

type Account struct {
	mu sync.Mutex

	logger *zap.Logger
	config *Config

	Name    string
	Index   int
	Address string
}

// NewAccount function creates the account keys based on name and mnemonic provided
func NewAccount(config *Config, logger *zap.Logger, name string, mnemonic string, index int) (*Account, error) {
	account := &Account{
		logger: logger,
		config: config,
		Index:  index,
		Name:   name,
	}
	// add key to the keyring
	address, err := account.addKey(name, mnemonic, index)
	if err != nil {
		return nil, err
	}
	account.Address = address

	return account, nil
}

// addKey adds key to the keyring, returns address
func (a *Account) addKey(name, mnemonic string, index int) (string, error) {
	ok := a.mu.TryLock()
	if !ok {
		return "", fmt.Errorf("account %s busy: %w", a, ErrResourceInUse)
	}
	defer a.mu.Unlock()

	args := "--output json --recover --keyring-backend=\"test\""
	cmdStr := fmt.Sprintf("echo \"%s\" | %s keys add %s --index %d %s", mnemonic, a.config.ChainBinary, name, index, args)
	a.logger.Debug(fmt.Sprintf("running command to add key: %s", cmdStr))

	out, err := runCommand(cmdStr)
	if err != nil {
		return "", err
	}
	a.logger.Debug(fmt.Sprintf("key added with output: %s", string(out)))

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

	cmdStr := fmt.Sprintf("%s keys delete %s --force --yes --keyring-backend=\"test\"", a.config.ChainBinary, name)
	a.logger.Debug(fmt.Sprintf("running command to delete key: %s", cmdStr))
	_, err := runCommand(cmdStr)

	return err
}

func (a *Account) String() string {
	return fmt.Sprintf("name: %s, addr: %s", a.Name, a.Address)
}

// sendTokens performs chain binary send txn from account
func (a *Account) sendTokens(address string, denom string, amount string) error {
	ok := a.mu.TryLock()
	if !ok {
		return fmt.Errorf("account %s busy: %w", a, ErrResourceInUse)
	}
	defer a.mu.Unlock()

	args := fmt.Sprintf("--chain-id=%s --fees=%s --keyring-backend=test --gas=auto --gas-adjustment=1.5 --yes --node=%s", a.config.ChainId, a.config.ChainFees, a.config.ChainRPCEndpoint)
	cmdStr := fmt.Sprintf("%s tx bank send %s %s %s%s %s", a.config.ChainBinary, a.Address, address, amount, denom, args)
	output, err := runCommand(cmdStr)
	if err != nil {
		a.logger.Error("send token failed", zap.String("cmd", cmdStr), zap.Error(err))
		return err
	}
	a.logger.Info("ran cmd to send tokens", zap.String("cmd", cmdStr), zap.String("stdout", string(output)))

	return nil
}

// SendTokens will perform send tokens with retries based on errors
func (a *Account) SendTokens(address string, denom string, amount string) error {
	err := a.sendTokens(address, denom, amount)
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "account sequence mismatch") {
		// retry sendTokens
		a.logger.Debug("got account sequence missmatch error, retrying send tokens recursively")
		return a.SendTokens(address, denom, amount)
	}

	return err
}

// GetBalance queries for balance of account and populates Balances
// since this is query, no need for blocking. Make http request to rest/lcd endpoint
func (a *Account) GetBalance() (Coins, error) {
	url := fmt.Sprintf("%s%s/%s", a.config.ChainRESTEndpoint, a.config.ChainBalancesURI, a.Address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	balances, ok := data["balances"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("balances not found in response: %v", data)
	}

	coins := Coins{}
	for _, balance := range balances {
		coins = append(coins, Coin{
			balance.(map[string]interface{})["denom"].(string),
			balance.(map[string]interface{})["amount"].(string),
		})
	}

	return coins, nil
}

// GetBalanceByDenom queries for balance based on the denom
func (a *Account) GetBalanceByDenom(denom string) (Coin, error) {
	coins, err := a.GetBalance()
	if err != nil {
		return Coin{}, err
	}
	coinAmt := coins.GetDenomAmount(denom)
	if coinAmt == "" {
		return Coin{}, fmt.Errorf("denom %s not found in balacne with denoms: %s", denom, coins.GetDenoms())
	}

	return Coin{Denom: denom, Amount: coinAmt}, nil
}
