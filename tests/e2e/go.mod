module github.com/cosmology-tech/starship/tests/e2e

go 1.19

require (
	github.com/cosmology-tech/starship/exposer v0.0.0-20230413092908-7da9e8a24b31
	github.com/cosmology-tech/starship/registry v0.0.0-20230411094226-129001b2f52a
	github.com/golang/protobuf v1.5.3
	github.com/stretchr/testify v1.8.2
	go.uber.org/zap v1.24.0
	google.golang.org/protobuf v1.29.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230320184635-7606e756e683 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace (
	github.com/cosmology-tech/starship/exposer => ../../exposer
	github.com/cosmology-tech/starship/registry => ../../registry
)
