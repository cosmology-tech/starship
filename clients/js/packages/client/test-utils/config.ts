import { readFileSync } from 'fs';
import * as yaml from 'js-yaml';
import { join, resolve } from 'path';

import { StarshipConfig } from '../src/config';

export const fixtureDir = resolve(join(__dirname, '/../../../__fixtures__'));
export const outputDir = resolve(join(__dirname, '/../../../__output__'));

function loadConfig(filename: string) {
  const configPath = join(fixtureDir, filename);
  const configAsYaml = readFileSync(configPath, 'utf-8');
  const config: StarshipConfig = yaml.load(configAsYaml) as StarshipConfig;
  return {
    configPath,
    configAsYaml,
    config
  };
}

const localConfig = loadConfig('local-config.yaml');
const config = loadConfig('config.yaml');

export { config, localConfig };
