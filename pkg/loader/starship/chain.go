package starship

import (
	"fmt"
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

func chainEnvVars(chainConfig types.Chain) []types.EnvVar {
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
	if chainConfig.Faucet != nil && chainConfig.Faucet.Enabled {
		faucetEnv.Value = "true"
	}
	envs = append(envs, faucetEnv)

	// add timeout envrionment vars
	for key, value := range chainConfig.Timeouts {
		envs = append(envs, types.EnvVar{strings.ToUpper(key), value})
	}

	return envs
}

// genesisInits returns object Init that indicates scripts files for
// init containers for genesis chain node
// genesisInit: [init-genesis, init-config]
func genesisInits(chainConfig types.Chain) ([]types.Init, error) {
	inits := []types.Init{}

	envs := chainEnvVars(chainConfig)

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
	if chainConfig.Faucet != nil && chainConfig.Faucet.Enabled {
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

// validatorInits returns all the init containers for the validator node
// todo: each validator node down the line needs to be injected with following
// todo: envrionment variables: NODE_INDEX (corresponding to replica number), GENESIS_HOST (host of genesis to connect to)
// todo: GENESIS_PORT_* (all ports of the genesis node that are exposed) which to connect to. The initialization scripts will use
// todo: use them as if they are present, and then they are injected later
func validatorInits(cc types.Chain) ([]types.Init, error) {
	inits := []types.Init{}

	envs := chainEnvVars(cc)
	// actual values of NodeIndex is injected in the transformers, since it
	// depends on the underlying implmentation
	envs = append(envs, types.EnvVar{"NODE_INDEX", "0"})

	command := `|
VAL_INDEX=${NODE_INDEX}
echo "Validator Index: $VAL_INDEX"
if [ -f $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN ]; then
	cp $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN /usr/bin
fi

if [ -f $CHAIN_DIR/config/genesis.json ]; then
	echo "Genesis file exists, exiting early"
	exit 0
fi

VAL_NAME=$(jq -r ".validators[0].name" $KEYS_CONFIG)-$VAL_INDEX
echo "Validator Index: $VAL_INDEX, Key name: $VAL_NAME"

echo "Recover validator $VAL_NAME"
$CHAIN_BIN init $VAL_NAME --chain-id $CHAIN_ID
jq -r ".validators[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN keys add $VAL_NAME --index $VAL_INDEX --recover --keyring-backend="test"

curl http://$GENESIS_HOST:$GENESIS_PORT_EXPOSER/genesis -o $CHAIN_DIR/config/genesis.json
echo "Genesis file that we got....."
cat $CHAIN_DIR/config/genesis.json

echo "Create node id json file"
NODE_ID=$($CHAIN_BIN tendermint show-node-id)
echo '{"node_id":"'$NODE_ID'"}' > $CHAIN_DIR/config/node_id.json
`

	validatorInit := types.Init{
		Name:        "init-validator",
		Image:       cc.Image,
		Command:     []string{"bash", "-c", command},
		Environment: envs,
	}
	inits = append(inits, validatorInit)

	// init-config container
	// Note: NODE_INDEX is an envrionment variable that is injected into
	// the system in the transformer, but the initialization scripts can use it.
	initCommand := `|
VAL_INDEX=${NODE_INDEX}
echo "Validator Index: $VAL_INDEX"
if [ -f $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN ]; then
	cp $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN /usr/bin
fi

echo "Running setup config script..."
bash -e /scripts/update-config.sh
`
	configInit := types.Init{
		Name:        "init-config",
		Image:       cc.Image,
		Environment: envs,
		Command:     []string{"bash", "-c", initCommand},
	}
	inits = append(inits, configInit)

	return inits, nil
}

func createExposerSidecar(chainConfig types.Chain) (types.NodeConfig, error) {
	exposer := types.NodeConfig{
		Name:          "exposer",
		ContainerName: "exposer",
		Image:         "ghcr.io/cosmology-tech/starship/exposer:20231011-cdbd60b",
		Environment: []types.EnvVar{
			{
				Name:  "EXPOSER_HTTP_PORT",
				Value: "8081",
			},
			{
				Name:  "EXPOSER_GRPC_PORT",
				Value: "9099",
			},
			{
				Name:  "EXPOSER_GENESIS_FILE",
				Value: fmt.Sprintf("%s/config/genesis.json", chainConfig.Home),
			},
			{
				Name:  "EXPOSER_MNEMONIC_FILE",
				Value: "/configs/keys.json",
			},
			{
				Name:  "EXPOSER_PRIV_VAL_FILE",
				Value: fmt.Sprintf("%s/config/priv_validator_key.json", chainConfig.Home),
			},
			{
				Name:  "EXPOSER_NODE_KEY_FILE",
				Value: fmt.Sprintf("%s/config/node_key.json", chainConfig.Home),
			},
			{
				Name:  "EXPOSER_NODE_ID_FILE",
				Value: fmt.Sprintf("%s/config/node_id.json", chainConfig.Home),
			},
			{
				Name:  "EXPOSER_PRIV_VAL_STATE_FILE",
				Value: fmt.Sprintf("%s/data/priv_validator_state.json", chainConfig.Home),
			},
		},
		Command:  []string{"exposer"},
		Replicas: 1,
		Resources: types.Resource{
			CPU:    "100m",
			Memory: "100Mi",
		},
		ImagePullPolicy: "IfNotPresent",
	}

	return exposer, nil
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
func convertChainToNodeConfig(chainConfig types.Chain) ([]types.NodeConfig, error) {
	allNodes := []types.NodeConfig{}

	mounts := []types.Mount{
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
	}

	inits, err := genesisInits(chainConfig)
	if err != nil {
		return nil, err
	}

	exposer, err := createExposerSidecar(chainConfig)
	if err != nil {
		return nil, err
	}

	envs := chainEnvVars(chainConfig)

	// initialize genesis node
	genesis := types.NodeConfig{
		Name:          fmt.Sprintf("%s-genesis", chainConfig.Name),
		ContainerName: strings.Replace(fmt.Sprintf("%s-genesis", chainConfig.Name), "_", "-", -1),
		Controller:    "statefulset", // this is specific to k8s, and is ignored for others
		Image:         chainConfig.Image,
		Port:          getChainPorts(chainConfig.Ports),
		Environment:   envs,
		Command:       nil,
		ScriptFiles:   nil,
		WorkingDir:    "",
		Init:          inits,
		DependsOn:     nil,
		Replicas:      1,
		Labels: map[string]string{
			"instance": chainConfig.GetName(),
			"type":     chainConfig.Type,
			"rawname":  chainConfig.Name,
		},
		Annotations: map[string]string{
			"quality": "release",
			"role":    "node",
			"sla":     "medium",
			"tier":    "genesis",
		},
		Sidecars:        []*types.NodeConfig{&exposer},
		Resources:       chainConfig.Resources,
		ImagePullPolicy: "IfNotPresent",
		Mounts:          mounts,
	}
	allNodes = append(allNodes, genesis)

	// return early incase of no validators
	if chainConfig.NumValidators <= 1 {
		return allNodes, nil
	}

	validatorInits, err := validatorInits(chainConfig)
	if err != nil {
		return nil, err
	}

	validatorExposer, err := createExposerSidecar(chainConfig)
	if err != nil {
		return nil, err
	}

	validatorCmd := `
set -eux
START_ARGS=""

# Starting the chain
if [ -f $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN ]; then
	cp $CHAIN_DIR/cosmovisor/genesis/bin/$CHAIN_BIN /usr/bin
	/usr/bin/cosmovisor start $START_ARGS
else
	$CHAIN_BIN start $START_ARGS
fi
`
	postStartCmd := `
set -eux
VAL_INDEX=${NODE_INDEX}
VAL_NAME="$(jq -r ".validators[0].name" $KEYS_CONFIG)-$VAL_INDEX"
echo "Validator Index: $VAL_INDEX, Key name: $VAL_NAME. Chain bin $CHAIN_BIN"

VAL_ADDR=$($CHAIN_BIN keys show $VAL_NAME -a --keyring-backend="test")
echo "Transfer tokens to address $VAL_ADDR before trying to create validator. Best effort"
bash -e /scripts/transfer-tokens.sh \
	$VAL_ADDR \
	$DENOM \
	http://$GENESIS_HOST:$GENESIS_PORT_FAUCET/credit \
	"$FAUCET_ENABLED" || true

$CHAIN_BIN keys list --keyring-backend test | jq
VAL_NAME=$VAL_NAME bash -e /scripts/create-validator.sh
`

	validatorNodes := types.NodeConfig{
		Name:            fmt.Sprintf("%s-validator", chainConfig.Name),
		ContainerName:   strings.Replace(fmt.Sprintf("%s-validator", chainConfig.Name), "_", "-", -1),
		Controller:      "statefulset",
		Image:           chainConfig.Image,
		Environment:     envs,
		Port:            getChainPorts(chainConfig.Ports),
		Command:         []string{"bash", "\"-c\"", validatorCmd},
		Init:            validatorInits,
		DependsOn:       []*types.NodeConfig{&genesis},
		Replicas:        chainConfig.NumValidators - 1,
		Sidecars:        []*types.NodeConfig{&validatorExposer},
		Resources:       chainConfig.Resources,
		ImagePullPolicy: "IfNotPresent",
		Mounts:          mounts,
		Labels: map[string]string{
			"instance": chainConfig.GetName(),
			"type":     chainConfig.Type,
			"rawname":  chainConfig.Name,
		},
		Annotations: map[string]string{
			"quality": "release",
			"role":    "node",
			"sla":     "medium",
			"tier":    "validator",
		},
		PostStart: []string{"bash", "\"-c\"", "\"-e\"", postStartCmd},
	}
	allNodes = append(allNodes, validatorNodes)

	return allNodes, nil
}
