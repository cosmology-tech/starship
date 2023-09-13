package main

import "regexp"

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

// HasDenom returns true if denom found in coins else false
func (c Coins) HasDenom(denom string) bool {
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
