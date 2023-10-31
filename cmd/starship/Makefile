DOCKER := $(shell which docker)
# DOCKER_REPO_NAME is the local docker repo used, can be set to individual dockerhub username
DOCKER_REPO_NAME := starship
DOCKER_IMAGE := registry
DOCKER_TAG_NAME := $(shell date '+%Y%m%d')-$(shell git rev-parse --short HEAD)

all: build run

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(CURDIR)/build/ ./...

RUN_ARGS = "../../tests/e2e/configs/two-chain.yaml"
.PHONY: run
run:
	CGO_ENABLED=0 go run -mod=readonly ./... generate $(RUN_ARGS)

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

.PHONY: build-arm
build-arm:
	GOOS=darwin GOARCH=arm64 $(MAKE) build
