import { createClient, expectClient } from '../test-utils/client';
import { config } from '../test-utils/config';

describe('StarshipClient', () => {
  it('setup', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach(dep => dep.installed = true);

    client.setConfig(config.config);

    // helm
    client.setup();
    client.install();
    
    client.startPortForward();

    client.stop();

    // remove helm chart
    client.teardown();

    expectClient(ctx, -1);
  });
});