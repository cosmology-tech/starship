package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	pb "github.com/cosmology-tech/starship/exposer/exposer"
)

var configEnvKey = "TEST_CONFIG_FILE"

type TestSuite struct {
	suite.Suite

	configFile string
	config     *Config
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	s.T().Log("setting up e2e integration test suite...")

	// read config file from yaml
	configFile := os.Getenv(configEnvKey)
	if configFile == "" {
		s.T().Log(fmt.Errorf("env var %s not set, using default value", configEnvKey))
		configFile = "configs/two-chain.yaml"
	}
	configFile = strings.Replace(configFile, "tests/e2e/", "", -1)
	yamlFile, err := os.ReadFile(configFile)
	s.Require().NoError(err)
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	s.Require().NoError(err)

	s.config = config
	s.configFile = configFile
}

func (s *TestSuite) MakeRequest(req *http.Request, expCode int) io.Reader {
	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err, "trying to make request", zap.Any("request", req))

	s.Require().Equal(expCode, resp.StatusCode, "response code did not match")

	return resp.Body
}

func (s *TestSuite) TestChains_Status() {
	s.T().Log("running test for /status endpoint for each chain")

	for _, chain := range s.config.Chains {
		url := fmt.Sprintf("http://0.0.0.0:%d/status", chain.Ports.Rpc)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		s.Require().NoError(err)

		body := s.MakeRequest(req, 200)
		resp := &pb.Status{}
		err = jsonpb.Unmarshal(body, resp)
		s.Assert().NoError(err)

		// assert chain id
		s.Assert().Equal(chain.Name, resp.Result.NodeInfo.Network)
	}
}

func (s *TestSuite) TestChains_StakingParams() {
	s.T().Log("running test for /staking/parameters endpoint for each chain")

	expUnbondingTime := "90s" // default value
	switch s.configFile {
	case "configs/one-chain.yaml":
		expUnbondingTime = "5s" // based on genesis override in one-chain.yaml file
	case "configs/one-chain-custom-scripts.yaml":
		expUnbondingTime = "15s" // based on genesis override in the custom script
	}

	url := fmt.Sprintf("http://0.0.0.0:%d/cosmos/staking/v1beta1/params", s.config.Chains[0].Ports.Rest)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	s.Require().NoError(err)
	body := s.MakeRequest(req, 200)
	data := map[string]interface{}{}

	err = json.NewDecoder(body).Decode(&data)
	s.Require().NoError(err)

	s.Require().Equal(expUnbondingTime, data["params"].(map[string]interface{})["unbonding_time"])
}
