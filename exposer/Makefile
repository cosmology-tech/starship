BINARY_NAME = exposer
BINARY_DIR = $(CURDIR)/../bin

DOCKER := $(shell which docker)
DOCKER_REPO_NAME := anmol1696
DOCKER_IMAGE := exposer
DOCKER_TAG_NAME := $(shell date '+%Y%m%d')-$(shell git rev-parse --short HEAD)

DOCKER_ARGS += --platform linux/amd64

all: build run

build:
	bazel build //$(BINARY_NAME):$(BINARY_NAME)

run:
	bazel run //$(BINARY_NAME):$(BINARY_NAME) -- \
		--genesis-file $(CURDIR)/testdata/genesis.json \
		--mnemonic-file $(CURDIR)/testdata/keys.json \
		--priv-val-file $(CURDIR)/testdata/priv_validator_key.json \
		--status-url https://rpc.osmosis.zone/status

## Docker commands
docker-build:
	$(DOCKER) buildx build $(DOCKER_ARGS) \
		-t $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME) .
	echo "Docker image: $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)"

docker-build-push: docker-build
	$(DOCKER) push $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)
