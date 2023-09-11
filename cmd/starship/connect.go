package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"sync"
)

// File responsible for having functions that will perform kubectl port-froward.
// for the first pass, we will just run a exec commond that runns the kubectl cmd in the
// background. We will also need to provide a function to kill this process as well

var defaultPorts = map[string]map[string]int{
	"chain": {
		"rest":    1317,
		"rpc":     26657,
		"grpc":    9091,
		"faucet":  8000,
		"exposer": 8081,
	},
	"explorer": {
		"rest": 8080,
	},
	"registry": {
		"rest": 8080,
		"grpc": 9090,
	},
}

// portForward function with perform port-forwarding based on
// kubectl port-forward <resource> <localPort>:<removePort>
func (c *Client) execPortForwardCmd(resource string, localPort, remotePort int) *exec.Cmd {
	cmdArgs := []string{"port-forward", resource, fmt.Sprintf("%v:%v", localPort, remotePort)}
	if c.config.Namespace != "" {
		cmdArgs = append(cmdArgs, []string{"-n", c.config.Namespace}...)
	}
	cmd := exec.Command("kubectl", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ() // pass through environment variables
	return cmd
}

// PortForwardCmds returns a list of exec commands for port-fowrarding from a config file
func (c *Client) PortForwardCmds() ([]*exec.Cmd, []string, error) {
	config := c.helmConfig

	var cmds []*exec.Cmd
	var msgs []string

	// port-forward all chains
	for _, chain := range config.Chains {
		for portType, remotePort := range defaultPorts["chain"] {
			port := chain.Ports.GetPort(portType)
			if port == 0 {
				continue
			}
			msgs = append(msgs, fmt.Sprintf("port-forwarding: %s: port %s: to: http://localhost:%d", chain.Name, portType, port))
			cmds = append(cmds, c.execPortForwardCmd(fmt.Sprintf("pods/%s-genesis-0", chain.Name), port, remotePort))
		}
	}
	// port-forward explorer
	if config.Explorer != nil {
		for portType, remotePort := range defaultPorts["explorer"] {
			port := config.Explorer.Ports.GetPort(portType)
			if port == 0 {
				continue
			}
			msgs = append(msgs, fmt.Sprintf("port-forwarding: %s: port %s: to: http://localhost:%d", "explorer", portType, port))
			cmds = append(cmds, c.execPortForwardCmd("svc/explorer", port, remotePort))
		}
	}
	// port-forward registry
	if config.Registry != nil {
		for portType, remotePort := range defaultPorts["registry"] {
			port := config.Registry.Ports.GetPort(portType)
			if port == 0 {
				continue
			}
			msgs = append(msgs, fmt.Sprintf("port-forwarding: %s: port %s: to: http://localhost:%d", "registry", portType, port))
			cmds = append(cmds, c.execPortForwardCmd("svc/registry", port, remotePort))
		}
	}

	return cmds, msgs, nil
}

// RunPortForward function performs the exec commands to run the port-forwarding
func (c *Client) RunPortForward(cliCtx context.Context) error {
	cmds, msgs, err := c.PortForwardCmds()
	if err != nil {
		return err
	}

	// run all cmds in parallel go-routines
	ctx, cancel := context.WithCancel(cliCtx)
	defer cancel()

	var wg sync.WaitGroup
	resultChan := make(chan error, len(cmds))

	c.logger.Debug("commands to run", zap.Int("num_cmds", len(cmds)))
	for _, cmd := range cmds {
		wg.Add(1)
		c.logger.Debug(fmt.Sprintf("runing cmd: %s", cmd.String()))
		go func(ctx context.Context, cmd *exec.Cmd, g *sync.WaitGroup, resultChan chan error) {
			if err := cmd.Run(); err != nil {
				select {
				case resultChan <- err:
				case <-ctx.Done():
				}
			}
		}(ctx, cmd, &wg, resultChan)
	}

	// log port forwarded endpoints to localhost
	c.logger.Info("Port forwarding")
	for _, msg := range msgs {
		c.logger.Info(msg)
	}

	// Start a goroutine to monitor the result channel and cancel if any error occurs
	go func() {
		for err := range resultChan {
			fmt.Printf("Command failed: %v\n", err)
			cancel()
		}
	}()

	// Wait for all commands to finish
	wg.Wait()

	// Close the result channel
	close(resultChan)

	c.logger.Debug("port-forward in progress... exit with Ctlr+C")

	return nil
}

// CheckPortForward verfify if all pods are in running state, and ready to be port-forwarded
func (c *Client) CheckPortForward() error {
	return nil
}
