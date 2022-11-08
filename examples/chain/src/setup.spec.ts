import anyTest, {TestFn} from 'ava';
import { connect } from "./utils";
import { testutils } from "@confio/relayer";
import { getMnemonic } from "./keys";
import { coins, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { getChainInfo } from "./network";
import { assertIsDeliverTxSuccess, calculateFee } from "@cosmjs/stargate";
import {SigningCosmWasmClient} from "@cosmjs/cosmwasm-stargate";

const { fundAccount, generateMnemonic } = testutils;

const test = anyTest as TestFn<{client: SigningCosmWasmClient, address: string}>;

test.before(async (t) => {
  t.log("Connecting to osmosis client");

  const chainInfo = getChainInfo("osmosis")
  const { client: osmoClient, address: osmoAddress } = await connect("osmosis");

  t.context = {
    client: osmoClient,
    address: osmoAddress,
  }

  t.pass();
});

test.serial("send amount", async (t) => {
  t.log("Begin send amount")
  const { client: osmoClient, address: osmoAddress } = t.context;

  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(
    generateMnemonic(),
    { prefix: getChainInfo("osmosis").prefix }
  );
  const address = (await wallet.getAccounts())[0].address;
  const beforeBalance = await osmoClient.getBalance(address, "uosmo")

  await osmoClient.sendTokens(
    osmoAddress,
    address,
    coins(22664, "uosmo"),
    calculateFee(100_000, "0.025uosmo"),
    "send amount for test"
  );

  // Check balance of test address after
  const afterBalance = await osmoClient.getBalance(address, "uosmo")

  t.is(parseInt(afterBalance.amount) - parseInt(beforeBalance.amount), 22664)

  t.pass();
});
