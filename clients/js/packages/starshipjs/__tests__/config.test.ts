import path from "path";

import { Config, ConfigContext } from "../src/config";
import { ChainRegistryFetcher } from "@chain-registry/client";

// it's more recommended to use ConfigContext.init to set the config file and registry.
it("1. throws without init;\n 2. init the setup and gets config;\n 3. throws when double init;\n", async () => {
  expect(() => ConfigContext.registry).toThrow();
  expect(() => ConfigContext.configFile).toThrow();

  const file = path.join(__dirname, "../../../__fixtures__", "config.yaml");

  // for unit test, only setup a chain registry fetcher without fetching.
  await Config.init(file, new ChainRegistryFetcher());

  const registry = ConfigContext.registry;
  const configFile = ConfigContext.configFile;

  expect(registry).toBeInstanceOf(ChainRegistryFetcher);
  expect(configFile).toBe(file);

  expect(
    async () => await ConfigContext.init(file, new ChainRegistryFetcher())
  ).rejects.toThrow();
});
