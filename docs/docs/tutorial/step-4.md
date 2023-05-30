# **Step 4:** All in one place

We have created some handy example directory for the tutorial you just went though.
You can find it at [examples/tutorial](https://github.com/cosmology-tech/starship/tree/main/examples/getting-started).

You can download the example directory with this [URL](https://download-directory.github.io/?url=https%3A%2F%2Fgithub.com%2Fcosmology-tech%2Fstarship%2Ftree%2Fmain%2Fexamples%2Fgetting-started)

Unzip it into a directory and run the following commands to get started.
You should see
```bash
Makefile
configs/
  starship.yaml
  tiny-starship.yaml
scripts/
  dev-setup.sh
  port-forward.sh
README.md
```

```bash
cd getting-started/

# Install dependencies, install startship helm chart, create kind cluster
make setup

# Install the starship instance and run port-forward
make start

# Stop the cluster with
make stop
```

Once you are done, you can all the resources with
```bash
make clean
```

## Low on resources
Simulating docker containers on your machine can be resource intensive, if you dont have enough
resources you can also run a smaller version of the starship instance with less resources
```bash
# instead of `make start` run following
make start-tiny
```
Note: This will be slower to spin up as well
