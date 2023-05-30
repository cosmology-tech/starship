# Starship Documentation

The Starship project is a user-friendly interchain development environment, complete with comprehensive testing capabilities for multi-chain scenarios. 
The current problem in the Cosmos ecosystem is the increasing complexity of interchain development and, in turn, the lack of good end-to-end tests. 
More than half of the projects in the Cosmos ecosystem, dont have any end to end testing setup. Other half has tests, and sets up the environment 
in Docker containers controlled either via custom Golang or Docker compose. Most projects have their own way of doing end-to-end testing, if 
they do it at all. Dapps being built in the Cosmos ecosystem currently also dont have easy to spin up the infra setup to mimic the current 
interchain development.

## Aim
Starship aims to create an easy to setup mini interchain environment by leveraging the power of Kubernetes, a production-grade container orchestration system, 
to set up a fully simulated Cosmos ecosystem with multiple chains (each with multiple nodes), relayers, explorers and even its own chain registry. 
We provide this interface to developers via a very simplistic yaml config file: example:

```yaml
chains:
  - name: osmosis-1
    type: osmosis
    numValidators: 4
    ports:
      rest: 1317
      rpc: 26657
  - name: cosmos-2
    type: cosmos
    numValidators: 3
    ports:
      rest: 1313
      rpc: 26653
  - name: juno-0
    type: juno
    numValidators: 2
    ports:
      rest: 1315
      rpc: 26655

relayers:
  - name: osmos-juno
    type: ts-relayer
    replicas: 1
    chains:
      - osmosis-1
      - juno-0
  - name: osmos-cosmos
    type: ts-relayer
    replicas: 1
    chains:
      - osmosis-1
      - cosmos-2

explorer:
  enabled: true
```

Along with the easy to setup infra, Starship also aim to create a testing framework to get rid of bash scripts and make scenario testing easy and repeatable.
The main plan is to add regression tests for various scenarios that can be run in the CI/CD to catch bugs sooner in the development lifecycle.

Starship will be used as a development environment for projects like OsmoJS, Telescope (frontend utils/apps for interaction with chain), 
Osmosis (cosmos based chain), Mesh Security (contract development over IBC). To achieve this, we will start developing a testing framework 
and integrate Starship as part of the projects which can run various tests in the CI/CD or locally.
