BINARY_NAME = registry
BINARY_DIR = $(CURDIR)/../bin

DOCKER := $(shell which docker)
DOCKER_REPO_NAME := anmol1696
DOCKER_IMAGE := registry
DOCKER_TAG_NAME := $(shell date '+%Y%m%d')-$(shell git rev-parse --short HEAD)

DOCKER_ARGS += --platform linux/amd64

all: build run

build:
	bazel build //$(BINARY_NAME):$(BINARY_NAME)

run:
	bazel run //$(BINARY_NAME):$(BINARY_NAME) -- \
	--http-port "8001" --grpc-port "9001" \
	--chain-registry $(CURDIR)/testdata/ \
	--chain-client-ids "osmosis-1,juno-0" \
	--chain-client-rpcs "https://rpc.osmosis.zone:443,https://rpc.juno.strange.love:443"

# Assumes a local starship cluster is running
# Uses tests/config.yaml as the base config
run-local:
	bazel run //$(BINARY_NAME):$(BINARY_NAME) -- \
	--http-port "8001" --grpc-port "9001" \
	--chain-registry $(CURDIR)/testdata/ \
	--chain-client-ids "osmosis-1,cosmoshub-4" \
	--chain-client-rpcs "http://localhost:26653,http://localhost:26657" \
	--chain-client-exposer "http://localhost:38083,http://localhost:38087"

## Docker commands
docker-build:
	bazel build //$(BINARY_NAME):image
	bazel build //$(BINARY_NAME):image_arm64

docker-build-push:
	bazel run --define=DOCKER_TAG_NAME=$(DOCKER_TAG_NAME) //$(BINARY_NAME):image_push

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)
