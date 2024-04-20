import { relative } from 'path';
import strip from 'strip-ansi'

import { StarshipClient } from '../src'
import { config } from './config';

interface TestClientCtx {
  commands: string[];
  logs: string[];
  code: number;
}

export const createClient = () => {
  const ctx: TestClientCtx = {
    commands: [],
    logs: [],
    code: -1
  };

  const client = new StarshipClient({
    helmName: 'osmojs',
    helmFile: relative(process.cwd(), config.configPath),
    helmRepo: 'starship',
    helmRepoUrl: 'https://cosmology-tech.github.io/starship/',
    helmChart: 'devnet',
    helmVersion: 'v0.1.38'
  });

  const handler = {
    get(target: any, prop: string, _receiver: any) {
      const originalMethod = target[prop];
      if (typeof originalMethod === 'function') {
        return function(...args: any[]) {
          const argsString = args.map(arg => strip(JSON.stringify(arg))).join(', ');
          ctx.logs.push(`Call: ${prop}(${argsString})`);
          // @ts-ignore
          return originalMethod.apply(this, args);
        };
      }
      return originalMethod;
    }
  };

  const proxiedClient: StarshipClient = new Proxy(client, handler);

  // Overriding the exit method
  // @ts-ignore
  proxiedClient.exit = (code: number) => {
    ctx.code = code;
  };
  
  // Overriding the exec method
  // @ts-ignore
  proxiedClient.exec = (cmd: string[]) => {
    // @ts-ignore
    proxiedClient.checkDependencies();
    ctx.commands.push(cmd.join(' '));
    ctx.logs.push(cmd.join(' '));
  };

  // Overriding the log method
  // @ts-ignore
  proxiedClient.log = (cmd: string) => {
    ctx.logs.push(strip(cmd));
  };

  return {
    client: proxiedClient,
    ctx
  }
}


export const expectClient = (ctx: TestClientCtx, code: number) => {
  expect(ctx.logs.join('\n')).toMatchSnapshot();
  expect(ctx.commands.join('\n')).toMatchSnapshot();
  expect(ctx.code).toBe(code);
}
