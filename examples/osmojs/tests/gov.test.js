import { generateMnemonic } from '@confio/relayer/build/lib/helpers';
import { assertIsDeliverTxSuccess, SigningStargateClient } from '@cosmjs/stargate';
import { coin, coins } from '@cosmjs/amino';
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import BigNumber from 'bignumber.js';
import {osmosis, cosmos} from "osmojs";

import { ChainClients } from './setup.test';

describe("Governance tests for osmosis", () => {
  let wallet;
  let baseDenom;
  let address;
  let chainClients;
  let chain;
  
  // Variables used accross testcases
  let queryClient;
  let proposalId;
  
  beforeAll(async () => {
    chainClients = ChainClients
    baseDenom = chainClients["osmosis-1"].getDenom();
    chain = chainClients["osmosis-1"];
    
    // Initialize wallet
    wallet = await DirectSecp256k1HdWallet.fromMnemonic(
      generateMnemonic(),
      { prefix: chainClients["osmosis-1"].getPrefix() },
    );
    address = (await wallet.getAccounts())[0].address;
    
    // Create custom cosmos interchain client
    queryClient = await cosmos.ClientFactory.createRPCQueryClient({rpcEndpoint: chain.rpc});
    
    // Transfer osmosis and ibc tokens to address, send only osmo to address
    await chain.sendTokens(address, "100000000000");
  }, 200000);
  
  it("check address has tokens", async () => {
    const { balance } = await queryClient.cosmos.bank.v1beta1.balance({
      address,
      denom: chain.getDenom(),
    });
    
    expect(balance.amount).toEqual("100000000000");
  }, 10000);
  
  it("submit a txt proposal", async () => {
    const signingClient = await SigningStargateClient.connectWithSigner(
      chain.rpc,
      wallet,
      chain.stargateClientOpts(),
    );
    
    const contentMsg = cosmos.gov.v1beta1.TextProposal.fromPartial({
      title: "Test Proposal",
      description: "Test text proposal for the e2e testing",
    });

    // Stake half of the tokens
    const msg = cosmos.gov.v1beta1.MessageComposer.withTypeUrl.submitProposal({
      proposer: address,
      initialDeposit: coins(1_000_000, baseDenom),
      content: {
        typeUrl: "/cosmos.gov.v1beta1.TextProposal",
        value: cosmos.gov.v1beta1.TextProposal.encode(contentMsg).finish(),
      },
    })
    
    const fee = {
      amount: coins(100_000, baseDenom),
      gas: "200000",
    };
    
    const result = await signingClient.signAndBroadcast(address, [msg], fee);
    assertIsDeliverTxSuccess(result);
    
    // Get proposal id from log events
    const proposalIdEvent = result.events.find((event) => event.type === "submit_proposal");
    proposalId = proposalIdEvent.attributes.find((attr) => attr.key === "proposal_id").value;
    
    expect(BigInt(proposalId)).toBeGreaterThan(BigInt(0));
  });
  
  it("query proposal", async () => {});
  it("vote on proposal", async () => {});
  it("query vote", async () => {});
  it("verify proposal passed", async () => {});
});
