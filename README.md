# Starship

<p align="center" width="100%">
    <img height="148" src="https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png" />
</p>

<p align="center" width="100%">
   <a href="https://github.com/cosmology-tech/starship/blob/main/LICENSE"><img height="20" src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/releases/latest"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/release.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/build.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/build.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/e2e-tests.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/e2e-tests.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/docs.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/docs.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/starship-docker.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/starship-docker.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/run-client-tests.yml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/run-client-tests.yml/badge.svg" /></a>
</p>

Universal interchain development environment in k8s. The vision of this project
is to have a single easy to use developer environment with full testing support
for multichain use cases

## Installation
In order to get started with starship, one needs to install the following
* `kubectl`: https://kubernetes.io/docs/tasks/tools/ (you can use [Docker Desktop](https://www.docker.com/products/docker-desktop/) for simple install)
* `helm`: https://helm.sh/docs/intro/install/

## Install

Install the test utilities `starshipjs` and the CLI `@starship-ci/cli`:

```sh
yarn add --dev starshipjs @starship-ci/cli
```

### Recommended Usage 📘

Stay tuned for a `create-cosmos-app` boilerplate! For now, this is the most recommended setup. Consider everything else after this section "advanced setup".

- We recommend studying the [osmojs starship integration](https://github.com/osmosis-labs/osmojs/tree/main/packages/osmojs/starship) and replicating it.
- Add your configs, similar to how it's done [here](https://github.com/osmosis-labs/osmojs/tree/main/packages/osmojs/starship/configs)
- Add your workflows for GitHub Actions [like this](https://github.com/osmosis-labs/osmojs/blob/main/.github/workflows/e2e-tests.yaml)
- Add `yarn starship` commands to your package.json scripts [like this](https://github.com/osmosis-labs/osmojs/blob/c456184666eda55cd6fee5cd09ba6c05c898d55c/packages/osmojs/package.json#L31-L34)
— Note the jest configurations in the [osmojs package](https://github.com/osmosis-labs/osmojs/tree/main/packages/osmojs)


This will allow you to run `yarn starship` to `start`, `setup`, `deploy`, `stop` and other `starship` commands:

#### Deploying `Starship` 🚀

```sh
yarn starship start
```

#### Running End-to-End Tests 🧪

```sh
# test
yarn starship:test

# watch 
yarn starship:watch
```

#### Teardown 🛠️

```sh
# stop ports and delete deployment
yarn starship stop
```

## Related

Checkout these related projects:

* [@cosmology/telescope](https://github.com/cosmology-tech/telescope) Your Frontend Companion for Building with TypeScript with Cosmos SDK Modules.
* [@cosmwasm/ts-codegen](https://github.com/CosmWasm/ts-codegen) Convert your CosmWasm smart contracts into dev-friendly TypeScript classes.
* [chain-registry](https://github.com/cosmology-tech/chain-registry) Everything from token symbols, logos, and IBC denominations for all assets you want to support in your application.
* [cosmos-kit](https://github.com/cosmology-tech/cosmos-kit) Experience the convenience of connecting with a variety of web3 wallets through a single, streamlined interface.
* [create-cosmos-app](https://github.com/cosmology-tech/create-cosmos-app) Set up a modern Cosmos app by running one command.
* [interchain-ui](https://github.com/cosmology-tech/interchain-ui) The Interchain Design System, empowering developers with a flexible, easy-to-use UI kit.
* [starship](https://github.com/cosmology-tech/starship) Unified Testing and Development for the Interchain.

## Credits

🛠 Built by Cosmology — if you like our tools, please consider delegating to [our validator ⚛️](https://cosmology.zone/validator)


## Disclaimer

AS DESCRIBED IN THE LICENSES, THE SOFTWARE IS PROVIDED “AS IS”, AT YOUR OWN RISK, AND WITHOUT WARRANTIES OF ANY KIND.

No developer or entity involved in creating this software will be liable for any claims or damages whatsoever associated with your use, inability to use, or your interaction with other users of the code, including any direct, indirect, incidental, special, exemplary, punitive or consequential damages, or loss of profits, cryptocurrencies, tokens, or anything else of value.
