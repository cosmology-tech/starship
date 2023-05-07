import {coins, makeCosmoshubPath} from '@cosmjs/amino';
import {DirectSecp256k1HdWallet} from '@cosmjs/proto-signing';
import {defaultRegistryTypes, QueryClient, setupIbcExtension, SigningStargateClient} from '@cosmjs/stargate';
import fetch from 'node-fetch';
import {getSigningOsmosisClient, getSigningOsmosisClientOptions} from 'osmojs';

import {getChainInfo, getStarshipConfig} from './config';

export class ChainClientRegistry {
  keys;

  constructor(chainId) {
    this.chainId = chainId;
    this.chainInfo = null;
    this.client = null;
    this.address = null;
}

  // initialize class object. chainInfo is initialized first before everything else
  async initialize() {
    this.chainInfo = await ChainClientRegistry.fetchChainInfo(this.chainId);
    this.keys = await ChainClientRegistry.fetchKeys();

    await this.initClient();
    this.queryClient = QueryClient.withExtensions(this.client.getTmClient(), setupIbcExtension)
  }

  stargateClientOpts() {
    let opts = {
        prefix: this.getPrefix(),
        gasPrice: `0${this.getDenom()}`,
    }

    if (this.getChainType() === "osmosis") {
      const {
        registry,
        aminoTypes
      } = getSigningOsmosisClientOptions({
        defaultRegistryTypes
      });
      Object.assign(opts, {registry, aminoTypes})
    }

    return opts;
  }

  // initialize client and address
  async initClient() {
    const prefix = this.chainInfo["bech32_prefix"];
    const denom = this.getDenom();
    const rpc = this.chainInfo["apis"]["rpc"][0]["address"];

    const hdPath = makeCosmoshubPath(0);

    // Setup signer
    const offlineSigner = await DirectSecp256k1HdWallet.fromMnemonic(this.getGenesisMnemonic(), {
      prefix,
      hdPaths: [hdPath],
    });
    const { address } = (await offlineSigner.getAccounts())[0];

    // Init SigningCosmWasmClient client
    const client = await SigningStargateClient.connectWithSigner(
      rpc,
      offlineSigner,
      this.stargateClientOpts(),
    );

    const chainId = await client.getChainId();
    if (chainId !== this.chainInfo["chain_id"]) {
      throw Error(`Given ChainId: ${this.getChainId()} doesn't match the clients ChainID: ${chainId}!`);
    }

    // set client and address
    this.client = client;
    this.address = address;
  }

  get rpc() {
    return this.chainInfo["apis"]["rpc"][0]["address"];
  }

  static async fetchChainInfo(chainId) {
    return await getChainInfo(chainId);
  }

  static async fetchKeys() {
    const url = "https://raw.githubusercontent.com/cosmology-tech/starship/main/charts/devnet/configs/keys.json";
    const response = await fetch(url, {});
    return await response.json();
  }

  getGenesisMnemonic() {
    return this.keys["genesis"][0]["mnemonic"]
  }

  getClient() {
    return this.client
  }

  getQueryClient() {
    return this.queryClient;
  }

  getChainInfo() {
    return this.chainInfo;
  }

  getPrefix() {
    return this.chainInfo["bech32_prefix"];
  }

  getChainId() {
    return this.chainInfo.chain_id;
  }

  getChainType() {
    return this.chainInfo["chain_name"];
  }

  getDenom() {
    return this.chainInfo.staking["staking_tokens"][0].denom
  }

  getDefaultFees() {
    return { amount: coins(10000, this.getDenom()), gas: "100000" };
  }

  // returns IBC info between current chain and given chain
  async getIBCInfo(chain2) {
    const config = getStarshipConfig();
    const port = config.registry.ports.rest;

    const resp = await fetch(`http://localhost:${port}/ibc/${this.chainId}/${chain2}`, {});
    return await resp.json();
  }

  static async withChainId(chainId) {
    const chainClient = new ChainClientRegistry(chainId);
    await chainClient.initialize();

    return chainClient;
  }

  // return a dict with key as chainId and value as the client
  static async withChainIds(chainIds) {
    const dict = {};
    await Promise.all(chainIds.map(async (chainId) => {
      dict[chainId] = await ChainClientRegistry.withChainId(chainId);
    }));
    return dict;
  }
}
