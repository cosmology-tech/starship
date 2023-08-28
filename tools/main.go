// tools adds a blank import to tools we use such that `go mod tidy`
// doesn't clean up needed dependencies when running `go install`.

//go:build tools
// +build tools

package main

import (
	"fmt"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)

func main() {
	fmt.Printf("You just lost the game\n")
}
