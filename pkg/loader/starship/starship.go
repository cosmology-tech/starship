package starship

import (
	"errors"
	"github.com/cosmology-tech/starship/pkg/types"
	"gopkg.in/yaml.v3"
	"os"
)

// Starship is starship config file loader, implements Loader interface
type Starship struct {
}

func (s *Starship) LoadFile(files []string, defaultFile string) (types.StarshipObject, error) {
	if len(files) > 0 {
		return types.StarshipObject{}, errors.New("loading from multiple files not supported, yet")
	}
	config, err := s.loadConfig(files[0])
	if err != nil {
		return types.StarshipObject{}, nil
	}
	// todo: override defaults into config

	return convertConfigToObject(config)
}

// loadConfig reads the file into config object
func (s *Starship) loadConfig(file string) (types.Config, error) {
	config := types.Config{}
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return types.Config{}, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return types.Config{}, err
	}

	return config, nil
}

// convertConfigToObject converts basic Config object into Starship object, this would prefill informations
// based on chains, relayers and explorers
func convertConfigToObject(config types.Config) (types.StarshipObject, error) {
	return types.StarshipObject{}, nil
}
