TYPE = chains
PROCESS = osmosis
DOCKER := $(shell which docker)
DOCKER_REPO_NAME := anmol1696
DOCKER_TAG_NAME := latest

DOCKER_SCRIPT=$(CURDIR)/../scripts/build-docker.sh

docker-build:
	$(DOCKER_SCRIPT) -t $(TYPE) -p $(PROCESS)

docker-build-push:
	$(DOCKER_SCRIPT) -t $(TYPE) -p $(PROCESS) --push

docker-run:
	$(DOCKER) run --rm -it --entrypoint /bin/bash $(DOCKER_REPO_NAME)/$(PROCESS):$(DOCKER_TAG_NAME)

