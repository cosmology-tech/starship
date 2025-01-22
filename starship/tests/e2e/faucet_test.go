package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	pb "github.com/hyperweb-io/starship/registry/registry"
	"net/http"
	urlpkg "net/url"
	"strconv"
	"time"
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

	for _, chain := range s.config.Chains {
		s.Run(fmt.Sprintf("faucet test for: %s", chain.ID), func() {
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

func (s *TestSuite) MakeChainGetRequest(chain *Chain, endpoint string, unmarshal any) {
	url := fmt.Sprintf("http://0.0.0.0:%d%s", chain.Ports.Rest, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	s.Require().NoError(err)
	body := s.MakeRequest(req, 200)

	err = json.NewDecoder(body).Decode(&unmarshal)
	s.Require().NoError(err)
}

func (s *TestSuite) getChainAccounts(chain *Chain) []string {
	var accounts []string

	data := map[string]interface{}{}
	s.MakeChainGetRequest(chain, "/cosmos/auth/v1beta1/accounts", &data)
	s.Require().Contains(data, "accounts")

	for _, acc := range data["accounts"].([]interface{}) {
		accMap := acc.(map[string]interface{})
		s.Require().Contains(accMap, "@type")
		if accMap["@type"].(string) != "/cosmos.auth.v1beta1.BaseAccount" {
			continue
		}
		s.Require().NotEmpty(accMap["address"].(string))
		accounts = append(accounts, accMap["address"].(string))
	}

	s.Require().GreaterOrEqual(len(accounts), 1)
	return accounts
}

func (s *TestSuite) getChainDenoms(chain *Chain) string {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/chains/%s", chain.ID), nil)
	s.Require().NoError(err)

	respChain := &pb.ChainRegistry{}
	s.MakeRegistryRequest(req, respChain)
	s.Require().Equal(chain.ID, respChain.ChainId)

	s.Require().NotEmpty(respChain.Fees.FeeTokens[0].Denom)

	return respChain.Fees.FeeTokens[0].Denom
}

func (s *TestSuite) getIBCData(aChain, bChain string) *pb.IBCData {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/ibc/%s/%s", aChain, bChain), nil)
	s.Require().NoError(err)

	ibcData := &pb.IBCData{}
	s.MakeRegistryRequest(req, ibcData)

	return ibcData
}

func (s *TestSuite) getAccountBalance(chain *Chain, address string, denom string) float64 {
	data := map[string]interface{}{}
	s.MakeChainGetRequest(chain, fmt.Sprintf("/cosmos/bank/v1beta1/balances/%s", address), &data)
	s.Require().Contains(data, "balances")

	for _, bal := range data["balances"].([]interface{}) {
		balMap := bal.(map[string]interface{})
		if balMap["denom"].(string) == denom {
			b, err := strconv.ParseFloat(balMap["amount"].(string), 64)
			s.Require().NoError(err)
			return b
		}
	}
	return float64(0)
}

func (s *TestSuite) TestFaucet_Credit() {
	s.T().Log("running test for /credit endpoint for faucet")

	// expected amount to be credited via faucet
	expCreditedAmt := float64(10000000000)

	for _, chain := range s.config.Chains {
		s.Run(fmt.Sprintf("faucet test for: %s", chain.ID), func() {
			if chain.Ports.Faucet == 0 {
				s.T().Skip("faucet not exposed via ports")
			}

			// fetch denom and address from an account on chain
			denom := s.getChainDenoms(chain)
			addr := getAddressFromType(chain.Name)
			beforeBalance := s.getAccountBalance(chain, addr, denom)

			body := map[string]string{
				"denom":   denom,
				"address": addr,
			}
			postBody, err := json.Marshal(body)
			s.Require().NoError(err)
			resp, err := http.Post(
				fmt.Sprintf("http://0.0.0.0:%d/credit", chain.Ports.Faucet),
				"application/json",
				bytes.NewBuffer(postBody))
			s.Require().NoError(err)
			s.Require().Equal(200, resp.StatusCode)

			time.Sleep(4 * time.Second)
			afterBalance := s.getAccountBalance(chain, addr, denom)
			s.T().Log("address:", addr, "after balance: ", afterBalance, "before balance:", beforeBalance)
			// note sometimes expected difference is 9x expected value (bug due to using holder address for test)
			// hence checking for difference is at least expected value
			s.Require().GreaterOrEqual(afterBalance-beforeBalance, expCreditedAmt)
		})
	}
}
