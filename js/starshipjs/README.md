# StarshipJS

<p align="center">
  <img src="https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png" width="80"><br />
  StarshipJS enables developers to efficiently set up and test chains, explorers, and validators, making it easier to handle development projects spanning several blockchain networks.
</p>

<p align="center" width="100%">
  <a href="https://github.com/cosmology-tech/starshipjs/actions/workflows/run-tests.yml">
    <img height="20" src="https://github.com/cosmology-tech/starshipjs/actions/workflows/run-tests.yml/badge.svg" />
  </a><a href="https://github.com/cosmology-tech/starshipjs/blob/main/LICENSE"><img height="20" src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
</p>

**StarshipJS** is the JS companion to deploy and manage [Starship](https://cosmology.zone/products/starship), tailored specifically for Node.js and TypeScript developers. This toolkit provides a seamless, easy-to-use interface that dramatically simplifies the development, testing, and deployment of interchain applications, whether on your local machine or CI/CD environments.

Designed with simplicity and speed in mind, **StarshipJS** enables developers to quickly integrate Starship into their blockchain projects without complex orchestration.

## Features

👨🏻‍💻 **Node.js and TypeScript Focused**: Tailored specifically for JavaScript ecosystems, **StarshipJS** brings simplicity to multi-chain development for Node.js and TypeScript environments, streamlining the setup and coding processes.

🚀 **Simplified Interchain Development**: Enables the straightforward creation of applications that span multiple blockchain networks. This simplifies complex blockchain interactions, enhancing interoperability and making it easier to build sophisticated interchain solutions.

🔒 **Security-First Approach**: **StarshipJS** incorporates security best practices from the ground up. Facilitates secure coding practices and configurations, helping developers build secure blockchain applications by default, reducing the risk of vulnerabilities.

## Table of contents

- [StarshipJS](#starshipjs)
- [Table of contents](#table-of-contents)
- [Install](#install)
- [Usage](#usage)
- [Developing](#developing)
- [Credits](#credits)

## install

Install the test utilities `starshipjs`:

```sh
npm install starshipjs

```

## Usage 

StarshipJS is a utility library that provides helpers to leverage [Starship](https://github.com/cosmology-tech/starship)'s internal chain registry, emulating the style of code used in projects like [cosmos-kit](https://github.com/cosmology-tech/cosmos-kit).

### Configuration

Before using StarshipJS, you need to set up the configuration for your blockchain network.

```js
import { Config } from 'starshipjs';
import path from 'path';

// Path to your YAML configuration file
const configFile = path.join(__dirname, 'path', 'to', 'your', 'config.yaml');

// Set the configuration file in StarshipJS
Config.setConfigFile = configFile;
```

### Registry

```js
import { useRegistry, Config } from 'starshipjs';

Config.setRegistry = await useRegistry(Config.configFile);
```

## Chain Info

Get detailed chain information about the blockchain network:

```js
const { chainInfo } = useChain('osmosis');

console.log(chainInfo);
```

## Credits and Faucets

If your blockchain network supports faucets, you can use them to get test tokens:

```js
const { creditFromFaucet } = useChain('osmosis');
const address = 'your-blockchain-address';

await creditFromFaucet(address);
```

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
