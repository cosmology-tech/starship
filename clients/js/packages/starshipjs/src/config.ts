import { ChainRegistryFetcher } from '@chain-registry/client';
export class Config {
  private static instance: Config;
  public registry?: ChainRegistryFetcher;
  public configFile?: string;

  // keep instantiation private to enforce singletone
  private constructor() {}

  public static getInstance(): Config {
    if (!Config.instance) {
      Config.instance = new Config();
    }
    return Config.instance;
  }

  setConfigFile(configFile: string) {
    this.configFile = configFile;
  }

  setRegistry(registry: ChainRegistryFetcher) {
    this.registry = registry;
  }
}

export interface ChainConfig {
  registry: {
    ports: {
      rest: number;
    }
  };
  chains: Array<{
    id: string;
    name: string;
    ports: {
      rpc: number;
      rest: number;
      faucet: number;
    };
  }>;
  relayers: Array<{
    chains: [string, string];
  }>;
}

export const ConfigContext = Config.getInstance();