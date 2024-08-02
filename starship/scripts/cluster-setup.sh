#!/bin/bash

set -euo pipefail

# Function to print colored messages
function color() {
  local color=$1
  shift
  local black=30 red=31 green=32 yellow=33 blue=34 magenta=35 cyan=36 white=37
  local color_code=${!color:-$green}
  printf "\033[%sm%s\033[0m\n" "$color_code" "$*"
}

# Function to check if a namespace exists
namespace_exists() {
    kubectl get namespace "$1" &> /dev/null
}

# Function to check if a Helm release exists
release_exists() {
    helm list --namespace "$1" -q | grep -w "$2" &> /dev/null
}

# Function to check if cert-manager is installed
cert_manager_installed() {
    kubectl get pods --namespace cert-manager &> /dev/null
}

# Function to install and setup ingress
setup_ingress() {
    if release_exists ingress nginx-ingress; then
        color yellow "Ingress is already installed. Skipping ingress setup."
    else
        color blue "Setting up ingress..."
        helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
        helm repo update
        helm install nginx-ingress ingress-nginx/ingress-nginx --namespace ingress --create-namespace
        color green "Ingress setup completed."
    fi
}

# Function to install cert-manager
setup_cert_manager() {
    if cert_manager_installed; then
        color yellow "Cert-manager is already installed. Skipping cert-manager setup."
    else
        color blue "Setting up cert-manager..."
        kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.15.2/cert-manager.yaml
        color green "Cert-manager setup completed."
    fi
}

# Main function to orchestrate setup
main() {
    color blue "Starting cluster setup..."

    # Setup ingress
    if namespace_exists ingress && release_exists ingress nginx-ingress; then
        color yellow "Ingress is already installed. Skipping ingress setup."
    else
        setup_ingress
    fi

    # Setup cert-manager
    if cert_manager_installed; then
        color yellow "Cert-manager is already installed. Skipping cert-manager setup."
    else
        setup_cert_manager
    fi

    color green "Cluster setup completed successfully."
}

main "$@"
