package starship

import (
	"github.com/cosmology-tech/starship/pkg/types"
	"strings"
)

func getChainPorts(chainConfig types.Chain) []types.Ports {
	// default ports
	ports := []types.Ports{
		{Name: "p2p", ContainerPort: 26656, Protocol: "TCP"},
		{Name: "rpc", ContainerPort: 26657, Protocol: "TCP"},
		{Name: "address", ContainerPort: 26658, Protocol: "TCP"},
		{Name: "grpc", ContainerPort: 9090, Protocol: "TCP"},
		{Name: "rest", ContainerPort: 1317, Protocol: "TCP"},
	}

	for i, port := range ports {
		hostPort := chainConfig.Ports.GetPort(port.Name)
		if hostPort != 0 {
			ports[i].HostPort = int32(hostPort)
		}
	}

	return ports
}

// convertChainToServiceConfig creates a list of serviceConfig objects based on chain defination in config
// when a chain is converted to NodeConfig, then these containers are created:
//   - {chain}-genesis
//     init: [init-genesis, init-config]
//     sidecars: [exposer, faucet(optional)]
//   - {chain}-validator
//     dependson: genesis
//     init: [init-validator, init-config]
//     sidecars: [exposer]
//   - {chain}-cometmock (optional)
//     dependson: genesis, validator
//     init: [init-cometmock]
func convertChainToServiceConfig(chainConfig types.Chain) ([]types.NodeConfig, error) {
	allNodes := []types.NodeConfig{}

	// initialize genesis node
	genesis := types.NodeConfig{
		Name:            chainConfig.Name,
		ContainerName:   strings.Replace(chainConfig.Name, "_", "-", -1),
		Controller:      "statefulsets", // this is specific to k8s, and is ignored for others
		Image:           chainConfig.Image,
		Port:            getChainPorts(chainConfig),
		Command:         nil,
		ScriptFiles:     nil,
		WorkingDir:      "",
		Init:            nil,
		DependsOn:       nil,
		Replicas:        0,
		Labels:          nil,
		Annotations:     nil,
		Sidecars:        nil,
		Resources:       types.Resource{},
		ImagePullPolicy: "",
		Files:           nil,
	}

	allNodes = append(allNodes, genesis)

	return allNodes, nil
}
