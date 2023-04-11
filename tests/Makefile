HELM_NAME = starship-e2e-tests
HELM_FILE = config.yaml

HELM_REPO = starship
HELM_CHART = devnet

###############################################################################
###                              Helm Charts                                ###
###############################################################################

install:
	helm install -f $(HELM_FILE) $(HELM_NAME) ../charts/$(HELM_CHART)

upgrade:
	helm upgrade --debug $(HELM_NAME) ../charts/$(HELM_CHART) -f $(HELM_FILE)

debug:
	helm install --dry-run --debug -f $(HELM_FILE) $(HELM_NAME) ../charts/$(HELM_CHART)

delete:
	-helm delete $(HELM_NAME)

###############################################################################
###                              Port forward                              ###
###############################################################################

.PHOY: port-forward
port-forward:
	$(CURDIR)/../scripts/port-forward.sh -c=$(HELM_FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"
