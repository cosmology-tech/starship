import { createClient, expectClient } from '../test-utils/client';
import { config } from '../test-utils/config';

describe('StarshipClient', () => {
  it('setup', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach((dep) => (dep.installed = true));

    client.setConfig(config.config);

    // helm
    client.setup();
    client.deploy();

    client.startPortForward();

    // remove helm chart
    client.stop();

    expectClient(ctx, -1);
  });
});
