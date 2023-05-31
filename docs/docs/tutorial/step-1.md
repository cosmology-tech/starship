# **Step 1:** Setup dependencies

_You can skip directly to **Step 4** for the quickest start._

In this step, we will setup the dependencies for our project. We will use
* `kind`
* `kubectl`
* `helm`
* `yq`

## Installation
Run the following commands to get the dependencies installed.
```bash
bash <(curl -Ls https://raw.githubusercontent.com/cosmology-tech/starship/main/scripts/dev-setup.sh)
```

This will fetch all the dependencies and install them in the `~/.local/bin` directory.
Please add it to your `PATH` variable with
```bash
export PATH=$PATH:~/.local/bin
```

## Manual Installation
If you want to install the dependencies manually, please follow the instructions of each of
the required dependencis

* `kubectl`: https://kubernetes.io/docs/tasks/tools/
* `kind`: https://kind.sigs.k8s.io/docs/user/quick-start/#installation
* `helm`: https://helm.sh/docs/intro/install/
* `yq`: https://github.com/mikefarah/yq/#install
