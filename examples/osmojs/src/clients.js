import { coin, coins } from '@cosmjs/amino';
import { DirectSecp256k1HdWallet } from '@cosmjs/proto-signing';
import {
  assertIsDeliverTxSuccess,
  QueryClient,
  setupIbcExtension,
  SigningStargateClient
} from '@cosmjs/stargate';
import yaml from "js-yaml";
import fs from "fs";
import Long from "long";
import fetch from 'node-fetch';

import { makeCosmoshubPath } from './utils';



// StarshipClient is a wrapper around StargateClient that provides additional
// funtaionalities to make it easier to use in tests with Starship backend infra
export class StarshipClient {
  keys;

  constructor(chainId, config, clientOpts) {
    this.chainId = chainId;
    this.config = config;
    this.clientOpts = clientOpts;
    this.chainInfo = null;
    this.client = null;
    this.address = null;
    this.wallet = null;
}

  // initialize class object. chainInfo is initialized first before everything else
  async initialize() {
    this.chainInfo = await this.fetchChainInfo(this.chainId, this.config);

    await this.initClient();
    this.queryClient = QueryClient.withExtensions(this.client.getTmClient(), setupIbcExtension)
  }

  stargateClientOpts() {
    return this.clientOpts || {
        prefix: this.getPrefix(),
        gasPrice: this.getGasPrice(),
    }
  }

  // initialize client and address
  async initClient() {
    const prefix = this.chainInfo["bech32_prefix"];
    const rpc = this.chainInfo["apis"]["rpc"][0]["address"];
    const slip44 = this.chainInfo["slip44"];

    const hdPath = makeCosmoshubPath(slip44, 0);

    // Setup signer
    const mnemonic = await this.getGenesisMnemonic();
    const offlineSigner = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
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
    this.wallet = offlineSigner;
    this.address = address;
  }

  get rpc() {
    const chain = this.config["chains"].find(chain => chain.name === this.chainId);
    return `http://localhost:${chain.ports.rpc}`;
  }

  async fetchKeys() {
    const url = `${this.getRegistryUrl()}/chains/${this.chainId}/keys`;
    const response = await fetch(url, {});
    return await response.json();
  }

  async getGenesisMnemonic() {
    const keys = await this.fetchKeys();
    return keys["genesis"][0]["mnemonic"];
  }

  getClient() {
    return this.client
  }

  getQueryClient() {
    return this.queryClient;
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
  
  getGasPrice() {
    const price = this.chainInfo["fees"]["fee_tokens"][0]["fixed_min_gas_price"];
    return `${price}${this.getDenom()}`;
  }

  getDefaultFees() {
    return { amount: coins(100000, this.getDenom()), gas: "1000000" };
  }
  
  getRegistryUrl() {
    const port = this.config.hasOwnProperty("registry") ? this.config.registry.ports.rest : 8080;
    return "http://localhost:" + port;
  }

  // returns IBC info between current chain and given chain
  async getIBCInfo(chain2) {
    const resp = await fetch(`${this.getRegistryUrl()}/ibc/${this.chainId}/${chain2}`, {});
    return await resp.json();
  }
  
  async fetchChainInfo() {
    const url = "http://localhost:" + (this.config.hasOwnProperty("registry") ? this.config.registry.ports.rest : 8080);
    const resp = await fetch(`${url}/chains/${this.chainId}`, {});
    const chainInfo = await resp.json();
    
    // Replace chain rpc with localhost rpc
    const chain = this.config["chains"].find(chain => chain.name === this.chainId);
    chainInfo.apis.rpc[0].address = `http://localhost:${chain.ports.rpc}`;
    
    return chainInfo;
  }
  
  // helper functions for testing utils
  waitForTx(hash, timeout = 10000, interval = 1000) {
    return new Promise((resolve, reject) => {
      const startTime = Date.now();
      const intervalId = setInterval(async () => {
        const tx = await this.client.getTx(hash);
        if (tx.code === 0) {
          clearInterval(intervalId);
          resolve(tx);
        } else if (Date.now() - startTime > timeout) {
          clearInterval(intervalId);
          reject(new Error('Transaction timeout'));
        }
      }, interval);
    });
  }
  
  async sendTokens(toAddress, amount) {
    const denom = this.getDenom();
    
    const resp = await this.client.sendTokens(
      this.address,
      toAddress,
      [coin(amount, denom)],
      this.getDefaultFees(),
    );
    assertIsDeliverTxSuccess(resp);
    
    await this.waitForTx(resp.transactionHash);
  }
  
  async sendIBCTokens(toAddress, amount, chain2) {
    const denom = this.getDenom();
    const ibcInfo = await this.getIBCInfo(chain2);
    const channel = ibcInfo.channels[0]["chain_1"];
    
    const resp = await this.client.sendIbcTokens(
      this.address,
      toAddress,
      coin(amount, denom),
      channel["port_id"],
      channel["channel_id"],
      { revisionHeight: Long.fromNumber(12300), revisionNumber: Long.fromNumber(45600) },
      Math.floor(Date.now() / 1000) + 60,
      this.getDefaultFees(),
      `starship: ibc transfer from ${this.chainId} to ${chain2}`,
    );
    
    assertIsDeliverTxSuccess(resp);
    await this.waitForTx(resp.transactionHash);
  }

  // Static functions for creating clients
  static async withChainId(chainId, config, chainOpts = {}) {
    const chainClient = new StarshipClient(chainId, config, chainOpts);
    await chainClient.initialize();

    return chainClient;
  }
  
  static async withConfigFile(configFile, chainOpts = {}) {
    const config = yaml.load(fs.readFileSync(configFile, 'utf8'));
    const chainClients = {};
    
    const chainIds = config.chains.map(chain => chain.name);
    await Promise.all(chainIds.map(async (chainId) => {
      const chainOpt = chainOpts.hasOwnProperty(chainId) ? chainOpts[chainId] : {};
      chainClients[chainId] = await StarshipClient.withChainId(chainId, config, chainOpt);
    }));

    return chainClients;
  }
}
