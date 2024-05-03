package main

import (
	"errors"
)

var (
	Version         = "v0"
	RequestIdCtxKey = &contextKey{"RequestId"}
)

const (
	Prog        = "registry"
	Description = "is a local chain registry api that exposes chain registry based on configs"
	envPrefix   = "REGISTRY_"
)

// Define default errors
var (
	ErrValidation       = errors.New("validation error")
	ErrNotFound         = errors.New("resource not found")
	ErrNotImplemented   = errors.New("not Implemented")
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrInternalServer   = errors.New("internal server error")
)

// copied and modified from net/http/http.go
// contextKey is a value for use with context.WithValue. It's used as
// a pointer, so it fits in an interface{} without allocation.
type contextKey struct {
	name string
}

func (k *contextKey) String() string { return Prog + " context value " + k.name }
