#!/bin/bash

# we use the starship `Dockerfile` to take the bin to build this to be used in test
# others can be found https://github.com/orgs/cosmology-tech/packages
docker buildx build \
      --platform linux/amd64,linux/arm64 \
      -t "ghcr.io/nymlab/cheqd-node:v2.0.0-starship" \
      . -f Dockerfile \
      --build-arg BASE_IMAGE="ghcr.io/nymlab/cheqd-node" \
      --build-arg VERSION="v2.0.0" \
      --build-arg IMAGE_SOURCE="https://github.com/nymlab/cheqd-node-starship" \
	  --push

