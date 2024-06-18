import chalk from 'chalk';
import deepmerge from 'deepmerge';
import { existsSync, readFileSync, writeFileSync } from 'fs';
import * as yaml from 'js-yaml';
import { mkdirp } from 'mkdirp';
import * as os from 'os';
import { dirname, resolve } from 'path';
import * as shell from 'shelljs';

import {Chain, Relayer, Script, StarshipConfig} from './config';
import { Ports } from './config';
import { dependencies as defaultDependencies, Dependency } from "./deps";
import { readAndParsePackageJson } from './package';

export interface StarshipContext {
  helmName?: string;
  helmFile?: string;
  helmRepo?: string;
  helmRepoUrl?: string;
  helmChart?: string;
  helmVersion?: string;
  helmNamespace?: string;
  verbose?: boolean;
  curdir?: string;
};

export const defaultStarshipContext: Partial<StarshipContext> = {
  helmName: 'starship',
  helmRepo: 'starship',
  helmRepoUrl: 'https://cosmology-tech.github.io/starship/',
  helmChart: 'starship/devnet',
  helmVersion: 'v0.2.4',
  helmNamespace: '',
};

export interface PodPorts {
  registry?: Ports,
  explorer?: Ports,
  chains?: {
    defaultPorts?: Ports,
    [chainName: string]: Ports
  },
  relayers?: {
    defaultPorts?: Ports,
    [relayerName: string]: Ports
  },
}

// TODO talk to Anmol about moving these into yaml, if not already possible?
const defaultPorts: PodPorts = {
  explorer: {
    rest: 8080
  },
  registry: {
    grpc: 9090,
    rest: 8080
  },
  chains: {
    defaultPorts: {
      rpc: 26657,
      grpc: 9090,
      rest: 1317,
      exposer: 8081,
      faucet: 8000,
      cometmock: 22331,
    }
  },
  relayers: {
    defaultPorts: {
      rest: 3000,
      exposer: 8081,
    }
  }
};
export interface StarshipClientI {
  ctx: StarshipContext;
  version: string;
  dependencies: Dependency[];
  depsChecked: boolean;
  config: StarshipConfig;
  podPorts: PodPorts
};

export interface PodStatus {
  phase: string;
  ready: boolean;
  restartCount: number;
  reason?: string;
}

export const formatChainID = (input: string): string => {
  // Replace underscores with hyphens
  let formattedName = input.replace(/_/g, '-');

  // Truncate the string to a maximum length of 63 characters
  if (formattedName.length > 63) {
    formattedName = formattedName.substring(0, 63);
  }

  return formattedName;
};

export class StarshipClient implements StarshipClientI {
  ctx: StarshipContext;
  version: string;
  dependencies: Dependency[] = defaultDependencies;
  depsChecked: boolean = false;
  config: StarshipConfig;
  podPorts: PodPorts = defaultPorts;

  private podStatuses = new Map<string, PodStatus>(); // To keep track of pod statuses

  // Define a constant for the restart threshold
  private readonly RESTART_THRESHOLD = 4;

  constructor(ctx: StarshipContext) {
    this.ctx = deepmerge(defaultStarshipContext, ctx);
    // TODO add semver check against net
    this.version = readAndParsePackageJson().version;
  }

  private exec(cmd: string[], log: boolean = true, silent: boolean = false): shell.ShellString {
    this.checkDependencies();
    const str = cmd.join(' ');
    if (log) this.log(str);
    return shell.exec(str, { silent });
  }

  private log(str: string): void {
    // add log level
    console.log(str);
  }

  private exit(code: number): void {
    shell.exit(code);
  }

