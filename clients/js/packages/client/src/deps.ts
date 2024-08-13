import * as shell from 'shelljs';

export type Dependency = {
  name: string;
  url: string;
  macUrl?: string; // Optional property for macOS-specific URLs
  installed: boolean;
};

export const dependencies: Dependency[] = [
  {
    name: 'kubectl',
    url: 'https://kubernetes.io/docs/tasks/tools/',
    macUrl: 'https://docs.docker.com/desktop/install/mac-install/',
    installed: !!shell.which('kubectl')
  },
  {
    name: 'docker',
    url: 'https://docs.docker.com/get-docker/',
    macUrl: 'https://docs.docker.com/desktop/install/mac-install/',
    installed: !!shell.which('docker')
  },
  {
    name: 'helm',
    url: 'https://helm.sh/docs/intro/install/',
    installed: !!shell.which('helm')
  }
];
