package transformer

import (
	"github.com/cosmology-tech/starship/pkg/types"
)

type Objects interface {
	WriteToFile(dir string) error
	Validate() error
}

type Transformer interface {
	Transform([]types.NodeConfig, types.ConvertOptions) (Objects, error)
}
