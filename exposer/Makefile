BINARY_NAME = exposer
BINARY_DIR = $(CURDIR)/../bin

DOCKER := $(shell which docker)
# DOCKER_REPO_NAME is the local docker repo used, can be set to individual dockerhub username
DOCKER_REPO_NAME := starship
DOCKER_IMAGE := exposer
DOCKER_TAG_NAME := $(shell date '+%Y%m%d')-$(shell git rev-parse --short HEAD)

all: build run

.PHONY: build
build:
	CGO_ENABLED=0 go build -mod=readonly -o $(CURDIR)/build/ ./...

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

.PHONY: build-arm
build-arm:
	GOOS=darwin GOARCH=arm64 $(MAKE) build

run-local:
	go run ./... \
		--genesis-file $(CURDIR)/testdata/genesis.json \
		--mnemonic-file $(CURDIR)/testdata/keys.json \
		--priv-val-file $(CURDIR)/testdata/priv_validator_key.json \
		--status-url https://rpc.osmosis.zone/status

## Docker commands
docker-setup:
	docker buildx create --use --name starship

docker-build:
	$(DOCKER) buildx build . --platform linux/amd64,linux/arm64 -t $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)

docker-build-push:
	$(DOCKER) buildx build . --platform linux/amd64,linux/arm64 -t $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME) --push

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)
