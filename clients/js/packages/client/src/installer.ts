import chalk from 'chalk';
import * as os from 'os';
import * as shell from 'shelljs';

type Installation = {
  mac: string;
  linux: string;
};

export class StarshipInstaller {
  private installations: Record<string, Installation> = {
    docker: {
      mac: 'Please install Docker. Follow: https://docs.docker.com/desktop/install/mac-install/',
      linux:
        'Please install Docker. Follow: https://docs.docker.com/engine/install/ubuntu/'
    },
    kubectl: {
      mac: 'brew install kubectl',
      linux: `curl -Lks "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" > ~/.local/bin/kubectl && chmod +x ~/.local/bin/kubectl`
    },
    helm: {
      mac: 'brew install helm',
      linux:
        'curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash'
    }
  };

  async checkAndInstallBinary(binaryName: string) {
    if (!shell.which(binaryName)) {
      console.log(`${binaryName} is not installed`);
      await this.installBinary(binaryName);
      if (!shell.which(binaryName)) {
        console.log(
          chalk.red(
            `Installation of ${binaryName} failed. Please install ${binaryName} manually.`
          )
        );
        process.exit(1);
      }
    }
  }

  async checkAndInstallDependencies() {
    for (const dependency of Object.keys(this.installations)) {
      console.log(`Checking ${dependency}...`);
      await this.checkAndInstallBinary(dependency);
    }
  }

  async installBinary(binaryName: string) {
    const platform = os.platform();
    const installation = this.installations[binaryName];
    if (platform === 'darwin') {
      await this.runInstallation(installation.mac);
    } else if (platform === 'linux') {
      await this.runInstallation(installation.linux);
    }
  }

  private async runInstallation(command: string) {
    shell.exec(command);
  }
}

// Example usage:
// const client = new StarshipClient();
// await client.checkAndInstallBinary('kubectl');
