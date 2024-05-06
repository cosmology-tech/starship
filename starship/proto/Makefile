all: build

.PHONY: build
build:
	buf generate --path registry/ -o $(CURDIR)/../registry
	buf generate --path exposer/ -o $(CURDIR)/../exposer
	buf generate --path faucet/ -o $(CURDIR)/../faucet

.PHONY: docker-build
docker-build:
	docker run --volume "$(CURDIR):/workspace" --workdir /workspace bufbuild/buf generate -o build/
	rm -rf $(CURDIR)/../registry/registry && mv -f build/registry $(CURDIR)/../registry
	rm -rf $(CURDIR)/../exposer/exposer && mv -f build/exposer $(CURDIR)/../exposer
	rm -rf $(CURDIR)/../faucet/faucet && mv -f build/faucet $(CURDIR)/../faucet