  private checkDependencies(): void {
    if (this.depsChecked) return;

    // so CI/CD and local dev work nicely
    const platform = process.env.NODE_ENV === 'test' ? 'linux' : os.platform();
    const messages: string[] = [];
    const depMessages: string[] = [];
    const missingDependencies = this.dependencies.filter(dep => !dep.installed);

    if (!missingDependencies.length) {
      this.depsChecked = true;
      return;
    }

    this.dependencies.forEach(dep => {
      if (missingDependencies.find(d => d.name === dep.name)) {
        depMessages.push(`${chalk.red('x')}${dep.name}`);
      } else {
        depMessages.push(`${chalk.green('✓')}${dep.name}`);
      }
    });

    messages.push('\n'); // Adding a newline for better readability

    missingDependencies.forEach(dep => {
      messages.push(chalk.bold.white(dep.name + ': ') + chalk.cyan(dep.url));

      if (dep.name === 'helm' && platform === 'darwin') {
        messages.push(chalk.gray("Alternatively, you can install using brew: ") + chalk.white.bold("`brew install helm`"));
      }

      if (dep.name === 'kubectl' && platform === 'darwin') {
        messages.push(chalk.gray("Alternatively, you can install Docker for Mac which includes Kubernetes: ") + chalk.white.bold(dep.macUrl));
      }

      if (dep.name === 'docker' && platform === 'darwin') {
        messages.push(chalk.gray("For macOS, you may also consider Docker for Mac: ") + chalk.white.bold(dep.macUrl));
      } else if (dep.name === 'docker') {
        messages.push(chalk.gray("For advanced Docker usage and installation on other platforms, see: ") + chalk.white.bold(dep.url));
      }

      messages.push('\n'); // Adding a newline for separation between dependencies
    });

    this.log(depMessages.join('\n'));
    this.log('\nPlease install the missing dependencies:');
    this.log(messages.join('\n'));
    this.exit(1);
  }

  public setup(): void {
    this.setupHelm();
  }

  private loadYaml(filename: string): any {
    const path = filename.startsWith('/') ? filename : resolve((process.cwd(), filename))
    const fileContents = readFileSync(path, 'utf8');
    return yaml.load(fileContents);
  }

  private saveYaml(filename: string, obj: any): any {
    const path = filename.startsWith('/') ? filename : resolve((process.cwd(), filename))
    const yamlContent = yaml.dump(obj);
    mkdirp.sync(dirname(path));
    writeFileSync(path, yamlContent, 'utf8');
  }

  public loadConfig(): void {
    this.ensureFileExists(this.ctx.helmFile);
    this.config = this.loadYaml(this.ctx.helmFile) as StarshipConfig;
  }

  public saveConfig(): void {
    this.saveYaml(this.ctx.helmFile, this.config);
  }

  public savePodPorts(filename: string): void {
    this.saveYaml(filename, this.podPorts);
  }

  public loadPodPorts(filename: string): void {
    this.ensureFileExists(filename);
    this.podPorts = this.loadYaml(filename) as PodPorts;
  }

  public setConfig(config: StarshipConfig): void {
    this.config = config;
  }

  public setContext(ctx: StarshipContext): void {
    this.ctx = ctx;
  }

  public setPodPorts(ports: PodPorts): void {
    this.podPorts = deepmerge(defaultPorts, ports);
  }

  public getArgs(): string[] {
    const args = [];
    if (this.ctx.helmNamespace) {
      args.push('--namespace', this.ctx.helmNamespace);
    }
    return args;
  }

  public getDeployArgs(): string[] {
    const args = this.getArgs();
    if (this.ctx.helmNamespace) {
      args.push('--create-namespace');
    }
    return args;
  }

  // TODO do we need this here?
  public test(): void {
    this.exec([
      'yarn',
      'run',
      'jest',
      `--testPathPattern=../${this.ctx.helmRepo}`,
      '--verbose',
      '--bail'
    ]);
  }

  public stop(): void {
    this.stopPortForward();
    this.deleteHelm();
  }

  public async start(): Promise<void> {
    this.setup();
    this.deploy();
    await this.waitForPods(); // Ensure waitForPods completes before starting port forwarding
    this.startPortForward();
  }

  public setupHelm(): void {
    this.exec([
      'helm',
      'repo',
      'add',
      this.ctx.helmRepo,
      this.ctx.helmRepoUrl
    ]);
    this.exec(['helm', 'repo', 'update']);
    this.exec([
      'helm',
      'search',
      'repo',
      this.ctx.helmChart,
      '--version',
      this.ctx.helmVersion
    ]);
  }

  private ensureFileExists(filename: string): void {
    const path = filename.startsWith('/') ? filename : resolve((process.cwd(), filename))
    if (!existsSync(path)) {
      throw new Error(`Configuration file not found: ${filename}`);
    }
  }

