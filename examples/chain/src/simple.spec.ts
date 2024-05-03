import anyTest, {TestFn} from 'ava';
import { connect } from "./utils";
import { testutils } from "@confio/relayer";
import { getMnemonic } from "./keys";
import { coins, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import {getChainInfo, Network} from "./network";
import { assertIsDeliverTxSuccess, calculateFee } from "@cosmjs/stargate";
import {SigningCosmWasmClient} from "@cosmjs/cosmwasm-stargate";

const { fundAccount, generateMnemonic } = testutils;

const test = anyTest as TestFn<{
  client: SigningCosmWasmClient,
  address: string,
  chainInfo: Network,
}>;

test.before(async (t) => {
  t.log("Connecting to osmosis client");

  const chainInfo = getChainInfo("osmosis")
  const { client: osmoClient, address: osmoAddress } = await connect("osmosis");

  t.context = {
    client: osmoClient,
    address: osmoAddress,
    chainInfo: chainInfo,
  };

  t.pass();
});

test.serial("send amount to address", async (t) => {
  t.log("Begin send amount")
  const { client: osmoClient, address: osmoAddress, chainInfo } = t.context;

  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(
    generateMnemonic(),
    { prefix: getChainInfo("osmosis").prefix }
  );
  const address = (await wallet.getAccounts())[0].address;
  const beforeBalance = await osmoClient.getBalance(address, chainInfo.denom)

  await osmoClient.sendTokens(
    osmoAddress,
    address,
    coins(22664, chainInfo.denom),
    calculateFee(100_000, chainInfo.gasPrice),
    "send amount for test"
  );

  // Check balance of test address after
  const afterBalance = await osmoClient.getBalance(address, chainInfo.denom)

  t.is(parseInt(afterBalance.amount) - parseInt(beforeBalance.amount), 22664)

  t.pass();
});
