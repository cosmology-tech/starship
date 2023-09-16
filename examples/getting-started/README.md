# Getting Started

Simple self-contained example to get started with Starship.

## TLDR
```bash
```bash
cd getting-started/

# Install the starship instance and run port-forward
make install
# OR, if you are low on resources on local machine
make install-tiny

# Once the pods are running, run
make port-forward

# Stop the cluster with
make stop
```

For more detailed step-by-step approach checkout the rest of the README.

## Dependencies
Make sure you have the dependencies installed. Run the following command to install the dependencies.
```bash
make setup-deps
# Will install
# - kind
# - kubectl
# - helm
# - yq
```

Additionally, you would also need `docker` to be running on your machine as well.

## Spin up a cluster (if you don't have one already)
Optional step to setup a cluster with kind, if you already have a cluster running, you can skip this step.
```bash
make setup-kind
```

## Deploy Starship
Run the following commands to deploy starship on the cluster.
```bash
make install
# OR, if low on resources on local machine
make install-tiny
```

This will run the following nodes:
* Osmosis node
* Gaia node
* Hermes relayer between the two
* Ping Pub explorer

## Connect to cluster
Port forward ports from the cluster to your local machine.
```bash
make port-forward
```

## Interact with the chains
You can then interact with the chain on localhost at
* Osmosis: http://localhost:26653/status
* Cosmos: http://localhost:26657/status
* Ping Pub: http://localhost:8080
