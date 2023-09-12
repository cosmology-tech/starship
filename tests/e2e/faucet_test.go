package e2e

import (
	"encoding/json"
	"fmt"
	"net/http"
	urlpkg "net/url"
)

func (s *TestSuite) MakeFaucetRequest(chain *Chain, req *http.Request, unmarshal map[string]interface{}) {
	host := fmt.Sprintf("http://0.0.0.0:%d%s", chain.Ports.Faucet, req.URL.String())

	url, err := urlpkg.Parse(host)
	s.Require().NoError(err)

	req.URL = url

	body := s.MakeRequest(req, 200)

	err = json.NewDecoder(body).Decode(&unmarshal)
	s.Require().NoError(err)
}

func (s *TestSuite) TestFaucet_Status() {
	s.T().Log("running test for /status endpoint for faucet")

	if s.config.Faucet != nil && !s.config.Faucet.Enabled {
		s.T().Skip("faucet disabled")
	}

	for _, chain := range s.config.Chains {
		s.Run(fmt.Sprintf("facuet test for: %s", chain.Name), func() {
			if chain.Faucet != nil && !chain.Faucet.Enabled {
				s.T().Skip("faucet disabled for chain")
			}
			if chain.Ports.Faucet == 0 {
				s.T().Skip("faucet not exposed via ports")
			}

			req, err := http.NewRequest(http.MethodGet, "/status", nil)
			s.Require().NoError(err)

			resp := map[string]interface{}{}
			s.MakeFaucetRequest(chain, req, resp)

			s.Require().Equal("ok", resp["status"])
		})
	}
}
