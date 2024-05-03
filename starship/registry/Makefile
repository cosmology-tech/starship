BINARY_NAME = registry

DOCKER := $(shell which docker)
# DOCKER_REPO_NAME is the local docker repo used, can be set to individual dockerhub username
DOCKER_REPO_NAME := starship
DOCKER_IMAGE := registry
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

# Assumes a local starship cluster is running
# Uses tests/config.yaml as the base config
.PHONY: run-local
run-local:
	go run ./... \
	--http-port "8001" --grpc-port "9001" \
	--chain-registry $(CURDIR)/testdata/ \
	--chain-client-ids "osmosis-1,cosmoshub-4" \
	--chain-client-rpcs "http://localhost:26653,http://localhost:26657" \
	--chain-client-exposer "http://localhost:38083,http://localhost:38087"

## Docker commands
docker-setup:
	-docker buildx rm starship
	docker buildx create --use --name starship

docker-build:
	$(DOCKER) buildx build . --platform linux/amd64,linux/arm64 -t $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)

docker-build-push:
	$(DOCKER) buildx build . --platform linux/amd64,linux/arm64 -t $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME) --push

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)
