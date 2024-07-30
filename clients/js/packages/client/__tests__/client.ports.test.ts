import { join, relative } from 'path';

import { createClient, expectClient } from '../test-utils/client';
import { config, outputDir } from '../test-utils/config';

describe('StarshipClient', () => {
  it('setup', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach((dep) => (dep.installed = true));

    client.setConfig(config.config);
    client.setPodPorts({
      chains: {
        osmosis: {
          exposer: 98988,
          faucet: 1000000,
          grpc: 909090,
          rest: 6767676
        }
      }
    });
    const portYaml = join(outputDir, 'custom-pod-ports.yaml');
    const relativePortYaml = relative(process.cwd(), portYaml);
    client.savePodPorts(relativePortYaml);

    // helm
    client.setup();
    client.deploy();

    client.startPortForward();

    // remove helm chart
    client.stop();

    expectClient(ctx, -1);
  });
});
