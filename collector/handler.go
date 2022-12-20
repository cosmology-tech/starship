package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

	byteValue, _ := io.ReadAll(jsonFile)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteValue)
}

func (a *AppServer) GetChains(w http.ResponseWriter, r *http.Request) {
}

func (a *AppServer) GetChainExports(w http.ResponseWriter, r *http.Request) {
}

func (a *AppServer) SetChainExport(w http.ResponseWriter, r *http.Request) {
}

func (a *AppServer) GetChainExport(w http.ResponseWriter, r *http.Request) {
}

func (a *AppServer) GetChainSnapshots(w http.ResponseWriter, r *http.Request) {
}

func (a *AppServer) SetChainSnapshot(w http.ResponseWriter, r *http.Request) {
}

func (a *AppServer) GetChainSnapshot(w http.ResponseWriter, r *http.Request) {
}
