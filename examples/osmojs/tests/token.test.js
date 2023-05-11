import { generateMnemonic } from '@confio/relayer/build/lib/helpers';
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

import { ibcCosmosToOsmosis, sendOsmoToAddress } from '../src/utils';

import { ChainClients } from './setup.test';

describe("Token transfers", () => {
  let wallet;
  let baseDenom;
  let address;
  let chainClients;
  
  beforeAll(async () => {
    chainClients = ChainClients
    wallet = await DirectSecp256k1HdWallet.fromMnemonic(
      generateMnemonic(),
      { prefix: chainClients["osmosis-1"].getPrefix() },
    );
    baseDenom = chainClients["osmosis-1"].getDenom();
    address = (await wallet.getAccounts())[0].address;
  });
  
  it("send osmosis token to address", async () => {
    let chainClient = chainClients["osmosis-1"]
    // Transfer uosmo tokens from faceut
    await sendOsmoToAddress(chainClient, address);
    
    const balance = await chainClient.client.getBalance(address, baseDenom);
  
    expect(balance.amount).toEqual("100000000000");
    expect(balance.denom).toEqual(baseDenom);
  }, 10000);
  
  it("send ibc atom tokens to address", async () => {
    // Transfer uatom tokens via IBC to osmosis
    await ibcCosmosToOsmosis(chainClients["cosmos-2"], chainClients["osmosis-1"], address);
    
    // Check atom in address
    const chain = chainClients["osmosis-1"];
    const client = chain.client;
    const queryClient = chain.getQueryClient();
    
    const balances = await client.getAllBalances(address);
    
    // check balances
    expect(balances.length).toEqual(2);
    const ibcBalance = balances.find(balance => {
      return balance.denom.startsWith("ibc/")
    });
    expect(ibcBalance.amount).toEqual("100000000000");
    expect(ibcBalance.denom).toContain("ibc/");
    
    // check ibc denom trace of the same
    const trace = await queryClient.ibc.transfer.denomTrace(ibcBalance.denom.replace("ibc/", ""));
    expect(trace.denomTrace.baseDenom).toEqual(chainClients["cosmos-2"].getDenom());
  }, 10000);
});
