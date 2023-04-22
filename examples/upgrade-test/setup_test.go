package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	expb "github.com/cosmology-tech/starship/exposer/exposer"
)

var configFile = "./config.yaml"

type TestSuite struct {
	suite.Suite

	config       *Config
	chainClients ChainClients
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	s.T().Log("setting up e2e integration test suite...")

	// read config file from yaml
	yamlFile, err := os.ReadFile(configFile)
	s.Require().NoError(err)
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	s.Require().NoError(err)
	s.config = config

	// create chain clients
	chainClients, err := NewChainClients(zap.L(), config)
	s.Require().NoError(err)
	s.chainClients = chainClients
}

func (s *TestSuite) MakeRequest(req *http.Request, expCode int) io.Reader {
	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err, "trying to make request", zap.Any("request", req))

	s.Require().Equal(expCode, resp.StatusCode, "response code did not match")

	return resp.Body
}

func (s *TestSuite) TestChains_Status() {
	s.T().Log("runing test for /status endpoint for each chain")

	for _, chainClient := range s.chainClients {
		url := fmt.Sprintf("%s/status", chainClient.GetRPCAddr())
		req, err := http.NewRequest(http.MethodGet, url, nil)
		s.Require().NoError(err)

		body := s.MakeRequest(req, 200)
		resp := &expb.Status{}
		err = jsonpb.Unmarshal(body, resp)
		s.Assert().NoError(err)

		// assert chain id
		s.Assert().Equal(chainClient.ChainID(), resp.Result.NodeInfo.Network)
	}
}
