package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
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
func execPortForward(resource string, localPort, remotePort int) error {
	cmd := exec.Command("kubectl", "port-forward", resource, fmt.Sprintf("%v:%v", localPort, remotePort))
	//cmd := exec.Command("kubectl", "get", resource)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	fmt.Println("out:", outb.String(), "err:", errb.String())
	if err != nil {
		return err
	}
	return nil
}

// PortForward function performs the exec commands to run the port-forwarding
func (c *Client) PortForward() error {
	config := c.helmConfig
	// port-forward all chains
	for _, chain := range config.Chains {
		for portType, remotePort := range defaultPorts["chain"] {
			port := chain.Ports.GetPort(portType)
			if port == 0 {
				continue
			}
			c.logger.Debug(fmt.Sprintf("port-forwarding: chain %s, %s to http://localhost:%v", chain.Name, portType, port))
			err := execPortForward(fmt.Sprintf("pods/%s-genesis-0", chain.Name), port, remotePort)
			if err != nil {
				return err
			}
			c.logger.Info(fmt.Sprintf("port-forwarded: chain %s, %s to http://localhost:%v", chain.Name, portType, port))
		}
	}
	// port-forward explorer
	if config.Explorer != nil {
		for portType, remotePort := range defaultPorts["explorer"] {
			port := config.Explorer.Ports.GetPort(portType)
			if port == 0 {
				continue
			}
			err := execPortForward("svc/explorer", port, remotePort)
			if err != nil {
				return err
			}
			c.logger.Info(fmt.Sprintf("port-forwarding: explorer, %s to http://localhost:%v", portType, port))
		}
	}
	// port-forward registry
	if config.Registry != nil {
		for portType, remotePort := range defaultPorts["registry"] {
			port := config.Registry.Ports.GetPort(portType)
			if port == 0 {
				continue
			}
			err := execPortForward("svc/registry", port, remotePort)
			if err != nil {
				return err
			}
			c.logger.Info(fmt.Sprintf("port-forwarding: registry, %s to http://localhost:%v", portType, port))
		}
	}
	time.Sleep(120 * time.Second)
	return nil
}

// StopPortForward will kill the processes that ran the kubectl exec commands
func (c *Client) StopPortForward(config HelmConfig) error {
	cmd := exec.Command("pkill", "-f", "port-forward")
	return cmd.Run()
}
