package loader

import "github.com/cosmology-tech/starship/pkg/types"

type Loader interface {
	LoadFile(files []string, defaultConfig types.DefaultConfig) (types.StarshipObject, error)
}
