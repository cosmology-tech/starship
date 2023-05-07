import { coin, coins } from '@cosmjs/amino';
import { assertIsDeliverTxSuccess } from '@cosmjs/stargate';
import Long from 'long';
import fetch from 'node-fetch';

export async function fetchJSON(url, options = {}) {
  const response = await fetch(url, {});
  return await response.json()
};

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

export async function ibcCosmosToOsmosis(cosmosChain, osmosisChain, address) {
  const client = cosmosChain.getClient();

  // Fetch open transfer channels and ports
  const ibcInfo = await cosmosChain.getIBCInfo(osmosisChain.getChainId());
  const channel = ibcInfo.channels[0]["chain_1"];

  const result = await client.sendIbcTokens(
    cosmosChain.address,
    address,
    coin(100_000_000_000, cosmosChain.getDenom()),
    channel["port_id"],
    channel["channel_id"],
    { revisionHeight: Long.fromNumber(12300), revisionNumber: Long.fromNumber(45600) },
    Math.floor(Date.now() / 1000) + 60,
    { amount: coins(10000, cosmosChain.getDenom()), gas: "200000" },
    "initial send atoms as part of setup",
  );

  // todo: fix this, better to wait for the broadcast to succed with a timeout
  await sleep(1*1000);

  assertIsDeliverTxSuccess(result);
}

// todo: use facuet here
export async function sendOsmoToAddress(osmosisChain, address) {
  const client = osmosisChain.getClient();
  const denom = osmosisChain.getDenom();

  const result = await client.sendTokens(
    osmosisChain.address,
    address,
    [coin(100_000_000_000, denom)],
    { amount: coins(10000, denom), gas: "200000" },
  );

  // todo: fix this, better to wait for the broadcast to succed with a timeout
  await sleep(1*1000);

  assertIsDeliverTxSuccess(result);
}
