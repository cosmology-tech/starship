#!/bin/bash

set -euo pipefail

function color() {
  local color=$1
  shift
  local black=30 red=31 green=32 yellow=33 blue=34 magenta=35 cyan=36 white=37
  local color_code=${!color:-$green}
  printf "\033[%sm%s\033[0m\n" "$color_code" "$*"
}

# Define a function to install a binary on macOS
install_macos() {
    case $1 in
        kubectl) brew install kubectl ;;
        helm) brew install helm ;;
        yq) brew install yq ;;
        kind) brew install kind ;;
    esac
}

# Define a function to install a binary on Linux
install_linux() {
    case $1 in
        kubectl) curl -LO "https://dl.k8s.io/release/$(curl -Ls https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" -o /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl ;;
        helm) curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash ;;
        yq) curl -L "https://github.com/mikefarah/yq/releases/download/v4.33.3/yq_linux_amd64" -o /usr/local/bin/yq && chmod +x /usr/local/bin/yq ;;
        kind) curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.18.1/kind-linux-amd64 && chmod +x ./kind && mv ./kind /usr/local/bin/kind ;;
    esac
}

# Define a function to install a binary
install_binary() {
    if [[ $(uname -s) == "Darwin" ]]; then
        install_macos $1
    else
        install_linux $1
    fi
}

# Define a function to check for the presence of a binary
check_binary() {
    if ! command -v $1 &> /dev/null
    then
        echo "$1 is not installed"
        if [ "$EUID" -ne 0 ]; then
            color yellow "Please run the script with sudo to install $1"
            exit 1
        fi
        install_binary $1
        if ! command -v $1 &> /dev/null
        then
            color red "Installation of $1 failed, exiting..."
            exit 1
        fi
    fi
}

# Check the binaries
check_binary kubectl
check_binary helm
check_binary yq
check_binary kind

color green "All binaries are installed"
