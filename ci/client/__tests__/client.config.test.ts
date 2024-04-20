import { join, relative } from 'path';

import { createClient, expectClient } from '../test-utils/client';
import { config, outputDir } from '../test-utils/config';

describe('StarshipClient', () => {
  it('setup', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach(dep => dep.installed = true);

    client.setConfig(config.config);
    const helmFile = client.ctx.helmFile;
    client.ctx.helmFile = join(outputDir, 'my-config.yaml');
    client.ctx.helmFile = relative(process.cwd(), client.ctx.helmChart)
    client.saveConfig();
    client.ctx.helmFile = helmFile;

    const portYaml = join(outputDir, 'default-pod-ports.yaml');
    const relativePortYaml = relative(process.cwd(), portYaml);
    client.savePodPorts(relativePortYaml);

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