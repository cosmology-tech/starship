package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	genesisFileKey = "GENESIS_FILE"
	portKey        = "GENESIS_PORT"
)

var statusURL = "http://0.0.0.0:26657/status"

type StatusResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	NodeInfo      NodeInfo      `json:"node_info"`
	ValidatorInfo ValidatorInfo `json:"validator_info"`
}

type NodeInfo struct {
	ID      string `json:"id"`
	Network string `json:"network"`
}

type ValidatorInfo struct {
	Address string `json:"address"`
	PubKey  PubKey `json:"pub_key"`
}

type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func getNodeStatus() StatusResponse {
	resp, err := http.Get(statusURL)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		log.Fatalln(err)
	}

	return statusResp
}

func getNodeIDHandler(w http.ResponseWriter, r *http.Request) {
	status := getNodeStatus()

	_, _ = io.WriteString(w, status.Result.NodeInfo.ID)
}

func getPubKeyHandler(w http.ResponseWriter, r *http.Request) {
	status := getNodeStatus()

	response := map[string]string{
		"@type": "/cosmos.crypto.ed25519.PubKey",
		"key":   status.Result.ValidatorInfo.PubKey.Value,
	}

	data, _ := json.Marshal(response)
	_, _ = io.WriteString(w, string(data))
}

func getGenesisHandler(w http.ResponseWriter, r *http.Request) {
	genesisFile := os.Getenv(genesisFileKey)
	jsonFile, err := os.Open(genesisFile)
	if err != nil {
		log.Fatalf("Error opening genesis file at %s", genesisFile)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(byteValue)
}

func main() {
	fmt.Println("Server started ...")
	http.HandleFunc("/node_id", getNodeIDHandler)
	http.HandleFunc("/pub_key", getPubKeyHandler)
	http.HandleFunc("/genesis", getGenesisHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(portKey)), nil)
	if err != nil {
		log.Fatalf("Fail to start server, %s\n", err)
	}
}
