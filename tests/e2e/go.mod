module github.com/cosmology-tech/starship/tests/e2e

go 1.21

require (
	github.com/cosmology-tech/starship/exposer v0.0.0-20230413092908-7da9e8a24b31
	github.com/cosmology-tech/starship/registry v0.0.0-20230411094226-129001b2f52a
	github.com/golang/protobuf v1.5.3
	github.com/stretchr/testify v1.8.4
	go.uber.org/zap v1.26.0
	google.golang.org/protobuf v1.32.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/grpc v1.60.1 // indirect
)

replace (
	github.com/cosmology-tech/starship/exposer => ../../exposer
	github.com/cosmology-tech/starship/registry => ../../registry
)
