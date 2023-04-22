HELM_NAME = update-test
HELM_FILE = config.yaml

HELM_REPO = starship
HELM_CHART = devnet

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
	helm install -f $(HELM_FILE) $(HELM_NAME) ../../charts/$(HELM_CHART) --wait --debug --timeout=20m

upgrade:
	helm upgrade --debug $(HELM_NAME) ../../charts/$(HELM_CHART) -f $(HELM_FILE)

debug:
	helm install --dry-run --debug -f $(HELM_FILE) $(HELM_NAME) ../../charts/$(HELM_CHART)

delete:
	-helm delete $(HELM_NAME)

###############################################################################
###                              Tests cases                                ###
###############################################################################

.PHONY: test
test:
	go test .


###############################################################################
###                              Port forward                              ###
###############################################################################

.PHOY: port-forward
port-forward:
	$(CURDIR)/../../scripts/port-forward.sh --config=$(HELM_FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"
