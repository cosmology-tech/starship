import { generateMnemonic } from '@confio/relayer/build/lib/helpers';
import {assertIsDeliverTxSuccess, setupIbcExtension, QueryClient, SigningStargateClient} from '@cosmjs/stargate';
import { coin, coins } from '@cosmjs/amino';
import Long from 'long';
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { osmosis } from 'osmojs';

import { ibcCosmosToOsmosis, sendOsmoToAddress } from '../src/utils';

import {calcShareOutAmount, daysToSeconds} from "./utils.js";
import { ChainClients } from './setup.test';


describe("Pool testing over IBC tokens", () => {
  let wallet;
  let baseDenom;
  let address;
  let chainClients;
  let chain;
  
  // Variables used accross testcases
  let poolId;
  let pool;

  beforeAll(async () => {
    chainClients = ChainClients
    wallet = await DirectSecp256k1HdWallet.fromMnemonic(
      generateMnemonic(),
      { prefix: chainClients["osmosis-1"].getPrefix() },
    );
    baseDenom = chainClients["osmosis-1"].getDenom();
    address = (await wallet.getAccounts())[0].address;
    chain = chainClients["osmosis-1"]
    
    // Transfer osmosis and ibc tokens to address
    await sendOsmoToAddress(chain, address);
    await ibcCosmosToOsmosis(chainClients["cosmos-2"], chainClients["osmosis-1"], address);
  }, 200000);

  it("check address has tokens", async () => {
    const balances = await chain.client.getAllBalances(address);
  
    expect(balances.length).toEqual(2);
  }, 10000);

  it("create ibc pools with ibc atom osmo", async () => {
    const signingClient = await SigningStargateClient.connectWithSigner(
      chainClients["osmosis-1"].rpc,
      wallet,
      chainClients["osmosis-1"].stargateClientOpts(),
    );
  
    const balances = await chain.client.getAllBalances(address);
  
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
      {amount: coins(10_000_000, chain.getDenom()), gas: "10000000"},
      "creating IBC pools",
    )
  
    assertIsDeliverTxSuccess(result);
  
    const poolCreated = result.events.find(x => x.type === "pool_created")
  
    // set poolid for the following test cases
    poolId = Long.fromString(poolCreated.attributes.find(x => x.key === "pool_id").value)
  
    expect(poolId.isPositive()).toBeTruthy()
  }, 200000);
  
  it("query pool via id, verify creation", async () => {
    // Query the created pool
    const queryClient = await osmosis.ClientFactory.createRPCQueryClient({
      rpcEndpoint: chain.rpc,
    });
    const poolResponse = await queryClient.osmosis.gamm.v1beta1.pool({poolId});
    
    expect(poolResponse).toBeTruthy()
    expect(poolResponse.pool.id.toInt()).toEqual(poolId.toInt())
    
    // Verify the address has gamm tokens
    const gammDenom = poolResponse.pool.totalShares.denom
    const gammBalance = await chain.client.getBalance(address, gammDenom);
    
    expect(gammBalance.denom).toEqual(gammDenom)
    expect(BigInt(gammBalance.amount)).toEqual(BigInt(poolResponse.pool.totalShares.amount))
    
    // Set pool var for other tests
    pool = poolResponse.pool
  }, 20000);

  it("join pool", async () => {
    const signingClient = await SigningStargateClient.connectWithSigner(
      chainClients["osmosis-1"].rpc,
      wallet,
      chainClients["osmosis-1"].stargateClientOpts(),
    );
    
    const allCoins = [coin("1000000", pool.poolAssets[0].token.denom), coin("1000000", pool.poolAssets[1].token.denom)]
    const shareOutAmount = calcShareOutAmount(pool, allCoins)
    const joinPoolMsg = osmosis.gamm.v1beta1.MessageComposer.withTypeUrl.joinPool({
      poolId: poolId,
      sender: address,
      shareOutAmount: shareOutAmount,
      tokenInMaxs: allCoins,
    })
    
    const resultPool = await signingClient.signAndBroadcast(
      address,
      [joinPoolMsg],
      { amount: coins(10_000_000, chain.getDenom()), gas: "10000000" },
      "join pool created",
    )
    
    assertIsDeliverTxSuccess(resultPool);
    
    // Verify new gamm tokens have been minted to the address
    const {denom: gammDenom, amount: totalgammAmount} = pool.totalShares
    const gammBalance = await chain.client.getBalance(address, gammDenom);
    
    expect(gammBalance.denom).toEqual(gammDenom)
    expect(BigInt(gammBalance.amount)).toEqual(BigInt(shareOutAmount) + BigInt(totalgammAmount))
  }, 200000);
  
  it("lockup tokens", () => {})
  it("swap tokens using pool", () => {})
});
