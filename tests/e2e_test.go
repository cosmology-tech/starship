package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	pb "github.com/cosmology-tech/starship/exposer/exposer"
)

var configFile = "./config.yaml"

type TestSuite struct {
	suite.Suite

	config *Config
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
}

func (s *TestSuite) MakeRequest(req *http.Request, expCode int) []byte {
	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err, "trying to make request", zap.Any("request", req))

	s.Require().Equal(expCode, resp.StatusCode, "response code did not match")

	resBody, err := io.ReadAll(resp.Body)
	s.Require().NoError(err, "unable to read response body into buffer")

	return resBody
}

func (s *TestSuite) TestChains_Status() {
	s.T().Log("runing test for /status endpoint for each chain")

	for _, chain := range s.config.Chains {
		url := fmt.Sprintf("http://0.0.0.0:%d/status", chain.Ports.Rpc)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		s.Require().NoError(err)

		body := s.MakeRequest(req, 200)
		resp := &pb.Status{}
		err = json.Unmarshal(body, &resp)
		s.Assert().NoError(err)

		// assert chain id
		s.Assert().Equal(chain.Name, resp.Result.NodeInfo.Network)
	}
}
