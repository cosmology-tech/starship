package loader

import "github.com/cosmology-tech/starship/pkg/types"

type Loader interface {
	LoadFile(files []string, defaultFile string) (types.StarshipObject, error)
}
