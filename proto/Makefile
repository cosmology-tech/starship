.PHONY: build-proto
build-proto: build-exposer build-registry

.PHONY: build-registry
build-registry:
	buf generate --path $(CURDIR)/registry/ --output $(CURDIR)/../registry

.PHONY: build-exposer
build-exposer:
	buf generate --path $(CURDIR)/exposer/ --output $(CURDIR)/../exposer