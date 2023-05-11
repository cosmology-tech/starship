import {describe} from "node:test";
import { ChainClientRegistry } from '../src/clients.js';

export let ChainClients

beforeAll(async () => {
  ChainClients = await ChainClientRegistry.withChainIds(["osmosis-1", "cosmos-2"]);
})

describe("Test clients", () => {
  let chainClients;
  
  beforeAll(() => {
    chainClients = ChainClients
  });
  
  it("check chain height", async () => {
    let chainClient = chainClients["osmosis-1"];
    const height = await chainClient.client.getHeight();
    
    expect(height).toBeGreaterThan(0);
  });
})