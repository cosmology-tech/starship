import chalk from 'chalk';
import * as os from 'os';
import * as yaml from 'js-yaml';
import * as shell from 'shelljs';
import { StarshipConfig } from './config';
import { readFileSync } from 'fs';
import { dependencies as defaultDependencies, Dependency } from "./deps";
import { readAndParsePackageJson } from './package';

export interface StarshipContext {
  helmName: string;
  helmFile: string;
  helmRepo: string;
  helmRepoUrl: string;
  helmChart: string;
  helmVersion: string;
  kindCluster?: string;
  verbose?: boolean;
  curdir?: string;
};

export interface StarshipClientI {
  ctx: StarshipContext;
  version: string;
  dependencies: Dependency[];
  depsChecked: boolean;
  config: StarshipConfig;
};

export class StarshipClient implements StarshipClientI{
  ctx: StarshipContext;
  version: string;
  dependencies: Dependency[] = defaultDependencies;
  depsChecked: boolean = false;
  config: StarshipConfig;

  constructor(ctx: StarshipContext) {
    this.ctx = ctx;
    // TODO add semver check against net
    this.version = readAndParsePackageJson().version;
  }

  private exec(cmd: string[]): shell.ShellString {
    this.checkDependencies();
    const str = cmd.join(' ');
    this.log(str);
    return shell.exec(str);
  }

  private log(str: string): void {
    // add log level
    this.log(str);
  }

  private exit(code: number): void {
    shell.exit(code);
  }

  private checkDependencies(): void {
    if (this.depsChecked) return;

    const platform = os.platform();
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
        messages.push(chalk.gray("For advanced Docker usage and installation on other platforms, see: ") + chalk.white.bold("https://docs.docker.com/engine/install/"));
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

  public teardown(): void {
    this.removeHelm();
  }

  private loadConfig(): void {
    const fileContents = readFileSync(this.ctx.helmFile, 'utf8');
    this.config = yaml.load(fileContents) as StarshipConfig;
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
    this.stopForward();
    this.deleteHelm();
  }

  public clean(): void {
    this.stop();
    this.cleanKind();
  }

  private setupHelm(): void {
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
      `${this.ctx.helmRepo}/${this.ctx.helmChart}`,
      '--version',
      this.ctx.helmVersion
    ]);
  }

  private removeHelm(): void {
    this.exec([
      'helm',
      'repo',
      'remove',
      this.ctx.helmRepo
    ]);
  }

  public install(): void {
    this.log("Installing the helm chart. This is going to take a while.....");
    this.exec([
      'helm',
      'install',
      '-f',
      this.ctx.helmFile,
      this.ctx.helmName,
      `${this.ctx.helmRepo}/${this.ctx.helmChart}`,
      '--version',
      this.ctx.helmVersion
    ]);
    this.log("Run \"kubectl get pods\" to check the status of the cluster");
  }

  public upgrade(): void {
    this.exec([
      'helm',
      'upgrade',
      '--debug',
      '-f',
      this.ctx.helmFile,
      this.ctx.helmName,
      `${this.ctx.helmRepo}/${this.ctx.helmChart}`,
      '--version',
      this.ctx.helmVersion
    ]);
  }

  public debug(): void {
    this.exec([
      'helm',
      'install',
      '--dry-run',
      '--debug',
      '-f',
      this.ctx.helmFile,
      this.ctx.helmName,
      `${this.ctx.helmRepo}/${this.ctx.helmChart}`
    ]);
  }

  private deleteHelm(): void {
    this.exec(['helm', 'delete', this.ctx.helmName]);
  }

  public startPortForward(): void {
    this.stopPortForward();
    this.log(chalk.magenta(`Port forwarding for config ${this.ctx.helmFile}`));
    this.log(chalk.magenta("Port forwarding all chains"));

    if (!this.config.chains.length) {
      this.log(chalk.red("No chains to port-forward."));
      return;
    }

    this.config.chains.forEach(chain => {
      const { rpc, rest, exposer, faucet } = chain.ports;
      const forwardPort = (localPort: number, targetPort: number) => {
        if (localPort !== undefined && targetPort !== undefined) {
          this.exec([
            "kubectl", "port-forward",
            `pods/${chain.name}-genesis-0`,
            `${localPort}:${targetPort}`,
            ">", "/dev/null",
            "2>&1", "&"
          ]);
        }
      };

      if (rpc) forwardPort(rpc, 26657);
      if (rest) forwardPort(rest, 1317);
      if (exposer) forwardPort(exposer, 8081);
      if (faucet) forwardPort(faucet, 8000);

      this.log(chalk.yellow(`chains: forwarded ${chain.name} lcd to http://localhost:${rest}, rpc to http://localhost:${rpc}, faucet to http://localhost:${faucet}`));
    });

    if (this.config.registry?.enabled) {
      this.exec([
        "kubectl", "port-forward",
        "service/registry",
        "8081:8080",
        ">", "/dev/null",
        "2>&1", "&"
      ]);
      this.exec([
        "kubectl", "port-forward",
        "service/registry",
        "9091:9090",
        ">", "/dev/null",
        "2>&1", "&"
      ]);
      this.log(chalk.yellow("registry: forwarded registry lcd to grpc http://localhost:8081, to http://localhost:9091"));
    }

    if (this.config.explorer?.enabled) {
      this.exec([
        "kubectl", "port-forward",
        "service/explorer",
        "8080:8080",
        ">", "/dev/null",
        "2>&1", "&"
      ]);
      this.log(chalk.green("Open the explorer to get started.... http://localhost:8080"));
    }
  }

  public stopForward(): void {
    this.exec(['pkill', '-f', 'port-forward']);
  }

  // TODO review with Anmol, which stopForward is better...
  public stopPortForward(): void {
    this.log(chalk.green("Trying to stop all port-forward, if any...."));
    const pids = this.exec([
      "ps", "-ef",
      "|", "grep", "-i", "'kubectl port-forward'",
      "|", "grep", "-v", "'grep'",
      "|", "awk", "'{print $2}'"
    ]).split('\n');
    pids.forEach(pid => {
      if (pid.trim()) {
        this.exec([
          "kill", "-15", pid
        ]);
      }
    });
    this.exec(['sleep', '2']);
  }


  private setupKind(): void {
    if (this.ctx.kindCluster) {
      this.exec(['kind', 'create', 'cluster', '--name', this.ctx.kindCluster]);
    }
  }

  private cleanKind(): void {
    if (this.ctx.kindCluster) {
      this.exec(['kind', 'delete', 'cluster', '--name', this.ctx.kindCluster]);
    }
  }
}
