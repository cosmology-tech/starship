package kubernetes

import "github.com/cosmology-tech/starship/pkg/types"

// Kubernetes implements Transformer interface and represents Kubernetes transformer
type Kubernetes struct {
	// the user provided options from the command line
	Opt types.ConvertOptions
}
