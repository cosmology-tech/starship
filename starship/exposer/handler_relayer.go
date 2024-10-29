package main

import (
	"context"
	"fmt"
	pb "github.com/cosmology-tech/starship/exposer/exposer"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

// CreateChannel function runs the hermes command to create a channel between 2 given chains
func (a *AppServer) CreateChannel(ctx context.Context, req *pb.RequestCreateChannel) (*pb.ResponseCreateChannel, error) {
	ok := a.mu.TryLock()
	if !ok {
		return nil, ErrResourceInUse
	}
	defer a.mu.Unlock()

	// note: for cli one time txn, use `config-cli.toml` file
	createCmd := fmt.Sprintf("hermes --config /root/.hermes/config-cli.toml create channel --a-chain %s --a-port %s --b-port %s --yes", req.AChain, req.APort, req.BPort)

	if req.AConnection != nil {
		createCmd += fmt.Sprintf(" --a-connection %s", *req.AConnection)
	} else {
		createCmd += " --new-connection"
	}
	if req.ChannelVersion != nil {
		createCmd += fmt.Sprintf(" --channel-version %s", *req.ChannelVersion)
	}
	if req.Order != nil {
		createCmd += fmt.Sprintf(" --order %s", *req.Order)
	}
	a.logger.Debug("running command:", zap.String("cmd", createCmd))

	output, err := runCommand(createCmd)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseCreateChannel{Status: string(output)}, nil
}

func (a *AppServer) getClients(chainID string) ([]string, error) {
	getCmd := fmt.Sprintf("hermes query clients --host-chain %s", chainID)
	a.logger.Debug("running command:", zap.String("cmd", getCmd))

	output, err := runCommand(getCmd)
	if err != nil {
		return nil, err
	}
	a.logger.Debug("output from get clients:", zap.ByteString("cmdOutput", output))

	// Regular expression to match the client_id in the output
	re := regexp.MustCompile(`(?m)^\s*client_id:\s*ClientId\(\s*"([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(output), -1)

	var clients []string
	for _, match := range matches {
		if len(match) > 1 {
			clients = append(clients, match[1])
		}
	}

	return clients, nil
}

// GetClients function runs the hermes command to get the clients of a given chain
func (a *AppServer) GetClients(ctx context.Context, req *pb.RequestGetClients) (*pb.ResponseGetClients, error) {
	clients, err := a.getClients(req.ChainId)
	if err != nil {
		return nil, err
	}

	// Check if clientIds are extracted correctly
	a.logger.Debug("parsed client IDs:", zap.Strings("client_ids", clients))

	// Return the parsed client IDs in a structured response
	return &pb.ResponseGetClients{Clients: clients}, nil
}

func (a *AppServer) updateClient(ctx context.Context, chainID, clientID string) (string, error) {
	// note: for cli one time txn, use `config-cli.toml` file
	updateCmd := fmt.Sprintf("hermes --config /root/.hermes/config-cli.toml update client --host-chain %s --client-id %s", chainID, clientID)
	a.logger.Debug("running command:", zap.String("cmd", updateCmd))

	output, err := runCommand(updateCmd)
	if err != nil {
		return "", err
	}
	// Check if clientIds are extracted correctly
	a.logger.Info("output from update client:",
		zap.String("chain", chainID),
		zap.String("client", clientID),
		zap.ByteString("cmdOutput", output))

	return string(output), nil
}

// UpdateClients function runs the herjson command to update a client of a given chain
func (a *AppServer) UpdateClients(ctx context.Context, req *pb.RequestUpdateClients) (*pb.ResponseUpdateClients, error) {
	clients, err := a.getClients(req.ChainId)
	if err != nil {
		return nil, err
	}

	for _, clientID := range clients {
		_, err = a.updateClient(ctx, req.ChainId, clientID)
		if err != nil {
			return nil, err
		}
	}

	return &pb.ResponseUpdateClients{Status: "ok"}, nil
}

// UpdateAllClients function runs the hermes command to update all clients of a given chain
func (a *AppServer) UpdateAllClients(ctx context.Context, req *pb.RequestUpdateAllClients) (*pb.ResponseUpdateAllClients, error) {
	ok := a.mu.TryLock()
	if !ok {
		return nil, ErrResourceInUse
	}
	defer a.mu.Unlock()

	chainIDs := strings.Split(a.config.ChainIDs, ",")
	for _, chainID := range chainIDs {
		clients, err := a.getClients(chainID)
		if err != nil {
			return nil, err
		}

		for _, client := range clients {
			_, err = a.updateClient(ctx, chainID, client)
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.ResponseUpdateAllClients{Status: "ok"}, nil
}
