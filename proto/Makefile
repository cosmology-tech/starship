.PHONY: build-proto
build-proto: build-exposer build-registry

.PHONY: build-registry
build-registry:
	bazel build //proto/registry:registry_go_proto

.PHONY: build-exposer
build-exposer:
	bazel build //proto/exposer:exposer_go_proto

.PHONY: update-go-pbs
update-go-pbs:
	cd $(CURDIR)/../ && bash scripts/update-go-pbs.sh
