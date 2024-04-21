# Starship CLI

<p align="center">
  <img src="https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png" width="80"><br />
  StarshipJS enables developers to efficiently set up and test chains, explorers, and validators, making it easier to handle development projects spanning several blockchain networks.
</p>

<p align="center" width="100%">
  <a href="https://github.com/cosmology-tech/starshipjs/actions/workflows/run-tests.yml">
    <img height="20" src="https://github.com/cosmology-tech/starshipjs/actions/workflows/run-tests.yml/badge.svg" />
  </a>
   <a href="https://github.com/cosmology-tech/starshipjs/blob/main/LICENSE"><img height="20" src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
</p>

**`starship`** is the command-line interface designed to deploy and manage [Starship](https://cosmology.zone/products/starship), catering specifically to Node.js and TypeScript developers. This CLI tool offers a seamless and intuitive interface that dramatically simplifies the development, testing, and deployment of interchain applications across both local environments and CI/CD pipelines.

Designed with simplicity and speed in mind, **Starship** enables developers to quickly integrate Starship into their blockchain projects without complex orchestration.

## Table of contents

- [StarshipJS](#starshipjs)
- [Table of contents](#table-of-contents)
- [Install](#install)
- [Examples](#examples)
- [Commands](#commands)
- [Options](#options)
- [Developing](#developing)
- [Credits](#credits)

## install

Install `@starship-ci/cli` globally to use the `starship` command:

```sh
npm install -g @starship-ci/cli
```

## Examples 

Here are some common usages of the starship CLI:

```sh
starship deploy --helmFile ./config/helm.yaml --helmName my-release
starship undeploy --config ./config/settings.json
starship setupKind
starship startPortForward
starship stopPortForward
starship upgrade
starship teardown
starship clean
```

## Commands

Here's a concise overview of the commands available in the `starship` CLI:


| Command          | Description                                     |
| ---------------- | ----------------------------------------------- |
| `deploy`         | Deploy starship using specified options or configuration file. |
| `setup`          | Setup initial configuration and dependencies.   |
| `startPortForward` | Start port forwarding for the deployed services. |
| `stopPortForward` | Stop port forwarding.                           |
| `teardown`       | Remove all components related to the deployment. |
| `upgrade`        | Upgrade the deployed application to a new version. |
| `undeploy`       | Remove starship deployment using specified options or configuration file. |
| `cleanKind`      | Clean up Kubernetes kind cluster resources.     |
| `deleteHelm`     | Delete a specific Helm release.                 |
| `removeHelm`     | Remove Helm chart from local configuration.     |
| `setupKind`      | Setup a Kubernetes kind cluster for development. |
| `clean`          | Perform a clean operation to tidy up resources. |
| `version`, `-v`  | Display the version of the Starship Client.     |

## Options

Options in the CLI allow for dynamic configuration of `starship`. You can specify individual options directly in the command line to fine-tune the behavior of the tool. Alternatively, you can use the `--config` option to specify a YAML configuration file that sets up the initial parameters. If both methods are used, any options provided directly in the command line will override the corresponding settings in the configuration file, giving you the flexibility to customize specific aspects of the deployment without altering the entire configuration.

| Option             | Description                                           |
|--------------------|-------------------------------------------------------|
| `--config <path>`  | Specify the path to the JSON configuration file containing all settings. |
| `--helmFile <path>`| Specify the path to the Helm file.                    |
| `--helmName <name>`| Specify the Helm release name.                        |
| `--helmRepo <repo>`| Specify the Helm repository.                          |
| `--helmRepoUrl <url>` | Specify the Helm repository URL.                    |
| `--helmChart <chart>` | Specify the Helm chart.                             |
| `--helmVersion <ver>` | Specify the version of the Helm chart.              |
| `--kindCluster <name>` | Specify the name of the Kubernetes kind cluster.   |
| `--verbose`        | Enable verbose output for debugging purposes.         |
| `--curdir <dir>`   | Specify the current working directory of the operation. |


## Developing

When first cloning the repo:
```
yarn
yarn build
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
