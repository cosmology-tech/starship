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
	--chain-registry $(CURDIR)/testdata/ \
	--chain-client-ids "osmosis-1" \
	--chain-client-rpcs "https://rpc.osmosis.zone:443"

## Docker commands
docker-build:
	bazel build //$(BINARY_NAME):image
	bazel build //$(BINARY_NAME):image_arm64

docker-build-push:
	bazel run --define=DOCKER_TAG_NAME=$(DOCKER_TAG_NAME) //$(BINARY_NAME):image_push

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)
