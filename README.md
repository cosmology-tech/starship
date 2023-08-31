# Starship

<p align="center" width="100%">
    <img height="148" src="https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png" />
</p>

<p align="center" width="100%">
   <a href="https://github.com/cosmology-tech/starship/releases/latest"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/release.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/blob/main/LICENSE"><img height="20" src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/e2e-tests.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/e2e-tests.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/build.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/build.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/docs.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/docs.yaml/badge.svg"></a>
   <a href="https://github.com/cosmology-tech/starship/actions/workflows/starship-docker.yaml"><img height="20" src="https://github.com/cosmology-tech/starship/actions/workflows/starship-docker.yaml/badge.svg"></a>
</p>

Universal interchain development environment in k8s. The vision of this project
is to have a single easy to use developer environment with full testing support
for multichain use cases

## Installation
In order to get started with starship, one needs to install the following
* `kubectl`: https://kubernetes.io/docs/tasks/tools/
* `kind`: https://kind.sigs.k8s.io/docs/user/quick-start/#installation
* `helm`: https://helm.sh/docs/intro/install/
* `jq`: https://stedolan.github.io/jq/download/
* `yq`: https://github.com/mikefarah/yq/#install

## Getting started
Follow the steps here: https://starship.cosmology.tech

## Using helm chart
Inorder to use the helm chart externally without this repo.
```bash
helm repo add starship https://cosmology-tech.github.io/starship
helm repo update

helm search repo starship/devnet
```
Fetch the values.yaml file and update them before installing the chart
```bash
helm show values starship/devnet > custom-vaules.yaml
# change custom-values.yaml file

helm install -f custom-values.yaml starship/devnet --generate-name
```

**NOTE: It is recommended to still copy the Makefile from the repo to use the handy commands**
