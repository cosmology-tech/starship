import { createClient, expectClient } from '../test-utils/client';

describe('StarshipClient', () => {
  it('setup', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach(dep => dep.installed = true);

    client.setup();
    client.install();
    client.teardown();

    expectClient(ctx, -1);
  });
});