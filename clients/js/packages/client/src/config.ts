export interface Ports {
  rest?: number;
  rpc?: number;
  faucet?: number;
  exposer?: number;
  grpc?: number;
}

export interface Resources {
  cpu: number;
  memory: string;
}

export interface FaucetConfig {
  enabled: boolean;
  type: string;
}

export interface Chain {
  name: string;
  type: string;
  image: string;
  numValidators: number;
  ports: Ports;
  faucet?: FaucetConfig;
  resources?: Resources;
}

export interface Relayer {
  name: string;
  type: string;
  replicas: number;
  chains: string[];
}

export interface Explorer {
  enabled: boolean;
  ports?: Ports;
}

export interface Registry {
  enabled: boolean;
  ports: Ports;
}

export interface StarshipConfig {
  chains: Chain[];
  relayers: Relayer[];
  explorer?: Explorer;
  registry?: Registry;
}
