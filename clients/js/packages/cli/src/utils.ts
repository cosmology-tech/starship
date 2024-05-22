import {defaultStarshipContext, StarshipConfig, StarshipContext} from '@starship-ci/client'; // Adjust the import path as necessary
import { type Question } from 'inquirerer';
import chalk from 'chalk';
import {readFileSync} from 'fs';
import * as yaml from 'js-yaml';
import {dirname, resolve} from 'path';

import {readAndParsePackageJson} from './package';
import deepmerge from 'deepmerge';

// Function to display the version information
export function displayVersion() {
    const pkg = readAndParsePackageJson();
    console.log(chalk.green(`Name: ${pkg.name}`));
    console.log(chalk.blue(`Version: ${pkg.version}`));
}


const resolvePath = (filename: string) =>
  filename.startsWith('/') ? filename : resolve((process.cwd(), filename));


const loadYaml = (filename: string): any => {
  const path = resolvePath(filename);
  const fileContents = readFileSync(path, 'utf8');
  return yaml.load(fileContents);
}

export interface Config {
  context: StarshipContext,
  starship: StarshipConfig
}

export const params: string[] = [
  'helmName',
  'helmFile',
  'helmRepo',
  'helmRepoUrl',
  'helmChart',
  'helmVersion',
]

export const loadConfig = (argv: any): Config => {
  console.log("argv: ", argv);
  console.log("argv.config: ", argv.config);
  let context: StarshipContext = { ...defaultStarshipContext } as StarshipContext;
  let starship: StarshipConfig = {} as StarshipConfig;

  if (argv.config) {
    context = deepmerge(defaultStarshipContext, loadYaml(argv.config)) as StarshipContext
  }

  console.log("context", context);

  // Override context with command-line arguments dynamically based on StarshipContext keys
  params.forEach(key => {
    if (argv[key] !== undefined) {
      console.log("key: ", key, " argv[key]: ", argv[key]);
      // @ts-ignore
      context[key] = argv[key];
    }
  });

  if (context.helmFile) {
    const dir = dirname(argv.config || process.cwd());
    context.helmFile = resolve(resolvePath(dir), context.helmFile);
    starship = loadYaml(context.helmFile) as StarshipConfig
  }

  console.log("starship: ", starship);

  return {context, starship};
}

export const usageText =`
Usage: starship <command> [options]

Commands:
  start              Start the Starship services.
  stop               Remove all components related to the deployment.
  deploy             Deploy starship using specified options or configuration file.
  setup              Setup initial configuration and dependencies.
  start-ports        Start port forwarding for the deployed services.
  stop-ports         Stop port forwarding.
  delete             Delete a specific Helm release.
  get-pods           Get the list of pods for the Helm release.
  clean              Perform a clean operation to tidy up resources.
  version, -v        Display the version of the Starship Client.

Configuration File:
  --config <path>       Specify the path to the JSON configuration file containing all settings.
                        Command-line options will override settings from this file if both are provided.

Command-line Options:
  --helmFile <path>     Specify the path to the Helm file, the actual config file. Required.
  --helmName <name>     Specify the Helm release name, default: starship.
  --helmVersion <ver>   Specify the version of the Helm chart, default: v0.2.0.

Examples:
  $ starship setup
  $ starship deploy --helmFile ./config/helm.yaml --helmName my-release
  $ starship start-ports --config ./config/settings.json
  $ starship stop --config ./config/settings.json

Additional Help:
  $ starship help          Display this help information.
`;

export function displayUsage() {
  console.log(usageText);
};