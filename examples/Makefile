EXAMPLE = mesh-security
FILE = config.yaml
DEV_FILE = config-dev.yaml

HELM_REPO = starship
HELM_CHART = devnet

setup:
	helm repo add $(HELM_REPO) https://cosmology-tech.github.io/starship/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART)

install:
	helm install -f $(EXAMPLE)/$(FILE) $(EXAMPLE) ../charts/$(HELM_CHART)

.PHONY: install-dev
install-dev:
	helm install -f $(EXAMPLE)/$(DEV_FILE) $(EXAMPLE) ../charts/$(HELM_CHART)

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
	bash $(CURDIR)/../scripts/port-forward.sh -c=$(EXAMPLE)/$(FILE)

.PHOY: port-forward-dev
port-forward-dev:
	bash $(CURDIR)/../scripts/port-forward.sh -c=$(EXAMPLE)/$(DEV_FILE)


.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"
