package starship

import (
	"errors"
	"github.com/cosmology-tech/starship/pkg/defaults"
	"github.com/cosmology-tech/starship/pkg/types"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

// Starship is starship config file loader, implements Loader interface
type Starship struct {
	logger *zap.Logger
}

func NewStarship(logger *zap.Logger) Starship {
	return Starship{logger: logger}
}

func (s *Starship) LoadFile(files []string, defaultConfig types.DefaultConfig) ([]types.NodeConfig, error) {
	if len(files) > 1 {
		return nil, errors.New("loading from multiple files not supported, yet")
	}
	config, err := s.loadConfig(files[0])
	if err != nil {
		return nil, err
	}
	mergedConfig, err := defaults.MergeConfigs(config, defaultConfig)

	return convertConfigToObject(mergedConfig)
}

// loadConfig reads the file into config object
func (s *Starship) loadConfig(file string) (types.Config, error) {
	config := &types.Config{}
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return types.Config{}, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return types.Config{}, err
	}

	return *config, nil
}

// convertConfigToObject converts basic Config object into Starship object, this would prefill informations
// based on chains, relayers and explorers
func convertConfigToObject(config types.Config) ([]types.NodeConfig, error) {
	nodes := []types.NodeConfig{}
	for _, cc := range config.Chains {
		chainNodes, err := convertChainToNodeConfig(*cc)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, chainNodes...)
	}

	return nodes, nil
}
