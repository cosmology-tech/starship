import test from "ava";
import { connect } from "./utils";
import { testutils } from "@confio/relayer";
import {getMnemonic} from "./keys";
import {coins, DirectSecp256k1HdWallet} from "@cosmjs/proto-signing";
import {getChainInfo} from "./network";
import {assertIsDeliverTxSuccess, calculateFee} from "@cosmjs/stargate";

const { fundAccount, generateMnemonic } = testutils;

test.before(async (t) => {
    console.debug("Setup osmosis pools")
    console.debug("Connecting to osmosis client")

    const { client: osmoClient, address: osmoAddress } = await connect("osmosis")

    t.context = {
        client: osmoClient,
        address: osmoAddress,
    };

    t.pass();
});

test.serial("send amount", async (t) => {
    const { client: osmoClient, address: osmoAddress } = await connect("osmosis")

    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(generateMnemonic(),
        { prefix: getChainInfo("osmosis").prefix }
    );
    const address = (await wallet.getAccounts())[0].address;
    const sendResp = await osmoClient.sendTokens(
        osmoAddress,
        address,
        coins(22664, "uosmo"),
        calculateFee(100_000, "0.025uosmo"),
        "send amount for test",
    );

    t.assert(assertIsDeliverTxSuccess(sendResp));

    t.pass();
});