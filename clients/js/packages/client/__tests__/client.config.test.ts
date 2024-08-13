import { join, relative } from 'path';

import { createClient, expectClient } from '../test-utils/client';
import { config, outputDir } from '../test-utils/config';

describe('StarshipClient', () => {
  it('setup', () => {
    const { client, ctx } = createClient();

    client.dependencies.forEach((dep) => (dep.installed = true));

    client.setConfig(config.config);
    const helmFile = client.ctx.config;
    client.ctx.config = join(outputDir, 'my-config.yaml');
    client.ctx.config = relative(process.cwd(), client.ctx.chart);
    // @ts-expect-error - Ignore lint error
    client.saveYaml = () => {};
    client.saveConfig();
    client.ctx.config = helmFile;

    const portYaml = join(outputDir, 'default-pod-ports.yaml');
    const relativePortYaml = relative(process.cwd(), portYaml);
    client.savePodPorts(relativePortYaml);

    // helm
    client.setup();
    client.deploy();

    client.startPortForward();

    client.stop();

    expectClient(ctx, -1);
  });
});
