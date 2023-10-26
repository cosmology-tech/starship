package starship

import (
	"github.com/cosmology-tech/starship/pkg/types"
	"strconv"
	"strings"
)

func getChainPorts(hostPorts types.HostPort) []types.Ports {
	// default ports
	ports := []types.Ports{
		{Name: "p2p", ContainerPort: 26656, Protocol: "TCP"},
		{Name: "rpc", ContainerPort: 26657, Protocol: "TCP"},
		{Name: "address", ContainerPort: 26658, Protocol: "TCP"},
		{Name: "grpc", ContainerPort: 9090, Protocol: "TCP"},
		{Name: "rest", ContainerPort: 1317, Protocol: "TCP"},
	}

	for i, port := range ports {
		hostPort := hostPorts.GetPort(port.Name)
		if hostPort != 0 {
			ports[i].HostPort = int32(hostPort)
		}
	}

	return ports
}

// genesisInits returns object Init that indicates scripts files for
// init containers for genesis chain node
// genesisInit: [init-genesis, init-config]
func genesisInits(chainConfig types.Chain) ([]types.Init, error) {
	inits := []types.Init{}

	envs := []types.EnvVar{
		{"KEYS_CONFIG", "/configs/keys.json"},
		{"NUM_VALIDATORS", strconv.Itoa(chainConfig.NumValidators)},
		{"DENOM", chainConfig.Denom},
		{"COINS", chainConfig.Coins},
		{"CHAIN_DIR", chainConfig.Home},
		{"CODE_REPO", chainConfig.Repo},
		{"DAEMON_HOME", chainConfig.Home},
		{"DAEMON_NAME", chainConfig.Binary},
		{"CHAIN_ID", chainConfig.GetChainID()},
	}

	buildEnv := types.EnvVar{"TO_BUILD", "false"}
	if chainConfig.Build.Enabled {
		buildEnv.Value = "true"
	}
	envs = append(envs, buildEnv)

	// add faucet environment vars
	faucetEnv := types.EnvVar{"FUACET_ENABLED", "false"}
	if chainConfig.Faucet.Enabled {
		faucetEnv.Value = "true"
	}
	envs = append(envs, faucetEnv)

	// add timeout envrionment vars
	for key, value := range chainConfig.Timeouts {
		envs = append(envs, types.EnvVar{strings.ToUpper(key), value})
	}

	command := `
		VAL_INDEX=0
		echo "Validator Index: $VAL_INDEX"
		if [ -f $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN ]; then
			cp $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN /usr/bin
		fi
	
		if [ -f $CHAIN_DIR/config/genesis.json ]; then
		  echo "Genesis file exists, exiting init container"
		  exit 0
		fi
	
		echo "Running setup genesis script..."
		bash -e /scripts/create-genesis.sh
		bash -e /scripts/update-genesis.sh
	
		echo "Create node id json file"
		NODE_ID=$($CHAIN_BIN tendermint show-node-id)
		echo '{"node_id":"'$NODE_ID'"}' > $CHAIN_DIR/config/node_id.json
`
	mounts := []types.Mount{
		{
			Name:      "addresses",
			MountPath: "/configs",
			Path:      "./configs",
		},
		{
			Name:      "scripts",
			MountPath: "/scripts",
			Path:      "./scripts",
		},
	}

	// deal with overwriting genesis config map
	// todo: need to implement this this
	if chainConfig.Genesis != nil {
	}

	genesisInit := types.Init{
		Name:        "init-genesis",
		Image:       chainConfig.Image,
		Command:     []string{"bash", "-c", command},
		Environment: envs,
		Mounts:      mounts,
	}
	inits = append(inits, genesisInit)

	// init-config container
	initCommand := `
		VAL_INDEX=${HOSTNAME##*-}
		echo "Validator Index: $VAL_INDEX"
		if [ -f $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN ]; then
			cp $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN /usr/bin
		fi

		echo "Running setup config script..."
		bash -e /scripts/update-config.sh
`
	configInit := types.Init{
		Name:        "init-config",
		Image:       chainConfig.Image,
		Environment: envs,
		Command:     []string{"bash", "-c", initCommand},
		Mounts: []types.Mount{
			{
				Name:      "configs",
				MountPath: "/configs",
				Path:      "./configs",
			},
			{
				Name:      "scripts",
				MountPath: "/scripts",
				Path:      "./scripts",
			},
			{
				Name:      "node",
				MountPath: chainConfig.Home,
			},
		},
	}
	inits = append(inits, configInit)

	// faucet init
	if chainConfig.Faucet.Enabled {
		faucetCmd := `
			cp /bin/faucet /faucet/faucet
			chmod +x /faucet/faucet
`
		faucetInit := types.Init{
			Name:       "init-faucet",
			Image:      chainConfig.Faucet.Image,
			Command:    []string{"bash", "-c", faucetCmd},
			WorkingDir: "",
			Mounts: []types.Mount{
				{
					Name:      "faucet",
					MountPath: "/faucet",
				},
			},
		}
		inits = append(inits, faucetInit)
	}

	return inits, nil
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

	inits, err := genesisInits(chainConfig)
	if err != nil {
		return nil, err
	}
	// initialize genesis node
	genesis := types.NodeConfig{
		Name:            chainConfig.Name,
		ContainerName:   strings.Replace(chainConfig.Name, "_", "-", -1),
		Controller:      "statefulsets", // this is specific to k8s, and is ignored for others
		Image:           chainConfig.Image,
		Port:            getChainPorts(chainConfig.Ports),
		Command:         nil,
		ScriptFiles:     nil,
		WorkingDir:      "",
		Init:            inits,
		DependsOn:       nil,
		Replicas:        0,
		Labels:          nil,
		Annotations:     nil,
		Sidecars:        nil,
		Resources:       chainConfig.Resources,
		ImagePullPolicy: "",
		Files:           nil,
	}

	allNodes = append(allNodes, genesis)

	return allNodes, nil
}
