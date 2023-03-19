# Infra

Starship infra is built with kubernetes. Inorder to simulate the cosmos ecosystem, we run multiple type of nodes, all of them in a scalable fashion.

## Design

![k8s setup](https://raw.githubusercontent.com/Anmol1696/starship/main/docs/docs/assets/images/k8s-setup.png "Kubernetes setup")

We support multiple kinds of nodes:

* Chains

    * Cosmoshub
    * Osmosis
    * Persistence
    * Juno
    * Generic Wasmd chain

* IBC relayer
    * TS-relayer
    * Hermes

* Explorers
    * Ping Pub

Soon we will also add ability to add:

* RPC Nodes
* Archival Nodes
