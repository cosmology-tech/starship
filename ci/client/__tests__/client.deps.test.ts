import { createClient, expectClient } from '../test-utils/client';

describe('StarshipClient', () => {
  it('missing deps', () => {
    const { client, ctx } = createClient();

    client.dependencies.find(dep => dep.name === 'kubectl')!.installed = false;
    client.dependencies.find(dep => dep.name === 'docker')!.installed = false;

    // @ts-ignore
    client.exec(['something'])

    expectClient(ctx, 1);
  })
  it('has all deps', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach(dep => dep.installed = true);

    // @ts-ignore
    client.exec(['something'])

    expectClient(ctx, -1);
  })
});