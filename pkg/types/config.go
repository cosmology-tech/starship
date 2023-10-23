package types

import (
	"fmt"
	"reflect"
	"strings"
)

type Chain struct {
	Name          string `name:"name" json:"name,omitempty" yaml:"name"`
	Type          string `name:"type" json:"type,omitempty" yaml:"type"`
	NumValidators int    `name:"num-validators" json:"num_validators,omitempty" yaml:"numValidators"`
	Image         string `name:"image" json:"image,omitempty" yaml:"image,omitempty"`
	// Chain specifics
	Home       string `name:"home" json:"home,omitempty" yaml:"home,omitempty"`
	Binary     string `name:"binary" json:"binary,omitempty" yaml:"binary,omitempty"`
	Prefix     string `name:"prefix" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Denom      string `name:"denom" json:"denom,omitempty" yaml:"denom,omitempty"`
	PrettyName string `name:"pretty-name" json:"pretty_name,omitempty" yaml:"prettyName,omitempty"`
	Coins      string `name:"coins" json:"coins,omitempty" yaml:"coins,omitempty"`
	HDPath     string `name:"hd-path" json:"hd_path,omitempty" yaml:"hdPath,omitempty"`
	CoinType   string `name:"coin-type" json:"coin_type,omitempty" yaml:"coinType,omitempty"`
	Repo       string `name:"repo" json:"repo,omitempty" yaml:"repo,omitempty"`
	// Custom modifications
	Scripts  map[string]ScriptData  `name:"scripts" json:"scripts,omitempty" yaml:"scripts"`
	Upgrade  Upgrade                `name:"upgrade" json:"upgrade,omitempty" yaml:"upgrade"`
	Genesis  map[string]interface{} `name:"genesis" json:"genesis,omitempty" yaml:"genesis"`
	Timeouts map[string]string      `name:"timeouts" json:"timeouts,omitempty" yaml:"timeouts"`
	// Feature toggles
	Build     Build    `name:"build" json:"build,omitempty" yaml:"build,omitempty"`
	Cometmock *Feature `name:"cometmock" json:"cometmock,omitempty" yaml:"cometmock,omitempty"`
	Faucet    *Feature `name:"facuet" json:"faucet,omitempty" yaml:"faucet,omitempty"`
	ICS       *Feature `name:"ics" json:"ics,omitempty" yaml:"ics,omitempty"`
	// Additional information
	Ports     HostPort `name:"ports" json:"ports,omitempty" yaml:"ports,omitempty"`
	Resources Resource `name:"resource" json:"resources,omitempty" yaml:"resources,omitempty"`
}

func (c *Chain) GetName() string {
	return strings.Replace(c.Name, "_", "-", -1)
}

func (c *Chain) GetChainID() string {
	return c.Name
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

type Build struct {
	Enabled bool   `name:"enabled" json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Source  string `name:"source" json:"source,omitempty" yaml:"source,omitempty"`
}

type HostPort struct {
	Rest    int `name:"rest" json:"rest" yaml:"rest"`
	Rpc     int `name:"rpc" json:"rpc" yaml:"rpc"`
	Grpc    int `name:"grpc" json:"grpc" yaml:"grpc"`
	Exposer int `name:"exposer" json:"exposer" yaml:"exposer"`
	Faucet  int `name:"faucet" json:"faucet" yaml:"faucet"`
}

func (p HostPort) GetPort(port string) int {
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Int()
		fieldName := field.Tag.Get("name")

		if fieldName == port {
			return int(value)
		}
	}
	return 0
}

type Relayer struct {
	Name     string      `name:"name" json:"name" yaml:"name"`
	Type     string      `name:"type" json:"type" yaml:"type"`
	Image    string      `name:"image" json:"image,omitempty" yaml:"image,omitempty"`
	Replicas int         `name:"replicas" json:"replicas" yaml:"replicas"`
	Chains   []string    `name:"chains" json:"chains" yaml:"chains"`
	Config   interface{} `name:"config" json:"config,omitempty" yaml:"config,omitempty"`
}

type Faucet struct {
	Enabled     bool     `name:"enabled" json:"enabled" yaml:"enabled"`
	Image       string   `name:"image" json:"image" yaml:"image"`
	Ports       HostPort `name:"ports" json:"ports" yaml:"ports"`
	Concurrency int      `name:"concurrency" json:"concurrency,omitempty" yaml:"concurrency,omitempty"`
}

type Feature struct {
	Enabled bool     `name:"enabled" json:"enabled" yaml:"enabled"`
	Image   string   `name:"image" json:"image" yaml:"image"`
	Ports   HostPort `name:"ports" json:"ports" yaml:"ports"`
}

func (f *Feature) GetRPCAddr() string {
	return fmt.Sprintf("http://localhost:%d", f.Ports.Rpc)
}

func (f *Feature) GetRESTAddr() string {
	return fmt.Sprintf("http://localhost:%d", f.Ports.Rest)
}

// Config is the struct for the config.yaml setup file
// Need not be fully compatible with the values.schema.json file, just need
// parts of the config file for performing  various functions, mainly port-forwarding
// todo: move this to a more common place, outside just tests
// todo: can be moved to proto defination
type Config struct {
	Chains   []*Chain   `name:"chains" json:"chains" yaml:"chains"`
	Relayers []*Relayer `name:"relayers" json:"relayers" yaml:"relayers"`
	Explorer *Feature   `name:"explorer" json:"explorer" yaml:"explorer"`
	Registry *Feature   `name:"registry" json:"registry" yaml:"registry"`
}

// HasChainId returns true if chain id found in list of chains
func (c *Config) HasChainId(chainId string) bool {
	for _, chain := range c.Chains {
		if chain.Name == chainId {
			return true
		}
	}

	return false
}

// GetChain returns the Chain object pointer for the given chain id
func (c *Config) GetChain(chainId string) *Chain {
	for _, chain := range c.Chains {
		if chain.Name == chainId {
			return chain
		}
	}

	return nil
}
