import { readFileSync } from 'fs';
import * as yaml from 'js-yaml';
import { join, resolve } from 'path';

const fixtureDir = resolve(join(__dirname, '/../../../__fixtures__'));

function loadConfig(filename: string) {
    const configPath = join(fixtureDir, filename);
    const configAsYaml = readFileSync(configPath, 'utf-8');
    const config = yaml.load(configAsYaml);
    return {
        configPath,
        configAsYaml,
        config
    };
}


const localConfig = loadConfig('local-config.yaml');
const config = loadConfig('config.yaml');

export { config, localConfig };
