import { readFileSync } from "fs";

import { AckWithMetadata, CosmWasmSigner, RelayInfo, testutils } from "@confio/relayer";
import { fromBase64, fromUtf8 } from "@cosmjs/encoding";
import { assert } from "@cosmjs/utils";

const { fundAccount, generateMnemonic, osmosis: oldOsmo, signingClient, signingCosmWasmClient, wasmd: oldWasmd } = testutils;

import { readFile } from "fs/promises";
import { load } from "js-yaml";

export async function readVaules(): any {
  const filepath = "../../charts/devnet/values.yaml"
  const { values } = await readFile(filepath, "utf8")

  return values
}

// prettier-ignore
export const wasmd = {
  tendermintUrlWs: 'ws://localhost:26659',
  tendermintUrlHttp: 'http://localhost:26659',
  chainId: 'wasmd',
  prefix: 'wasm',
  denomStaking: 'stake',
  denomFee: 'stake',
  minFee: '0.025stake',
  blockTime: 1500,
  faucet: {
    mnemonic: 'enlist hip relief stomach skate base shallow young switch frequent cry park',
    pubkey0: {
      type: 'tendermint/PubKeySecp256k1',
      value: 'A9cXhWb8ZpqCzkA8dQCPV29KdeRLV3rUYxrkHudLbQtS',
    },
    address0: 'wasm14qemq0vw6y3gc3u3e0aty2e764u4gs5lndxgyk',
  },
  ics20Port: 'transfer',
  estimatedBlockTime: 1500,
  estimatedIndexerTime: 1500,
};

// prettier-ignore
export const osmosis = {
  tendermintUrlWs: 'ws://localhost:26653',
  tendermintUrlHttp: 'http://localhost:26653',
  chainId: 'osmosis',
  prefix: 'osmo',
  denomStaking: 'uosmo',
  denomFee: 'uosmo',
  minFee: '0.030uosmo',
  blockTime: 1500,
  faucet: {
    mnemonic: 'remain fragile remove stamp quiz bus country dress critic mammal office need',
    pubkey0: {
      type: 'tendermint/PubKeySecp256k1',
      value: 'A0d/GxY+UALE+miWJP0qyq4/EayG1G6tsg24v+cbD6By',
    },
    address0: 'osmo1lvrwcvrqlc5ktzp2c4t22xgkx29q3y83hdcc5d',
  },
  ics20Port: 'transfer',
  estimatedBlockTime: 1500,
  estimatedIndexerTime: 1500,
};

export const IbcVersion = "mesh-security-v0.1";

export async function setupContracts(
  cosmwasm: CosmWasmSigner,
  contracts: Record<string, string>
): Promise<Record<string, number>> {
  const results: Record<string, number> = {};

  for (const name in contracts) {
    const path = contracts[name];
    console.info(`Storing ${name} from ${path}...`);
    const wasm = await readFileSync(path);
    const receipt = await cosmwasm.sign.upload(cosmwasm.senderAddress, wasm, "auto", `Upload ${name}`);
    console.debug(`Upload ${name} with CodeID: ${receipt.codeId}`);
    results[name] = receipt.codeId;
  }

  return results;
}

export interface FundingOpts {
  readonly tendermintUrlHttp: string;
  readonly prefix: string;
  readonly denomFee: string;
  readonly denomStaking: string;
  readonly minFee: string;
  readonly estimatedBlockTime: number;
  readonly estimatedIndexerTime: number;
  readonly faucet: {
    readonly mnemonic: string;
  };
}

export async function fundStakingAccount(opts: FundingOpts, rcpt: string, amount: string): Promise<void> {
  const client = await signingClient(opts, opts.faucet.mnemonic);
  const stakeTokens = { amount, denom: opts.denomStaking };
  await client.sendTokens(rcpt, [stakeTokens]);
}

// This creates a client for the CosmWasm chain, that can interact with contracts
export async function setupWasmClient(): Promise<CosmWasmSigner> {
  // create apps and fund an account
  const mnemonic = generateMnemonic();
  const cosmwasm = await signingCosmWasmClient(wasmd, mnemonic);
  await fundAccount(wasmd, cosmwasm.senderAddress, "4000000");
  await fundStakingAccount(wasmd, cosmwasm.senderAddress, "4000000");
  return cosmwasm;
}

// This creates a client for the CosmWasm chain, that can interact with contracts
export async function setupOsmosisClient(): Promise<CosmWasmSigner> {
  // create apps and fund an account
  const mnemonic = generateMnemonic();
  const cosmwasm = await signingCosmWasmClient(osmosis, mnemonic);
  await fundAccount(osmosis, cosmwasm.senderAddress, "4000000");
  return cosmwasm;
}

// throws error if not all are success
export function assertAckSuccess(acks: AckWithMetadata[]) {
  for (const ack of acks) {
    const parsed = JSON.parse(fromUtf8(ack.acknowledgement));
    if (parsed.error) {
      throw new Error(`Unexpected error in ack: ${parsed.error}`);
    }
    if (!parsed.result) {
      throw new Error(`Ack result unexpectedly empty`);
    }
  }
}

// throws error if not all are errors
export function assertAckErrors(acks: AckWithMetadata[]) {
  for (const ack of acks) {
    const parsed = JSON.parse(fromUtf8(ack.acknowledgement));
    if (parsed.result) {
      throw new Error(`Ack result unexpectedly set`);
    }
    if (!parsed.error) {
      throw new Error(`Ack error unexpectedly empty`);
    }
  }
}

export function assertPacketsFromA(relay: RelayInfo, count: number, success: boolean) {
  if (relay.packetsFromA !== count) {
    throw new Error(`Expected ${count} packets, got ${relay.packetsFromA}`);
  }
  if (relay.acksFromB.length !== count) {
    throw new Error(`Expected ${count} acks, got ${relay.acksFromB.length}`);
  }
  if (success) {
    assertAckSuccess(relay.acksFromB);
  } else {
    assertAckErrors(relay.acksFromB);
  }
}

export function assertPacketsFromB(relay: RelayInfo, count: number, success: boolean) {
  if (relay.packetsFromB !== count) {
    throw new Error(`Expected ${count} packets, got ${relay.packetsFromB}`);
  }
  if (relay.acksFromA.length !== count) {
    throw new Error(`Expected ${count} acks, got ${relay.acksFromA.length}`);
  }
  if (success) {
    assertAckSuccess(relay.acksFromA);
  } else {
    assertAckErrors(relay.acksFromA);
  }
}

export function parseAcknowledgementSuccess(ack: AckWithMetadata): any {
  const response = parseString(ack.acknowledgement);
  assert(response.result);
  return parseBinary(response.result);
}

export function parseString(str: Uint8Array): any {
  return JSON.parse(fromUtf8(str));
}

export function parseBinary(bin: string): any {
  return JSON.parse(fromUtf8(fromBase64(bin)));
}
