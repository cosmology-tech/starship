EXAMPLE = multi-chain
FILE = config.yaml

HELM_REPO = starship
HELM_CHART = devnet

setup:
	helm repo add $(HELM_REPO) https://hyperweb-io.github.io/starship/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART)

install:
	helm install -f $(EXAMPLE)/$(FILE) $(EXAMPLE) ../charts/$(HELM_CHART)

upgrade:
	helm upgrade --debug $(EXAMPLE) ../charts/$(HELM_CHART) -f $(EXAMPLE)/$(FILE)

debug:
	helm install --dry-run --debug -f $(EXAMPLE)/$(FILE) $(EXAMPLE) ../charts/$(HELM_CHART)

install-local:
	$(MAKE) install FILE=local-values.yaml

install-custom:
	$(MAKE) install FILE=custom-values.yaml

delete:
	-helm delete $(EXAMPLE)

###############################################################################
###                              Port forward                              ###
###############################################################################

.PHOY: port-forward-all
port-forward-all:
	bash $(CURDIR)/../scripts/port-forward.sh -c=$(FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"

