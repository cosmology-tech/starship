build: build-collector build-exposer build-registy

build-collector:
	buf generate --path collector/ --output ../collector

build-registy:
	buf generate --path registry/ --output ../registry

build-exposer:
	buf generate --path exposer/ --output ../exposer
