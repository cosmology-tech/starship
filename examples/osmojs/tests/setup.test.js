import { defaultRegistryTypes } from "@cosmjs/stargate";
import { describe } from "node:test";
import { getSigningOsmosisClientOptions } from "osmojs";
import path from "path";
import { StarshipClient } from '../src/clients.js';

// Global variable to store chain clients for the tests
export let ChainClients

beforeAll(async () => {
  const configFile = path.join(__dirname, "..", "configs", "config.yaml");
  
  // Instantiate osmosis chain client options
  const {
    registry,
    aminoTypes
  } = getSigningOsmosisClientOptions({
    defaultRegistryTypes
  });
  
  ChainClients = await StarshipClient.withConfigFile(configFile, {
    "osmosis-1": {
      gasPrice: "0.025uosmo",
      gasAdjustment: 1.5,
      registry,
      aminoTypes,
    },
  });
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
