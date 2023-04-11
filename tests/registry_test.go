package tests

import (
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"

	pb "github.com/cosmology-tech/starship/registry/registry"
)

func (s *TestSuite) MakeRegistryRequest(req *http.Request, unmarshal interface{}) {
	req.Host = fmt.Sprintf("0.0.0.0:%d", s.config.Registry.Ports.Rest)

	body := s.MakeRequest(req, 200)
	err := json.Unmarshal(body, unmarshal)
	s.Require().NoError(err)
}

func (s *TestSuite) TestRegistryChainIds() {
	s.T().Log("runing test for /chain_ids endpoint for registry")

	req, err := http.NewRequest(http.MethodGet, "/chain_ids", nil)
	s.Require().NoError(err)

	respChains := &pb.ResponseChainIDs{}
	s.MakeRegistryRequest(req, respChains)

	// assert results to expected values
	s.Assert().Len(respChains.ChainIds, len(s.config.Chains))
}

func (s *TestSuite) TestRegistryChains() {
	s.T().Log("runing test for /chains endpoint for registry")

	req, err := http.NewRequest(http.MethodGet, "/chains", nil)
	s.Require().NoError(err)

	respChains := &pb.ResponseChains{}
	s.MakeRegistryRequest(req, respChains)

	// assert results to expected values
	s.Assert().Len(respChains.Chains, len(s.config.Chains))
	for i := range respChains.Chains {
		chain := respChains.Chains[i]
		expChain := s.config.GetChain(chain.ChainId)
		s.Assert().NotNil(expChain)

		s.Assert().Equal(expChain.Type, chain.ChainName)
	}
}
