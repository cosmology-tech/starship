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

func TestMergeChains(t *testing.T) {
	testCases := map[string]struct {
		chain1   *Chain
		chain2   *Chain
		expected *Chain // merged chain
	}{
		"non-conflicting": {
			chain1: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
			},
			chain2: &Chain{
				Type: "osmosis",
				Home: "/root/.osmosisd",
			},
			expected: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Home: "/root/.osmosisd",
			},
		},
		"override-values": {
			chain1: &Chain{
				Name: "osmosis-1",
			},
			chain2: &Chain{
				Name: "osmosis-2",
			},
			expected: &Chain{
				Name: "osmosis-1",
			},
		},
		"num-override": {
			chain1: &Chain{
				Name:          "osmosis-1",
				Type:          "osmosis",
				NumValidators: 2,
			},
			chain2: &Chain{
				Name:          "osmosis-1",
				Type:          "osmosis",
				NumValidators: 1,
			},
			expected: &Chain{
				Name:          "osmosis-1",
				Type:          "osmosis",
				NumValidators: 2,
			},
		},
		"bool-overrides": {
			chain1: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Exposer: &Feature{
					Enabled: true,
				},
			},
			chain2: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
			},
			expected: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Exposer: &Feature{
					Enabled: true,
				},
			},
		},
		"reverse-bool-override": {
			chain1: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
			},
			chain2: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Exposer: &Feature{
					Enabled: true,
				},
			},
			expected: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Exposer: &Feature{
					Enabled: true,
				},
			},
		},
		"explicit-overrides": {
			chain1: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Exposer: &Feature{
					Enabled: false,
				},
			},
			chain2: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Exposer: &Feature{
					Enabled: true,
					Image:   "test-image",
				},
			},
			expected: &Chain{
				Name: "osmosis-1",
				Type: "osmosis",
				Exposer: &Feature{
					Enabled: false,
					Image:   "test-image",
				},
			},
		},
	}

	for _, testCase := range testCases {
		merged, err := testCase.chain1.Merge(testCase.chain2)
		require.NoError(t, err)
		require.Equal(t, testCase.expected, merged)
	}
}
