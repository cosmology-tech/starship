import { defaultStarshipContext, StarshipContext } from '@starship-ci/client'; // Adjust the import path as necessary
import { StarshipConfig } from '@starship-ci/client';
import chalk from 'chalk';
import { readFileSync } from 'fs';
import * as yaml from 'js-yaml';
import { dirname, resolve } from 'path';

import { readAndParsePackageJson } from './package';
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

export const loadConfig = (argv: any): Config => {
  if (argv.config) {
    const context = deepmerge(defaultStarshipContext, loadYaml(argv.config)) as StarshipContext
    if (context.helmFile) {
      const dir = dirname(argv.config);
      const configPath = resolve(resolvePath(dir), context.helmFile);
      context.helmFile = configPath;
      const starship = loadYaml(configPath) as StarshipConfig

      return {
        context,
        starship
      }
    }
  }
  return {
    // @ts-ignore
    context: {},
    // @ts-ignore
    starship: {}
  }
}

export const usageText =`
Usage: starship <command> [options]

Commands:
  deploy             Deploy starship using specified options or configuration file.
  setup              Setup initial configuration and dependencies.
  start-ports        Start port forwarding for the deployed services.
  stop-ports         Stop port forwarding.
  teardown           Remove all components related to the deployment.
  upgrade            Upgrade the deployed application to a new version.
  undeploy           Remove starship deployment using specified options or configuration file.
  delete-helm        Delete a specific Helm release.
  remove-helm        Remove Helm chart from local configuration.
  clean-kind         Clean up Kubernetes kind cluster resources.
  setup-kind         Setup a Kubernetes kind cluster for development.
  clean              Perform a clean operation to tidy up resources.
  version, -v        Display the version of the Starship Client.

Configuration File:
  --config <path>       Specify the path to the JSON configuration file containing all settings.
                        Command-line options will override settings from this file if both are provided.

Command-line Options:
  --helmFile <path>     Specify the path to the Helm file.
  --helmName <name>     Specify the Helm release name.
  --helmRepo <repo>     Specify the Helm repository.
  --helmRepoUrl <url>   Specify the Helm repository URL.
  --helmChart <chart>   Specify the Helm chart.
  --helmVersion <ver>   Specify the version of the Helm chart.
  --kindCluster <name>  Specify the name of the Kubernetes kind cluster.

Examples:
  $ starship setup
  $ starship deploy --helmFile ./config/helm.yaml --helmName my-release
  $ starship start-ports --config ./config/settings.json
  $ starship undeploy --config ./config/settings.json
  $ starship teardown

Additional Help:
  $ starship help          Display this help information.
`;

export function displayUsage() {
  console.log(usageText);
};