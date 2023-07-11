# Starship

<p align="center" width="100%">
    <img height="148" src="https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png" />
</p>

<p align="center" width="100%">
   <a href="https://github.com/cosmology-tech/starship/releases/latest"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/release.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/blob/main/LICENSE"><img height="20" src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/e2e-tests.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/e2e-tests.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/build.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/build.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/docs.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/docs.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/starship-docker.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/starship-docker.yaml/badge.svg"></a>
</p>

Universal interchain development environment in k8s. The vision of this project
is to have a single easy to use developer environment with full testing support
for multichain use cases

## Installation
In order to get started with starship, one needs to install the following
* `kubectl`: https://kubernetes.io/docs/tasks/tools/
* `kind`: https://kind.sigs.k8s.io/docs/user/quick-start/#installation
* `helm`: https://helm.sh/docs/intro/install/
* `jq`: https://stedolan.github.io/jq/download/
* `yq`: https://github.com/mikefarah/yq/#install

## Getting started
Follow the steps here: https://starship.cosmology.tech

## Using helm chart
Inorder to use the helm chart externally without this repo.
```bash
helm repo add starship https://cosmology-tech.github.io/starship
helm repo update

helm search repo starship/devnet
```
Fetch the values.yaml file and update them before installing the chart
```bash
helm show values starship/devnet > custom-vaules.yaml
# change custom-values.yaml file

helm install -f custom-values.yaml starship/devnet --generate-name
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