  public deploy(options: string[] = []): void {
    this.ensureFileExists(this.ctx.helmFile);
    this.log("Installing the helm chart. This is going to take a while.....");

    const cmd: string[] = [
      'helm',
      'install',
      '-f',
      this.ctx.helmFile,
      this.ctx.helmName,
      this.ctx.helmChart,
      '--version',
      this.ctx.helmVersion,
      ...this.getDeployArgs(),
      ...options,
    ];

    // Determine the data directory of the config file
    const datadir = resolve(dirname(this.ctx.helmFile!));

    // Iterate through each chain to add script arguments
    this.config.chains.forEach((chain, chainIndex) => {
      if (chain.scripts) {
        Object.keys(chain.scripts).forEach(scriptKey => {
          const script = chain.scripts?.[scriptKey as keyof Chain['scripts']];
          if (script && script.file) {
            const scriptPath = resolve(datadir, script.file);
            cmd.push(`--set-file chains[${chainIndex}].scripts.${scriptKey}.data=${scriptPath}`);
          }
        });
      }
    });

    this.exec(cmd);
    this.log("Run \"starship get-pods\" to check the status of the cluster");
  }

  public debug(): void {
    this.ensureFileExists(this.ctx.helmFile);
    this.deploy(['--dry-run', '--debug']);
  }

  public deleteHelm(): void {
    this.exec(['helm', 'delete', this.ctx.helmName]);
  }

  public getPods(): void {
    this.exec([
      "kubectl",
      "get pods",
      ...this.getArgs(),
      // "--all-namespaces"
    ]);
  }

  private getPodNames(): string[] {
    const result = this.exec([
      'kubectl',
      'get',
      'pods',
      '--no-headers',
      '-o',
      'custom-columns=:metadata.name',
      ...this.getArgs(),
    ], false, true)
  
    // Split the output by new lines and filter out any empty lines
    const podNames = result.split('\n').filter(name => name.trim() !== '');
  
    return podNames;
  }

  public areAllPodsRunning(): boolean {
    let allRunning = true;
    this.podStatuses.forEach((status) => {
      if (status.phase !== 'Running' || !status.ready) {
        allRunning = false;
      }
    });
    return allRunning;
  }

  private checkPodStatus(podName: string): void {
    const result = this.exec([
      'kubectl',
      'get',
      'pods',
      podName,
      '--no-headers',
      '-o',
      'custom-columns=:status.phase,:status.containerStatuses[*].ready,:status.containerStatuses[*].restartCount,:status.containerStatuses[*].state.waiting.reason',
      ...this.getArgs(),
    ], false, true).trim();

    const [phase, readyList, restartCountList, reason] = result.split(/\s+/);
    const ready = readyList.split(',').every(state => state === 'true');
    const restarts = restartCountList.split(',').reduce((acc, count) => acc + parseInt(count, 10), 0);

    this.podStatuses.set(podName, { phase, ready, restartCount: restarts, reason });

    if (restarts > this.RESTART_THRESHOLD) {
      this.log(`${chalk.red('Error:')} Pod ${podName} has restarted more than ${this.RESTART_THRESHOLD} times.`);
      this.exit(1);
    }
  }

  public async waitForPods(): Promise<void> {
    const podNames = this.getPodNames();

    // Check the status of each pod retrieved
    podNames.forEach(podName => {
      this.checkPodStatus(podName);
    });

    this.displayPodStatuses();

    if (!this.areAllPodsRunning()) {
      await new Promise(resolve => setTimeout(resolve, 2500));
      await this.waitForPods(); // Recursive call
    } else {
      this.log(chalk.green('All pods are running!'));
      // once the pods are in running state, wait for 10 more seconds
      await new Promise(resolve => setTimeout(resolve, 5000));
    }
  }

  private displayPodStatuses(): void {
    console.clear();
    this.podStatuses.forEach((status, podName) => {
      let statusColor;
      if (status.phase === 'Running' && status.ready) {
        statusColor = chalk.green(status.phase);
      } else if (status.phase === 'Running' && !status.ready) {
        statusColor = chalk.yellow('RunningButNotReady');
      } else if (status.phase === 'Terminating') {
        statusColor = chalk.gray(status.phase);
      } else {
        statusColor = chalk.red(status.phase);
      }

      console.log(`[${chalk.blue(podName)}]: ${statusColor}, ${chalk.gray(`Ready: ${status.ready}, Restarts: ${status.restartCount}`)}`);
    });
  }

  private forwardPort(chain: Chain, localPort: number, externalPort: number): void {
    if (localPort !== undefined && externalPort !== undefined) {
      this.exec([
        "kubectl", "port-forward",
        `pods/${formatChainID(chain.id)}-genesis-0`,
        `${localPort}:${externalPort}`,
        ...this.getArgs(),
        ">", "/dev/null",
        "2>&1", "&"
      ]);
      this.log(chalk.yellow(`Forwarded ${formatChainID(chain.id)}: local ${localPort} -> target (host) ${externalPort}`));
    }
  }

