import { relative } from 'path';
import strip from 'strip-ansi';

import { StarshipClient } from '../src';
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
    name: 'osmojs',
    config: relative(process.cwd(), config.configPath)
  });

  const handler = {
    get(target: any, prop: string, _receiver: any) {
      const originalMethod = target[prop];
      if (typeof originalMethod === 'function') {
        return function (...args: any[]) {
          const argsString = args
            .map((arg) => strip(JSON.stringify(arg)))
            .join(', ');
          ctx.logs.push(`Call: ${prop}(${argsString})`);
          // if you want to see nested Call, replace target with this
          // I double checked both this and target, it does not call the exec in the methods when used internally
          return originalMethod.apply(target, args);
        };
      }
      return originalMethod;
    }
  };

  const proxiedClient: StarshipClient = new Proxy(client, handler);

  // Overriding the exit method
  // @ts-expect-error - Ignore lint error
  proxiedClient.exit = (code: number) => {
    ctx.code = code;
  };

  // @ts-expect-error - Ignore lint error
  proxiedClient.ensureFileExists = (_filename: string) => {};

  // Overriding the exec method
  // @ts-expect-error - Ignore lint error
  proxiedClient.exec = (cmd: string[]) => {
    // @ts-expect-error - Ignore lint error
    client.checkDependencies();
    ctx.commands.push(cmd.join(' '));
    ctx.logs.push(cmd.join(' '));
    return '';
  };

  // Overriding the log method
  // @ts-expect-error - Ignore lint error
  proxiedClient.log = (cmd: string) => {
    const str = strip(cmd);
    if (/\n/.test(str)) {
      ctx.logs.push('Log⬇\n' + str + '\nEndLog⬆');
    } else {
      ctx.logs.push('Log: ' + str);
    }
  };

  return {
    client: proxiedClient,
    ctx
  };
};

export const expectClient = (ctx: TestClientCtx, code: number) => {
  expect(ctx.logs.join('\n')).toMatchSnapshot();
  expect(ctx.commands.join('\n')).toMatchSnapshot();
  expect(ctx.code).toBe(code);
};
