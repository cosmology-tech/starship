package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"io/ioutil"
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
	defer jsonFile.Close()

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
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	return result, nil
}

func (a *AppServer) GetChains(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(a.config.ChainRegistry)
	if err != nil {
		a.renderError(w, r, err)
		return
	}

	var chains []interface{}
	for _, f := range files {
		filename := filepath.Join(a.config.ChainRegistry, f.Name(), "chain.json")
		chainInfo, err := readJSONFile(filename)
		if err != nil {
			a.renderError(w, r, fmt.Errorf("unable to read file %s, err: %d", filename, err))
			return
		}
		chains = append(chains, chainInfo)
	}

	render.JSON(w, r, NewItemsResponse(chains...))
}

func (a *AppServer) GetChain(w http.ResponseWriter, r *http.Request) {
	chain := chi.URLParam(r, "chain")

	filename := filepath.Join(a.config.ChainRegistry, chain, "chain.json")

	chainInfo, err := readJSONFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		render.Render(w, r, ErrNotFound)
		return
	} else if err != nil {
		a.renderError(w, r, fmt.Errorf("unable to read file %s, err: %d", filename, err))
		return
	}

	render.JSON(w, r, chainInfo)
}
