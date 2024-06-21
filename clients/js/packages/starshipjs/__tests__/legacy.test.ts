import path from "path";

import { Config, ConfigContext } from "../src/config";
import { ChainRegistryFetcher } from "@chain-registry/client";

// people can still use legacy ConfigContext to set the config file and registry.
it("1. throws without init;\n 2. throws only init partially;\n 3. init the setup and gets config;\n 4. throws when double init;\n", async () => {
  expect(() => ConfigContext.registry).toThrow();
  expect(() => ConfigContext.configFile).toThrow();

  const file = path.join(__dirname, "../../../__fixtures__", "config.yaml");

  ConfigContext.setConfigFile(file);

  expect(() => ConfigContext.registry).toThrow();
  expect(() => ConfigContext.configFile).toThrow();

  ConfigContext.setRegistry(new ChainRegistryFetcher());

  const registry = ConfigContext.registry;
  const configFile = ConfigContext.configFile;

  expect(registry).toBeInstanceOf(ChainRegistryFetcher);
  expect(configFile).toBe(file);

  expect(
    async () => await ConfigContext.init(file, new ChainRegistryFetcher())
  ).rejects.toThrow();
});
