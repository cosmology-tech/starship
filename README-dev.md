# Starship

<p align="center" width="100%">
    <img height="148" src="https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png" />
</p>

<p align="center" width="100%">
   <a href="https://github.com/hyperweb-io/starship/releases/latest"><img height="20" src="https://github.com/hyperweb-io/starship/actions/workflows/release.yaml/badge.svg"></a>
   <a href="https://github.com/hyperweb-io/starship/blob/main/LICENSE"><img height="20" src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
   <a href="https://github.com/hyperweb-io/starship/actions/workflows/e2e-tests.yaml"><img height="20" src="https://github.com/hyperweb-io/starship/actions/workflows/e2e-tests.yaml/badge.svg"></a>
   <a href="https://github.com/hyperweb-io/starship/actions/workflows/build.yaml"><img height="20" src="https://github.com/hyperweb-io/starship/actions/workflows/build.yaml/badge.svg"></a>
   <a href="https://github.com/hyperweb-io/starship/actions/workflows/docs.yaml"><img height="20" src="https://github.com/hyperweb-io/starship/actions/workflows/docs.yaml/badge.svg"></a>
   <a href="https://github.com/hyperweb-io/starship/actions/workflows/starship-docker.yaml"><img height="20" src="https://github.com/hyperweb-io/starship/actions/workflows/starship-docker.yaml/badge.svg"></a>
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
Follow the steps here: https://docs.cosmology.zone/starship

## Using helm chart
In order to use the helm chart externally without this repo.
```bash
helm repo add starship https://cosmology-tech.github.io/starship
helm repo update

helm search repo starship/devnet
```
Fetch the values.yaml file and update them before installing the chart
```bash
helm show values starship/devnet > custom-values.yaml
# change custom-values.yaml file

helm install -f custom-values.yaml starship/devnet --generate-name
```

**NOTE: It is recommended to still copy the Makefile from the repo to use the handy commands**

## Related

Checkout these related projects:

* [telescope](https://github.com/hyperweb-io/telescope) Your Frontend Companion for Building with TypeScript with Cosmos SDK Modules.
* [@cosmwasm/ts-codegen](https://github.com/CosmWasm/ts-codegen) Convert your CosmWasm smart contracts into dev-friendly TypeScript classes.
* [chain-registry](https://github.com/hyperweb-io/chain-registry) Everything from token symbols, logos, and IBC denominations for all assets you want to support in your application.
* [cosmos-kit](https://github.com/hyperweb-io/cosmos-kit) Experience the convenience of connecting with a variety of web3 wallets through a single, streamlined interface.
* [create-cosmos-app](https://github.com/hyperweb-io/create-cosmos-app) Set up a modern Cosmos app by running one command.
* [interchain-ui](https://github.com/hyperweb-io/interchain-ui) The Interchain Design System, empowering developers with a flexible, easy-to-use UI kit.
* [starship](https://github.com/hyperweb-io/starship) Unified Testing and Development for the Interchain.

## Credits

🛠 Built by Hyperweb (formerly Cosmology) — if you like our tools, please checkout and contribute to [our github ⚛️](https://github.com/hyperweb-io)


## Disclaimer

AS DESCRIBED IN THE LICENSES, THE SOFTWARE IS PROVIDED “AS IS”, AT YOUR OWN RISK, AND WITHOUT WARRANTIES OF ANY KIND.

No developer or entity involved in creating this software will be liable for any claims or damages whatsoever associated with your use, inability to use, or your interaction with other users of the code, including any direct, indirect, incidental, special, exemplary, punitive or consequential damages, or loss of profits, cryptocurrencies, tokens, or anything else of value.
