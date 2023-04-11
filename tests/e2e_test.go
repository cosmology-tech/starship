package tests

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var configFile = "./config.yaml"

type TestSuite struct {
	suite.Suite

	config *Config
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
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
	s.T().Log("making request for", zap.Any("request", req))

	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err)

	s.Require().Equal(expCode, resp.StatusCode, "response code did not match")

	resBody, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)

	return resBody
}
