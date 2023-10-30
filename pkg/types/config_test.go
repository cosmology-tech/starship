package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHostPorts(t *testing.T) {
	baseInput := HostPort{
		Rest:    1317,
		Rpc:     26657,
		Grpc:    9090,
		Exposer: 8000,
		Faucet:  8001,
	}

	testCases := map[string]struct {
		input    HostPort
		expected map[string]int
	}{
		"happy-path": {
			input: baseInput,
			expected: map[string]int{
				"rest":        1317,
				"rpc":         26657,
				"grpc":        9090,
				"exposer":     8000,
				"faucet":      8001,
				"notexistent": 0,
			},
		},
		"partial-ports": {
			input: HostPort{Rpc: 26657},
			expected: map[string]int{
				"rpc":          26657,
				"grpc":         0,
				"rest":         0,
				"faucet":       0,
				"nonexsistent": 0,
			},
		},
	}

	for _, testCase := range testCases {
		for portType, expectedPort := range testCase.expected {
			port := testCase.input.GetPort(portType)
			require.Equal(t, expectedPort, port)
		}
	}
}
