import { createClient, expectClient } from '../test-utils/client';
import { config } from '../test-utils/config';

describe('StarshipClient', () => {
  it('setup', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach(dep => dep.installed = true);

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

    // helm
    client.setup();
    client.deploy();
    
    client.startPortForward();

    client.undeploy();

    // remove helm chart
    client.teardown();

    expectClient(ctx, -1);
  });
});