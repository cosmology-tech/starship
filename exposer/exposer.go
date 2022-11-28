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
	genesisFileKey  = "GENESIS_FILE"
	portKey         = "GENESIS_PORT"
	mnemonicFileKey = "KEYS_FILE"
)

var statusURL = "http://0.0.0.0:26657/status"

type aStatusResponse struct {
	Result Result `json:"result"`
}

type aResult struct {
	NodeInfo      NodeInfo      `json:"node_info"`
	ValidatorInfo ValidatorInfo `json:"validator_info"`
}

type aNodeInfo struct {
	ID      string `json:"id"`
	Network string `json:"network"`
}

type aValidatorInfo struct {
	Address string `json:"address"`
	PubKey  PubKey `json:"pub_key"`
}

type aPubKey struct {
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

func NodeIDHandler(w http.ResponseWriter, r *http.Request) {
	status := getNodeStatus()

	_, _ = io.WriteString(w, status.Result.NodeInfo.ID)
}

func PubKeyHandler(w http.ResponseWriter, r *http.Request) {
	status := getNodeStatus()

	response := map[string]string{
		"@type": "/cosmos.crypto.ed25519.PubKey",
		"key":   status.Result.ValidatorInfo.PubKey.Value,
	}

	data, _ := json.Marshal(response)
	_, _ = io.WriteString(w, string(data))
}

func handleJSONFile(w http.ResponseWriter, r *http.Request, file string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error opening file at %s, err: %s", file, err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(byteValue)
}

func GenesisHandler(w http.ResponseWriter, r *http.Request) {
	genesisFile := os.Getenv(genesisFileKey)
	handleJSONFile(w, r, genesisFile)
}

func KeysHandler(w http.ResponseWriter, r *http.Request) {
	mnemonicFile := os.Getenv(mnemonicFileKey)
	handleJSONFile(w, r, mnemonicFile)
}

func Routes() {
	http.HandleFunc("/node_id", NodeIDHandler)
	http.HandleFunc("/pub_key", PubKeyHandler)
	http.HandleFunc("/genesis", GenesisHandler)
	http.HandleFunc("/keys", KeysHandler)
}

func main() {
	fmt.Println("Server started ...")

	Routes()

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(portKey)), nil)
	if err != nil {
		log.Fatalf("Fail to start server, %s\n", err)
	}
}
