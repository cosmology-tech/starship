package main

import "errors"

var (
	Version         = "v0"
	RequestIdCtxKey = &contextKey{"RequestId"}
)

const (
	Prog        = "exposer"
	Description = "is a sidecar for running cosmos chain nodes for debugging"
	envPrefix   = "EXPOSER_"
)

var (
	ErrResourceInUse = errors.New("resource in use")
)

// copied and modified from net/http/http.go
// contextKey is a value for use with context.WithValue. It's used as
// a pointer, so it fits in an interface{} without allocation.
type contextKey struct {
	name string
}

func (k *contextKey) String() string { return Prog + " context value " + k.name }
