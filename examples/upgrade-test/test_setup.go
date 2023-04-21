package upgrade_test

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