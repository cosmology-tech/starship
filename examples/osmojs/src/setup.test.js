import { generateMnemonic } from '@confio/relayer/build/lib/helpers';
import {assertIsDeliverTxSuccess, setupIbcExtension, QueryClient, SigningStargateClient} from '@cosmjs/stargate';
import { coin, coins } from '@cosmjs/amino';
import Long from 'long';
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { osmosis } from 'osmojs';

import { ibcCosmosToOsmosis, sendOsmoToAddress } from './utils';
import { ChainClientRegistry } from './clients.js';


describe("IBC transfer of atom", () => {
  let wallet;
  let baseDenom;
  let address;
  let chainClients;

  beforeAll(async () => {
    chainClients = await ChainClientRegistry.withChainIds(["osmosis-1", "cosmos-2"]);
  });

  beforeEach(async () => {
    wallet = await DirectSecp256k1HdWallet.fromMnemonic(
      generateMnemonic(),
      { prefix: chainClients["osmosis-1"].getPrefix() },
    );
    baseDenom = chainClients["osmosis-1"].getDenom();
    address = (await wallet.getAccounts())[0].address;
  });

  it("check address has osmosis token", async () => {
    let chainClient = chainClients["osmosis-1"]
    // Transfer uosmo tokens from faceut
    await sendOsmoToAddress(chainClient, address);

    const genesisBalance = await chainClient.client.getBalance(chainClient.address, baseDenom);
    const balance = await chainClient.client.getBalance(address, baseDenom);

    expect(balance.amount).toEqual("100000000000");
    expect(balance.denom).toEqual(baseDenom);
  }, 10000);

  it("check address has IBC tokens", async () => {
    // Transfer uatom tokens via IBC to osmosis
    await ibcCosmosToOsmosis(chainClients["cosmos-2"], chainClients["osmosis-1"], address);

    // Check atom in address
    const chain = chainClients["osmosis-1"];
    const client = chain.client;
    const queryClient = chain.getQueryClient();

    const balances = await client.getAllBalances(address);

    // check balances
    expect(balances.length).toEqual(1);
    expect(balances[0].amount).toEqual("100000000000");
    expect(balances[0].denom).toContain("ibc/");

    // check ibc denom trace of the same
    const trace = await queryClient.ibc.transfer.denomTrace(balances[0].denom.replace("ibc/", ""));
    expect(trace.denomTrace.baseDenom).toEqual(chainClients["cosmos-2"].getDenom());
  }, 10000);

  it("create ibc pools with ibc atom osmo", async () => {
    const signingClient = await SigningStargateClient.connectWithSigner(
      chainClients["osmosis-1"].rpc,
      wallet,
      chainClients["osmosis-1"].stargateClientOpts(),
    );

    // Transfer uosmo tokens from faceut
    await sendOsmoToAddress(chainClients["osmosis-1"], address);
    // Transfer uatom tokens via IBC to osmosis
    await ibcCosmosToOsmosis(chainClients["cosmos-2"], chainClients["osmosis-1"], address);

    const chain = chainClients["osmosis-1"];
    const client = chain.client;

    const balances = await client.getAllBalances(address);

    const msg = osmosis.gamm.poolmodels.balancer.v1beta1.MessageComposer.fromPartial.createBalancerPool({
      sender: address,
      poolParams: {
        swapFee: "1",
        exitFee: "0",
      },
      poolAssets: [
        {
          token: coin("10000000", balances[0].denom),
          weight: "100",
        },
        {
          token: coin("10000000", balances[1].denom),
          weight: "100",
        }
      ],
      futurePoolGovernor: "",
    });

    const result = await signingClient.signAndBroadcast(
      address,
      [msg],
      { amount: coins(1_000_000_000, chain.getDenom()), gas: "10000000" },
      "creating IBC pools",
    )

    assertIsDeliverTxSuccess(result);

    const poolCreated = result.events.find(x => x.type === "pool_created")
    const poolId = Long.fromString(poolCreated.attributes.find(x => x.key === "pool_id").value)

    expect(poolId.isPositive()).toBeTruthy()
  }, 20000);

  it("join and actions on pool", async () => {
    const signingClient = await SigningStargateClient.connectWithSigner(
      chainClients["osmosis-1"].rpc,
      wallet,
      chainClients["osmosis-1"].stargateClientOpts(),
    );

    // Transfer uosmo tokens from faceut
    await sendOsmoToAddress(chainClients["osmosis-1"], address);
    // Transfer uatom tokens via IBC to osmosis
    await ibcCosmosToOsmosis(chainClients["cosmos-2"], chainClients["osmosis-1"], address);

    const chain = chainClients["osmosis-1"];
    const client = chain.client;

    const balances = await client.getAllBalances(address);

    const msg = osmosis.gamm.poolmodels.balancer.v1beta1.MessageComposer.fromPartial.createBalancerPool({
      sender: address,
      poolParams: {
        swapFee: "1",
        exitFee: "0",
      },
      poolAssets: [
        {
          token: coin("10000000", balances[0].denom),
          weight: "100",
        },
        {
          token: coin("10000000", balances[1].denom),
          weight: "100",
        }
      ],
      futurePoolGovernor: "",
    });

    const result = await signingClient.signAndBroadcast(
      address,
      [msg],
      { amount: coins(1_000_000_000, chain.getDenom()), gas: "10000000" },
      "creating IBC pools",
    )

    assertIsDeliverTxSuccess(result);

    const poolCreated = result.events.find(x => x.type === "pool_created")
    const poolId = Long.fromString(poolCreated.attributes.find(x => x.key === "pool_id").value)

    expect(poolId.isPositive()).toBeTruthy()

    const queryClient = await osmosis.ClientFactory.createRPCQueryClient({
      rpcEndpoint: chain.rpc,
    });

    const poolResponse = await queryClient.osmosis.gamm.v1beta1.pool({poolId});

    expect(poolResponse).toBeTruthy()
  }, 200000);
});