  private forwardPortCometmock(chain: Chain, localPort: number, externalPort: number): void {
    if (localPort !== undefined && externalPort !== undefined) {
      this.exec([
        "kubectl", "port-forward",
        `pods/${formatChainID(chain.id)}-cometmock-0`,
        `${localPort}:${externalPort}`,
        ...this.getArgs(),
        ">", "/dev/null",
        "2>&1", "&"
      ]);
      this.log(chalk.yellow(`Forwarded ${formatChainID(chain.id)}: local ${localPort} -> target (host) ${externalPort}`));
    }
  }

  private forwardPortRelayer(relayer: Relayer, localPort: number, externalPort: number): void {
    if (localPort !== undefined && externalPort !== undefined) {
      this.exec([
        "kubectl", "port-forward",
        `pods/${relayer.type}-${relayer.name}-0`,
        `${localPort}:${externalPort}`,
        ...this.getArgs(),
        ">", "/dev/null",
        "2>&1", "&"
      ]);
      this.log(chalk.yellow(`Forwarded ${relayer.name}: local ${localPort} -> target (host) ${externalPort}`));
    }
  }

  private forwardPortService(serviceName: string, localPort: number, externalPort: number): void {
    if (localPort !== undefined && externalPort !== undefined) {
      this.exec([
        "kubectl", "port-forward",
        `service/${serviceName}`,
        `${localPort}:${externalPort}`,
        ...this.getArgs(),
        ">", "/dev/null",
        "2>&1", "&"
      ]);
      this.log(`Forwarded ${serviceName} on port ${localPort} to target port ${externalPort}`);
    }
  }

  public startPortForward(): void {
    if (!this.config) {
      throw new Error('no config!');
    }
    this.log("Attempting to stop any existing port-forwards...");
    this.stopPortForward();
    this.log("Starting new port forwarding...");

    this.config.chains?.forEach(chain => {
      const chainPodPorts = this.podPorts.chains[chain.name] || this.podPorts.chains.defaultPorts;

      if (chain.cometmock?.enabled) {
        if (chain.ports?.rpc) this.forwardPortCometmock(chain, chain.ports.rpc, chainPodPorts.cometmock);
      } else {
        if (chain.ports?.rpc) this.forwardPort(chain, chain.ports.rpc, chainPodPorts.rpc);
      }
      if (chain.ports?.rest) this.forwardPort(chain, chain.ports.rest, chainPodPorts.rest);
      if (chain.ports?.grpc) this.forwardPort(chain, chain.ports.grpc, chainPodPorts.grpc);
      if (chain.ports?.exposer) this.forwardPort(chain, chain.ports.exposer, chainPodPorts.exposer);
      if (chain.ports?.faucet) this.forwardPort(chain, chain.ports.faucet, chainPodPorts.faucet);
    });

    this.config.relayers?.forEach(relayer => {
        const relayerPodPorts = this.podPorts.relayers[relayer.name] || this.podPorts.relayers.defaultPorts;
        if (relayer.ports?.rest) this.forwardPortRelayer(relayer, relayer.ports.rest, relayerPodPorts.rest);
        if (relayer.ports?.exposer) this.forwardPortRelayer(relayer, relayer.ports.exposer, relayerPodPorts.exposer);
    });

    if (this.config.registry?.enabled) {
      this.forwardPortService("registry", this.config.registry.ports.rest, this.podPorts.registry.rest);
      this.forwardPortService("registry", this.config.registry.ports.grpc, this.podPorts.registry.grpc);
    }

    if (this.config.explorer?.enabled) {
      this.forwardPortService("explorer", this.config.explorer.ports.rest, this.podPorts.explorer.rest);
    }
  }

  private getForwardPids(): string[] {
    const result = this.exec([
      "ps", "-ef",
      "|", "grep", "-i", "'kubectl port-forward'",
      "|", "grep", "-v", "'grep'",
      "|", "awk", "'{print $2}'"
    ]);
    const pids = (result || '').split('\n').map(pid => pid.trim()).filter(a => a !== '')
    return pids;
  }

  public stopPortForward(): void {
    this.log(chalk.green("Trying to stop all port-forward, if any...."));
    const pids = this.getForwardPids();
    pids.forEach(pid => {
      this.exec([
        "kill", "-15", pid
      ]);
    });
    this.exec(['sleep', '2']);
  }

  public printForwardPids(): void {
    const pids = this.getForwardPids();
    pids.forEach(pid => {
      console.log(pid);
    });
  }
}
