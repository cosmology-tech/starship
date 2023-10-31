package defaults

import (
	_ "embed"
	"fmt"
	"github.com/cosmology-tech/starship/pkg/types"
	"gopkg.in/yaml.v3"
)

//go:embed defaults.yaml
var defaultFile []byte

func DefaultConfig() types.DefaultConfig {
	config := types.DefaultConfig{}

	err := yaml.Unmarshal(defaultFile, &config)
	if err != nil {
		panic(fmt.Sprintf("unable to load default config file: err: %s", err))
	}

	return config
}

func MergeConfigs(config types.Config, defaultConfig types.DefaultConfig) (types.Config, error) {
	// merge chain config
	chains := []*types.Chain{}
	for _, chain := range config.Chains {
		defaultChain, ok := defaultConfig.Chains[chain.Type]
		if !ok {
			continue
		}
		// set sidecars, scripts etc to defaultChain from defaultConfig
		defaultChain.Faucet = defaultConfig.Faucet
		defaultChain.Timeouts = defaultConfig.Timeouts
		defaultChain.Scripts = defaultConfig.Scripts
		defaultChain.Exposer = defaultConfig.Exposer

		// merge defaultChain into chain
		chains = append(chains, chain.Merge(defaultChain))
	}

	// merge relayers config
	relayers := []*types.Relayer{}
	for _, relayer := range config.Relayers {
		defaultRelayer, ok := defaultConfig.Relayers[relayer.Type]
		if !ok {
			continue
		}

		relayers = append(relayers, relayer.Merge(defaultRelayer))
	}

	return types.Config{}, nil
}
