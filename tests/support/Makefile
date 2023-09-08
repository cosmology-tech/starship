HELM_NAME ?= starship-support-tests
HELM_FILE ?= chains/injective.yaml

HELM_REPO ?= starship
HELM_CHART ?= devnet

###############################################################################
###                              Helm Charts                                ###
###############################################################################

.PHONY: start
start: install port-forward

.PHONY: stop
stop: stop-forward delete

###############################################################################
###                              Helm Charts                                ###
###############################################################################

install:
	bash $(CURDIR)/../../scripts/install.sh --config $(HELM_FILE) --name $(HELM_NAME) --chart ../../charts/$(HELM_CHART)

upgrade:
	helm upgrade --debug $(HELM_NAME) ../../charts/$(HELM_CHART) -f $(HELM_FILE)

debug:
	bash $(CURDIR)/../../scripts/install.sh --config $(HELM_FILE) --name $(HELM_NAME) --dry-run --chart ../../charts/$(HELM_CHART)

delete:
	-helm delete $(HELM_NAME)

###############################################################################
###                              Port forward                              ###
###############################################################################

.PHOY: port-forward
port-forward:
	$(CURDIR)/../../scripts/port-forward.sh --config=$(HELM_FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"
