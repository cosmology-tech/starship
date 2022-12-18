package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

func fetchNodeStatus(url string) (StatusResponse, error) {
	var statusResp StatusResponse

	resp, err := http.Get(url)
	if err != nil {
		return statusResp, fmt.Errorf("unable to fetch status, err: %d", err)
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return statusResp, fmt.Errorf("unable to parse status response, err: %d", err)
	}

	return statusResp, nil
}

func (a *AppServer) renderJSONFile(w http.ResponseWriter, r *http.Request, filePath string) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		a.logger.Error("Error opening file",
			zap.String("file", filePath),
			zap.Error(err))
		a.renderError(w, r, fmt.Errorf("error opening json file: %s", filePath))
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteValue)
}

func (a *AppServer) GetNodeID(w http.ResponseWriter, r *http.Request) {
	status, err := fetchNodeStatus(a.config.StatusURL)
	if err != nil {
		a.renderError(w, r, err)
		return
	}

	render.PlainText(w, r, status.Result.NodeInfo.ID)
}

func (a *AppServer) GetPubKey(w http.ResponseWriter, r *http.Request) {
	status, err := fetchNodeStatus(a.config.StatusURL)
	if err != nil {
		a.renderError(w, r, err)
		return
	}

	response := map[string]string{
		"@type": "/cosmos.crypto.ed25519.PubKey",
		"key":   status.Result.ValidatorInfo.PubKey.Value,
	}

	render.JSON(w, r, response)
}

func (a *AppServer) GetGenesisFile(w http.ResponseWriter, r *http.Request) {
	a.renderJSONFile(w, r, a.config.GenesisFile)
}

func (a *AppServer) GetKeysFile(w http.ResponseWriter, r *http.Request) {
	a.renderJSONFile(w, r, a.config.MnemonicFile)
}

func (a *AppServer) GetPrivKeysFile(w http.ResponseWriter, r *http.Request) {
	a.renderJSONFile(w, r, a.config.PrivValFile)
}

// SetPrivKeysFile sets a priv keys json file at the place
func (a *AppServer) SetPrivKeysFile(w http.ResponseWriter, r *http.Request) {
	var privKey PrivValKey
	if err := render.DecodeJSON(r.Body, &privKey); err != nil {
		a.logger.Warn(ErrRequestBind.MessageText)
		render.Render(w, r, ErrRequestBind)
		return
	}

	// Write request body as priv_validator_json key
	file, err := json.MarshalIndent(privKey, "", " ")
	if err != nil {
		a.renderError(w, r, fmt.Errorf("unable to marshal struct to file, err: %s", err))
		return
	}
	err = os.WriteFile(a.config.PrivValFile, file, 0644)
	if err != nil {
		a.renderError(w, r, fmt.Errorf("unable to write file, err: %s", err))
		return
	}

	render.NoContent(w, r)
}
