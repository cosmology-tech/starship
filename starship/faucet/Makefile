BINARY_NAME = faucet

DOCKER := $(shell which docker)
# DOCKER_REPO_NAME is the local docker repo used, can be set to individual dockerhub username
DOCKER_REPO_NAME := starship
DOCKER_IMAGE := faucet
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
	GOOS=linux GOARCH=arm64 $(MAKE) build

## Need to be running osmosis node at localhost:1313 and 26653
run-docker: build-arm
	docker run --rm -v $(CURDIR)/build:/build -p 8800:8000 -it ghcr.io/hyperweb-io/starship/osmosis:v15.1.0 /build/faucet \
		--mnemonic="razor dog gown public private couple ecology paper flee connect local robot diamond stay rude join sound win ribbon soup kidney glass robot vehicle" \
		--chain-binary="osmosisd" \
		--concurrency=11 \
		--credit-coins="10000000uosmo,10000000uion" \
		--chain-id="osmosis-1" \
		--chain-rest-endpoint="http://host.docker.internal:1313" \
		--chain-rpc-endpoint="http://host.docker.internal:26653" \
		--chain-fees="10000uosmo"

test-credit:
	curl --header "Content-Type: application/json" \
        --request POST --data '{"denom":"uosmo","address":"osmo1tkrwspedcqwm6ve8vtk0vuzgr4c0m5203era0x"}' http://localhost:8800/credit

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
