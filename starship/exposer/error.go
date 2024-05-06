package main

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	MessageText string `json:"message"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// NewErrResponse create http aware errors from custom errors
func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		MessageText:    err.Error(),
	}
}

var (
	ErrInternalServer = &ErrResponse{HTTPStatusCode: http.StatusInternalServerError, MessageText: "Internal server error."}
)
