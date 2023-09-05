package e2e

import (
	"fmt"
	"net/http"
	urlpkg "net/url"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	pb "github.com/cosmology-tech/starship/registry/registry"
)

func (s *TestSuite) MakeRegistryRequest(req *http.Request, unmarshal proto.Message) {
	host := fmt.Sprintf("http://0.0.0.0:%d%s", s.config.Registry.Ports.Rest, req.URL.String())

	url, err := urlpkg.Parse(host)
	s.Require().NoError(err)

	req.URL = url

	body := s.MakeRequest(req, 200)
	err = jsonpb.Unmarshal(body, unmarshal)
	//err = json.Unmarshal(body, unmarshal)
	s.Require().NoError(err)
}

func (s *TestSuite) TestRegistry_ListChainIds() {
	s.T().Log("running test for /chain_ids endpoint for registry")

	req, err := http.NewRequest(http.MethodGet, "/chain_ids", nil)
	s.Require().NoError(err)

	respChains := &pb.ResponseChainIDs{}
	s.MakeRegistryRequest(req, respChains)

	// assert results to expected values
	s.Assert().Len(respChains.ChainIds, len(s.config.Chains))
}

func (s *TestSuite) TestRegistry_ListChains() {
	s.T().Log("running test for /chains endpoint for registry")

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

func (s *TestSuite) TestRegistry_GetChain() {
	s.T().Log("running test for /chains/{chain} endpoint for registry")

	for _, chain := range s.config.Chains {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/chains/%s", chain.Name), nil)
		s.Require().NoError(err)

		respChain := &pb.ChainRegistry{}
		s.MakeRegistryRequest(req, respChain)

		s.Assert().Equal(chain.Name, respChain.ChainId)
		s.Assert().Equal(chain.Type, respChain.ChainName)
		if chain.Ports.Rpc != 0 {
			s.Assert().Equal(fmt.Sprintf("http://localhost:%d", chain.Ports.Rpc), respChain.Apis.Rpc[0].Address)
		}
		if chain.Ports.Grpc != 0 {
			s.Assert().Equal(fmt.Sprintf("http://localhost:%d", chain.Ports.Grpc), respChain.Apis.Grpc[0].Address)
		}
		if chain.Ports.Rest != 0 {
			s.Assert().Equal(fmt.Sprintf("http://localhost:%d", chain.Ports.Rest), respChain.Apis.Rest[0].Address)
		}
	}
}

func (s *TestSuite) TestRegistry_ListChainPeers() {
	s.T().Log("running test for /chains/{chain}/peers endpoint for registry")

	for _, chain := range s.config.Chains {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/chains/%s/peers", chain.Name), nil)
		s.Require().NoError(err)

		respPeers := &pb.Peers{}
		s.MakeRegistryRequest(req, respPeers)

		s.Assert().Len(respPeers.Seeds, 1)
		s.Assert().Len(respPeers.PersistentPeers, chain.NumValidators-1)
	}
}

func (s *TestSuite) TestRegistry_ListChainApis() {
	s.T().Log("running test for /chains/{chain}/apis endpoint for registry")

	for _, chain := range s.config.Chains {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/chains/%s/apis", chain.Name), nil)
		s.Require().NoError(err)

		respAPIs := &pb.APIs{}
		s.MakeRegistryRequest(req, respAPIs)

		s.Assert().Len(respAPIs.Rpc, 1)
	}
}

func (s *TestSuite) TestRegistry_GetChainAssets() {
	s.T().Log("running test for /chains/{chain}/assets endpoint for registry")

	for _, chain := range s.config.Chains {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/chains/%s/assets", chain.Name), nil)
		s.Require().NoError(err)

		respAssets := &pb.ResponseChainAssets{}
		s.MakeRegistryRequest(req, respAssets)

		s.Assert().Equal(chain.Type, respAssets.ChainName)
		s.Assert().Len(respAssets.Assets, 1)
	}
}

func (s *TestSuite) TestRegistry_ListIBC() {
	s.T().Log("running test for /ibc endpoint for registry")
	if len(s.config.Relayers) == 0 {
		s.T().Skip("skip running ibc test for chains without relayers")
	}

	req, err := http.NewRequest(http.MethodGet, "/ibc", nil)
	s.Require().NoError(err)

	respIBC := &pb.ResponseListIBC{}
	s.MakeRegistryRequest(req, respIBC)

	// assert results to expected values
	s.Assert().Len(respIBC.Data, len(s.config.Relayers)*2, "number of ibc information should be double the number of relayers")
}

func (s *TestSuite) TestRegistry_GetChainKeys() {
	s.T().Log("running test for /chains/{chain}/keys endpoint for registry")

	for _, chain := range s.config.Chains {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/chains/%s/keys", chain.Name), nil)
		s.Require().NoError(err)

		respKeys := &pb.Keys{}
		s.MakeRegistryRequest(req, respKeys)

		// assert results to expected values
		s.Assert().Len(respKeys.Genesis, 1)
		s.Assert().Len(respKeys.Validators, 4)
		s.Assert().Len(respKeys.Keys, 3)
		s.Assert().Len(respKeys.Relayers, 2)
	}
}
