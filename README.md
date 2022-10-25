# shuttle
Universal interchain development environment in k8s. The vision of this project
is to have a single easy to use developer environment with full testing support
for multichain use cases

## Installation
Inorder to get started with shuttle, one needs to install the following
* `kubectl`: https://kubernetes.io/docs/tasks/tools/
* `kind`: https://kind.sigs.k8s.io/docs/user/quick-start/#installation
* `helm`: https://helm.sh/docs/intro/install/
* `jq`: https://stedolan.github.io/jq/download/
* `yq`: https://github.com/mikefarah/yq/#install

## Getting started
Update the `vaules.yaml` in `charts/devnet/templates`. Recommeded one creates a copy
of the values file and update it as per your requirements.

### Setup local k8s cluster (optional)
Create a local k8s cluster using `kind`. Can be setup with `minikube` or
docker-desktop as well.
One can use the handy make commands in the `Makefile` like following
```bash
make setup-kind
```
This will create a local kind cluster in docker and set the correct context in
your current kubectl. Check the kubectl context with
```bash
kubectl config current-context
# check: kind-shuttle
```

### Setup k8s cluster
1. Connect to a k8s cluster, make sure you are able to access following command
   ```bash
   kubectl get pods
   ```
2. Create a namespace in which the setup will be deployed.
   ```bash
   kubectl create namespace <namespace-name>
   # example
   kubectl create namespace shuttle
   ```
3. Make sure you have set the namespace in the current context, so the devnet is deployed
   without conflict to your current workloads

### Start
1. Debug the k8s yaml configuration files
   ```bash
   make debug VALUES_NAME=<custom-filename>
   # output all yaml files that will be deployed 
   # default values file run
   make debug
   ```
2. Start the cluster
   ```bash
   make install VALUES_NAME=<custom-filename>
   # default values file run
   make install
   ```
   Optionally you can use k9s, to watch all the fun
3. Once you make any changes to the system or values, run
   ```bash
   make upgrade VALUES_NAME=<custom-filename>
   # default values file run
   make upgrade
   ```
4. Run local port forwarding based on local port info in the `values.yaml`
   ```bash
   # port-forward all the local ports locally, runs in background
   make port-forward-all
   
   # Run following to stop port forwarding once done
   make stop-forward
   ```
   Sometime one might need to run connection updates so the port-forward does not
   get timed out. Run `make check-forward-all`
5. Open the explorer at `http://localhost:8080`
6. To clean up everything run
   ```bash
   # Kill any portforwarding
   make stop-forward
   # Delete current helm chart deployment
   make delete
   # If running local kind cluster, cleanup
   make clean-kind
   ```

## Custom setup
When one wants to change the configuration for setting up the devnet.
1. Copy the `charts/devnet/values.yaml` to `custom-values.yaml` file.
2. Update the yaml file
3. Install chart `make install VALUES_FILE=custom-values.yaml`

While making changes to the `vaules.yaml` file here are the modifications one
can perform
* Add chains to setup at `.chains` key, copy and paste an existing chain and make changes
* Update how many validators per chain at `.chains[].numValidators`
* Add relayer between chains with adding dict to `.relayers`, mention the chains to
  connect between in `.relayers[].chains`, use the chain name defined in `.chains`
* Toggle explorer with `.explorer.enabled` boolen flag

**NOTE: `values.yaml` still needs to be optimized with default values and less user inputs**

### ProTip: Local setup
* Set `.chains[].numValidators` to `1` when using local cluster
* Comment out chains and relayers not needed
* Disable explorer for local setup

## Using helm chart
Inorder to use the helm chart externally without this repo.
```bash
helm repo add shuttle https://anmol1696.github.io/shuttle
helm repo update

helm search repo shuttle/devnet
```
Fetch the values.yaml file and update them before installing the chart
```bash
helm show values shuttle/devnet > custom-vaules.yaml
# change custom-values.yaml file

helm install -f custom-values.yaml shuttle/devnet --generate-name
```

**NOTE: It is recommended to still copy the Makefile from the repo to use the handy commands**

# Future works
Some features that are to be added
* Default values for each chain, addition of `type` for chains
* Move scripts directly to configmaps
* Add support for hermes and go relayers
* How to perform ingress for local or cluster to get traffic to the system, this will
  remove the need for any sort of local port forwarding

## Improvements
* Make faster kye initialization, currently most time at startup is taken by adding
  keys to keyring from mnemonics
* Add comments to `values.yaml`
* Currently, the setup runs well for 12hrs before hitting the memory limits and the
  storage limits. Need to set proper pruning setup for optimal devnet

## Major design considerations
### GH Actions
Run this in CI/CD.

### E2E tests
The actual tests we can run.

### Docker images
Currently, for each chain, the docker image needs to be built stored in `docker/`
directory. Most of the docker images just need `jq`, `bash` and `sed`.
Can look into using strangelove-ventures/heighliner for docker images creation
instead of having self-hosted

### Chain binary
The docker images also require the chain binary to be already part of the container.
Need the ability for users to either pass the binary as input, or to build binary on
the fly, so that we can use a standard base docker image, and add binaries on top.

### Mainnet exports
Ability to get an exported state on a mainnet/testnet for a chain, and update locally.
LocalOsmosis has some functions and scripts to perform the state change.

Ideally we will spin up a new type of container, fetch the data, perform transformations 
and spread it out to all the following validator nodes.

### Upgrade testing: ToDo
* Run chain in current state, need to install cosmovisor on all nodes
  * Might neeed to binary pre-installed for both upgrades
  * Can fetch binary via a wget in init-containers
  * Can build the code in init container as well
* run some pre-upgrade txns state
* create upgrade proposal
* cosmovisor will do the upgrade automatically
* After upgrade, we loose the touch with ability to update `genesis.json`
  directly, 

### Bazel: Servicify
* Serviceify this, blaze tests, run all the test for any language
  * Bazal integrates with remote test runner
* Service running that listens to commands, and run commands on the cluster
  * Test RPC: takes test request and spits out an output
  * Create a web frontend that displays your results
* If we provide any sort of devnet service, we need to figure out **Authn** for all the
  calls, knowing that for local debugging there is no authn for actual blockchains.
  Port-forwarding does not sound like a bad idea, since the burden of authn is on
  `kubectl`.

### Stress testing and profiling
Inorder to understand the topology and deeper insights into how the system works, it
would make sense to try and use a simulated environment and perform stress tests as
a blackbox testing.
We can create new binaries with pprof or profiling tool embedded, run test workflow
and look at the analytics.

Currently, the way to do this, is very manual and some tools are being built for local
profiling, but large scale profiling still needs to happen.
