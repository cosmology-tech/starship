package main

import (
	"fmt"
	"strings"
)

type Chain struct {
	Name          string                 `name:"name" json:"name,omitempty" yaml:"name"`
	Type          string                 `name:"type" json:"type,omitempty" yaml:"type"`
	Image         string                 `name:"image" json:"image,omitempty" yaml:"image"`
	NumValidators int                    `name:"num-validators" json:"num_validators,omitempty" yaml:"numValidators"`
	Scripts       map[string]ScriptData  `name:"scripts" json:"scripts,omitempty" yaml:"scripts"`
	Ports         Port                   `name:"ports" json:"ports,omitempty" yaml:"ports"`
	Upgrade       Upgrade                `name:"upgrade" json:"upgrade,omitempty" yaml:"upgrade"`
	Genesis       map[string]interface{} `name:"genesis" json:"genesis,omitempty" yaml:"genesis"`
}

func (c *Chain) GetName() string {
	return strings.Replace(c.Name, "_", "-", -1)
}

func (c *Chain) GetRPCAddr() string {
	return fmt.Sprintf("http://localhost:%d", c.Ports.Rpc)
}

func (c *Chain) GetRESTAddr() string {
	return fmt.Sprintf("http://localhost:%d", c.Ports.Rest)
}

type ScriptData struct {
	File string `name:"file" json:"file,omitempty" yaml:"file"`
	Data string `name:"data" json:"data,omitempty" yaml:"data"`
}

type Upgrade struct {
	Enabled  bool   `name:"eanbled" json:"enabled" yaml:"enabled"`
	Type     string `name:"type" json:"type" yaml:"type"`
	Genesis  string `name:"genesis" json:"genesis" yaml:"genesis"`
	Upgrades []struct {
		Name    string `name:"name" json:"name" yaml:"name"`
		Version string `name:"version" json:"version" yaml:"version"`
	} `name:"upgrades" json:"upgrades" yaml:"upgrades"`
}

type Port struct {
	Rest    int `name:"rest" json:"rest" yaml:"rest"`
	Rpc     int `name:"rpc" json:"rpc" yaml:"rpc"`
	Grpc    int `name:"grpc" json:"grpc" yaml:"grpc"`
	Exposer int `name:"exposer" json:"exposer" yaml:"exposer"`
	Faucet  int `name:"faucet" json:"faucet" yaml:"faucet"`
}

func (p Port) GetPort(port string) int {
	switch port {
	case "rpc":
		return p.Rpc
	case "rest":
		return p.Rest
	case "grpc":
		return p.Grpc
	case "exposer":
		return p.Exposer
	case "faucet":
		return p.Faucet
	default:
		return 0
	}
}

type Relayer struct {
	Name     string   `name:"name" json:"name" yaml:"name"`
	Type     string   `name:"type" json:"type" yaml:"type"`
	Image    string   `name:"image" json:"image" yaml:"image"`
	Replicas int      `name:"replicas" json:"replicas" yaml:"replicas"`
	Chains   []string `name:"chains" json:"chains" yaml:"chains"`
}

type Feature struct {
	Enabled bool   `name:"enabled" json:"enabled" yaml:"enabled"`
	Image   string `name:"image" json:"image" yaml:"image"`
	Ports   Port   `name:"ports" json:"ports" yaml:"ports"`
}

func (f *Feature) GetRPCAddr() string {
	return fmt.Sprintf("http://localhost:%d", f.Ports.Rpc)
}

func (f *Feature) GetRESTAddr() string {
	return fmt.Sprintf("http://localhost:%d", f.Ports.Rest)
}

// HelmConfig is the struct for the config.yaml setup file
// Need not be fully compatible with the values.schema.json file, just need
// parts of the config file for performing  various functions, mainly port-forwarding
// todo: move this to a more common place, outside just tests
// todo: can be moved to proto defination
type HelmConfig struct {
	Chains   []*Chain   `name:"chains" json:"chains" yaml:"chains"`
	Relayers []*Relayer `name:"relayers" json:"relayers" yaml:"relayers"`
	Explorer *Feature   `name:"explorer" json:"explorer" yaml:"explorer"`
	Registry *Feature   `name:"registry" json:"registry" yaml:"registry"`
}

// HasChainId returns true if chain id found in list of chains
func (c *HelmConfig) HasChainId(chainId string) bool {
	for _, chain := range c.Chains {
		if chain.Name == chainId {
			return true
		}
	}

	return false
}

// GetChain returns the Chain object pointer for the given chain id
func (c *HelmConfig) GetChain(chainId string) *Chain {
	for _, chain := range c.Chains {
		if chain.Name == chainId {
			return chain
		}
	}

	return nil
}
