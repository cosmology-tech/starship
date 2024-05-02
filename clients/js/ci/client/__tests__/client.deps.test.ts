import { createClient, expectClient } from '../test-utils/client';

describe('StarshipClient', () => {
  it('missing deps', () => {
    const { client, ctx } = createClient();

    client.dependencies = client.dependencies.map(dep=>{
      if (['kubectl', 'docker'].includes(dep.name)) {
        return {
          ...dep,
          installed: false
        };
      }
      return dep;
    });

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