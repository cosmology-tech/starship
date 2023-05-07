import yaml from 'js-yaml';
import fetch from 'node-fetch';
import path from 'path';

const fs = require('fs');

let KEYS;

// returns starship config used to spun up the cluster
export function getStarshipConfig() {
  const configPath = path.join(__dirname, "configs", "config.yaml");
  return yaml.load(fs.readFileSync(configPath, "utf-8"));
}

// todo: use @chain-registry/types and @chain-registry/client
// fetches chain registry info from local chain registry spun up by Starship
export async function getChainInfo(name) {
  const config = getStarshipConfig();
  const port = config.registry.ports.rest;

  const resp = await fetch(`http://localhost:${port}/chains/${name}`, {});
  const chainInfo = await resp.json();

  // Replace chain rpc with localhost
  const chain = config["chains"].find(x => x.name === name);
  chainInfo["apis"]["rpc"][0]["address"] = `http://localhost:${chain.ports.rpc}`;

  return chainInfo;
}
