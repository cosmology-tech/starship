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
		// set default resources for the chain from resources.node
		defaultChain.Resources = defaultConfig.Node.Resources

		// merge defaultChain into chain
		mc, err := chain.Merge(defaultChain)
		if err != nil {
			return types.Config{}, err
		}
		chains = append(chains, mc)
	}

	// merge relayers config
	relayers := []*types.Relayer{}
	for _, relayer := range config.Relayers {
		defaultRelayer, ok := defaultConfig.Relayers[relayer.Type]
		if !ok {
			continue
		}
		defaultRelayer.Resources = defaultConfig.Node.Resources

		relayers = append(relayers, relayer.Merge(defaultRelayer))
	}

	mergedConfig := types.Config{
		Chains:   chains,
		Relayers: relayers,
	}

	if config.Registry != nil {
		merged, err := config.Registry.Merge(defaultConfig.Registry)
		if err != nil {
			return types.Config{}, err
		}
		mergedConfig.Registry = merged
	}

	if config.Explorer != nil {
		merged, err := config.Explorer.Merge(defaultConfig.Explorer)
		if err != nil {
			return types.Config{}, err
		}
		mergedConfig.Explorer = merged
	}

	if config.Monitoring != nil {
		merged, err := config.Monitoring.Merge(defaultConfig.Monitoring)
		if err != nil {
			return types.Config{}, err
		}
		mergedConfig.Monitoring = merged
	}

	return mergedConfig, nil
}
