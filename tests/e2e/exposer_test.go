package e2e

import (
	"fmt"
	"net/http"
	urlpkg "net/url"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	pb "github.com/cosmology-tech/starship/exposer/exposer"
)

func (s *TestSuite) MakeExposerRequest(chain *Chain, req *http.Request, unmarshal proto.Message) {
	host := fmt.Sprintf("http://0.0.0.0:%d%s", chain.Ports.Exposer, req.URL.String())

	url, err := urlpkg.Parse(host)
	s.Require().NoError(err)

	req.URL = url

	body := s.MakeRequest(req, 200)
	err = jsonpb.Unmarshal(body, unmarshal)
	s.Require().NoError(err)
}

func (s *TestSuite) TestExposer_GetNodeID() {
	s.T().Log("running test for /node_id endpoint for exposer")

	chain := s.config.Chains[0]

	req, err := http.NewRequest(http.MethodGet, "/node_id", nil)
	s.Require().NoError(err)

	resp := &pb.ResponseNodeID{}
	s.MakeExposerRequest(chain, req, resp)

	// assert results to expected values
	s.Assert().NotEmpty(resp.NodeId)
}

func (s *TestSuite) TestExposer_GetGenesisFile() {
	s.T().Log("running test for /genesis endpoint for exposer")

	chain := s.config.Chains[0]

	req, err := http.NewRequest(http.MethodGet, "/genesis", nil)
	s.Require().NoError(err)

	// todo: fix unmarshalling of genesis into proto
	resp := &structpb.Struct{}
	s.MakeExposerRequest(chain, req, resp)

	// assert results to expected values
	s.Assert().NotNil(resp)
	s.Assert().Equal(chain.Name, resp.AsMap()["chain_id"])
}

func (s *TestSuite) TestExposer_GetPubKey() {
	s.T().Log("running test for /pub_key endpoint for exposer")

	chain := s.config.Chains[0]

	req, err := http.NewRequest(http.MethodGet, "/pub_key", nil)
	s.Require().NoError(err)

	resp := &pb.ResponsePubKey{}
	s.MakeExposerRequest(chain, req, resp)

	// assert results to expected values
	s.Assert().NotNil(resp)
	s.Assert().NotEmpty(resp.Key)
	s.Assert().Equal("/cosmos.crypto.ed25519.PubKey", resp.Type)
}

func (s *TestSuite) TestExposer_GetPrivKey() {
	s.T().Log("running test for /priv_key endpoint for exposer")
	s.T().Skip("not implemented yet")

	chain := s.config.Chains[0]

	req, err := http.NewRequest(http.MethodGet, "/priv_key", nil)
	s.Require().NoError(err)

	resp := &pb.PrivValidatorKey{}
	s.MakeExposerRequest(chain, req, resp)

	// assert results to expected values
	s.Assert().NotNil(resp)
	s.Assert().NotEmpty(resp.PrivKey)
	s.Assert().NotEmpty(resp.PubKey)
}

func (s *TestSuite) TestExposer_GetKeys() {
	s.T().Log("running test for /keys endpoint for exposer")

	chain := s.config.Chains[0]

	req, err := http.NewRequest(http.MethodGet, "/keys", nil)
	s.Require().NoError(err)

	resp := &pb.Keys{}
	s.MakeExposerRequest(chain, req, resp)

	// assert results to expected values
	s.Assert().NotNil(resp)
	s.Assert().Len(resp.Genesis, 1)
	s.Assert().Len(resp.Validators, 4)
	s.Assert().Len(resp.Keys, 3)
	s.Assert().Len(resp.Relayers, 2)
}
