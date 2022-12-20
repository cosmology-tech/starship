DOCKER := $(shell which docker)
DOCKER_REPO_NAME := anmol1696
DOCKER_IMAGE := exposer
DOCKER_TAG_NAME := latest

DOCKER_ARGS += --platform linux/amd64

docker-build:
	$(DOCKER) buildx build $(DOCKER_ARGS) \
		-t $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME) .

docker-build-push: docker-build
	$(DOCKER) push $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(DOCKER_IMAGE):$(DOCKER_TAG_NAME)
