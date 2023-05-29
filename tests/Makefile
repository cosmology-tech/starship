HELM_NAME ?= starship-e2e-tests
HELM_FILE ?= configs/two-chain.yaml

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
	helm install -f $(HELM_FILE) $(HELM_NAME) ../charts/$(HELM_CHART) --wait --debug

upgrade:
	helm upgrade --debug $(HELM_NAME) ../charts/$(HELM_CHART) -f $(HELM_FILE)

debug:
	helm install --dry-run --debug -f $(HELM_FILE) $(HELM_NAME) ../charts/$(HELM_CHART)

delete:
	-helm delete $(HELM_NAME)

###############################################################################
###                              Tests cases                                ###
###############################################################################

.PHONY: test
test:
	bazel test --cache_test_results=no --test_output=all --test_env=TEST_CONFIG_FILE=$(HELM_FILE) //tests:tests

###############################################################################
###                              Port forward                              ###
###############################################################################

.PHOY: port-forward
port-forward:
	$(CURDIR)/../scripts/port-forward.sh --config=$(HELM_FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"
