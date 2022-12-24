package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

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

func readJSONFile(file string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	return result, nil
}

func (a *AppServer) GetChains(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(a.config.ChainRegistry)
	if err != nil {
		a.renderError(w, r, err)
		return
	}

	var chains []interface{}
	for _, f := range files {
		filename := filepath.Join(a.config.ChainRegistry, f.Name(), "chain.json")
		info, err := readJSONFile(filename)
		if err != nil {
			a.renderError(w, r, fmt.Errorf("unable to read file %s, err: %d", filename, err))
			return
		}
		chains = append(chains, info)
	}

	render.JSON(w, r, NewItemsResponse(chains))
}

func (a *AppServer) GetChainIDs(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(a.config.ChainRegistry)
	if err != nil {
		a.renderError(w, r, err)
		return
	}

	var chainIDs []string
	for _, f := range files {
		filename := filepath.Join(a.config.ChainRegistry, f.Name(), "chain.json")
		info, err := readJSONFile(filename)
		if err != nil {
			a.renderError(w, r, fmt.Errorf("unable to read file %s, err: %d", filename, err))
			return
		}
		chainID, ok := info["chain_id"].(string)
		if !ok {
			a.renderError(w, r, fmt.Errorf("unable to get chain id for %s", filename))
			return
		}
		chainIDs = append(chainIDs, chainID)
	}

	render.JSON(w, r, NewItemsResponse(chainIDs))
}

// GetChain handles the incoming request for a single chain given the chain id
// Note, we use chain-id instead of chain type, since it is expected, that there
// can be multiple chains of same type by unique chain ids
func (a *AppServer) GetChain(w http.ResponseWriter, r *http.Request) {
	chainID := chi.URLParam(r, "chain")

	filename := filepath.Join(a.config.ChainRegistry, chainID, "chain.json")

	info, err := readJSONFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		render.Render(w, r, ErrNotFound)
		return
	} else if err != nil {
		a.renderError(w, r, fmt.Errorf("unable to read file %s, err: %d", filename, err))
		return
	}

	render.JSON(w, r, info)
}

func (a *AppServer) GetChainAssets(w http.ResponseWriter, r *http.Request) {
	chainID := chi.URLParam(r, "chain")

	filename := filepath.Join(a.config.ChainRegistry, chainID, "assetlist.json")

	info, err := readJSONFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		render.Render(w, r, ErrNotFound)
		return
	} else if err != nil {
		a.renderError(w, r, fmt.Errorf("unable to read file %s, err: %d", filename, err))
		return
	}

	render.JSON(w, r, info)
}

func (a *AppServer) GetAllIBC(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}

func (a *AppServer) GetIBCChainsData(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}

func (a *AppServer) SetIBCChainsData(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}

func (a *AppServer) GetIBCChainsChannels(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}

func (a *AppServer) AddIBCChainChannel(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}
