package main

import (
	"fmt"
	"regexp"
	"strings"
)

var reCoins = regexp.MustCompile("([0-9]+)([a-zA-Z]+)")

type Coin struct {
	Denom  string `name:"denom" json:"denom,omitempty" yaml:"denom"`
	Amount string `name:"amount" json:"amount,omitempty" yaml:"amount"`
}

func (c Coin) String() string {
	return fmt.Sprintf("denom: %s, amount: %s", c.Denom, c.Amount)
}

type Coins []Coin

// NewCoinFromStr given a comma seperated string of coins, returns Coins
func NewCoinFromStr(coinsStr string) (Coins, error) {
	coins := Coins{}
	for _, cs := range strings.Split(coinsStr, ",") {
		matches := reCoins.FindStringSubmatch(cs)
		if len(matches) < 2 {
			return nil, fmt.Errorf("validation error: coin expected to be <amount><denom>, found: %s", cs)
		}
		coins = append(coins, Coin{Denom: matches[2], Amount: matches[1]})
	}

	return coins, nil
}

func (c Coins) String() string {
	coinsStrs := []string{}
	for _, coin := range c {
		coinsStrs = append(coinsStrs, coin.String())
	}

	return strings.Join(coinsStrs, ",")
}

func (c Coins) GetCoinByDenom(denom string) (Coin, error) {
	for _, coin := range c {
		if coin.Denom == denom {
			return coin, nil
		}
	}

	return Coin{}, fmt.Errorf("denom %s not found in coins with denoms %s", denom, c.GetDenoms())
}

func (c Coins) MustGetCoinByDenom(denom string) Coin {
	coin, err := c.GetCoinByDenom(denom)
	if err != nil {
		panic(err)
	}

	return coin
}

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

// HasDenom returns true if denom found in coins else false
func (c Coins) HasDenom(denom string) bool {
	for _, coin := range c {
		if coin.Denom == denom {
			return true
		}
	}

	return false
}
