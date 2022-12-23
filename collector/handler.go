package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"os"
	"strings"

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

func (a *AppServer) GetChains(w http.ResponseWriter, r *http.Request) {
	chains, err := a.db.GetChains()
	if err != nil {
		a.renderError(w, r, err, "unable to get chains")
		return
	}

	render.JSON(w, r, NewItemsResponse(chains))
}

func (a *AppServer) GetChainExports(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}

func (a *AppServer) SetChainExport(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}

func (a *AppServer) GetChainExport(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotImplemented)
}

func (a *AppServer) GetChainSnapshots(w http.ResponseWriter, r *http.Request) {
	chainID := chi.URLParam(r, "chain")
	valID := chi.URLParam(r, "validator")

	snapshotsIDs, err := a.db.ListSnapshots(chainID, valID)
	if err != nil {
		a.renderError(w, r, err)
		return
	}

	render.JSON(w, r, NewItemsResponse(snapshotsIDs))
}

func (a *AppServer) SetChainSnapshot(w http.ResponseWriter, r *http.Request) {
	chainID := chi.URLParam(r, "chain")
	valID := chi.URLParam(r, "validator")
	snapshotID := chi.URLParam(r, "id")

	if !strings.HasSuffix(snapshotID, a.config.SnapshotExt) {
		a.renderError(w, r, fmt.Errorf("snapshot id must contain suffix :%s", a.config.SnapshotExt))
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		a.logger.Warn(ErrRequestBind.MessageText, zap.Error(err))
		render.Render(w, r, ErrRequestBind)
		return
	}

	err = a.db.StoreSnapshot(chainID, valID, snapshotID, data)
	if err != nil {
		a.logger.Warn("unable to store snapshot",
			zap.String("snapshotID", snapshotID),
			zap.Error(err))
		a.renderError(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated) // Set response status to 201
}

func (a *AppServer) GetChainSnapshot(w http.ResponseWriter, r *http.Request) {
	chainID := chi.URLParam(r, "chain")
	valID := chi.URLParam(r, "validator")
	snapshotID := chi.URLParam(r, "id")

	data, err := a.db.GetSnapshot(chainID, valID, snapshotID)
	if err != nil {
		a.logger.Warn("unable to read snapshot",
			zap.String("snapshotID", snapshotID),
			zap.Error(err))
		a.renderError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/gzip")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
