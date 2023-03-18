PROCESS = osmosis
DOCKER := $(shell which docker)
DOCKER_REPO_NAME := anmol1696
DOCKER_TAG_NAME := latest

DOCKER_ARGS += --platform linux/amd64

docker-build:
	$(DOCKER) buildx build $(DOCKER_ARGS) \
		-t $(DOCKER_REPO_NAME)/$(PROCESS):$(DOCKER_TAG_NAME) \
		. -f $(PROCESS)/Dockerfile

docker-build-arm64:
	$(DOCKER) buildx build --platform linux/arm64 \
		-t $(DOCKER_REPO_NAME)/$(PROCESS):$(DOCKER_TAG_NAME)-arm64 \
		. -f $(PROCESS)/Dockerfile.arm64

docker-build-push: docker-build
	$(DOCKER) push $(DOCKER_REPO_NAME)/$(PROCESS):$(DOCKER_TAG_NAME)

docker-build-push-arm64: docker-build-arm64
	$(DOCKER) push $(DOCKER_REPO_NAME)/$(PROCESS):$(DOCKER_TAG_NAME)-arm64

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(PROCESS):$(DOCKER_TAG_NAME)

## Handy commands for all
process_dir = .
docker-build-push-all: $(process_dir)/*
	for p in $^ ; do \
  	$(MAKE) docker-build-push PROCESS=$${p}; \
	done
