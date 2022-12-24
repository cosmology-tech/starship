package main

import (
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
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
	ErrValidation       = &ErrResponse{HTTPStatusCode: http.StatusBadRequest, MessageText: "Validation error."}
	ErrNotFound         = &ErrResponse{HTTPStatusCode: http.StatusNotFound, MessageText: "Resource not found."}
	ErrNotImplemented   = &ErrResponse{HTTPStatusCode: http.StatusNotImplemented, MessageText: "Not Implemented."}
	ErrMethodNotAllowed = &ErrResponse{HTTPStatusCode: http.StatusMethodNotAllowed, MessageText: "Method not allowed."}
	ErrRequestBind      = &ErrResponse{HTTPStatusCode: http.StatusBadRequest, MessageText: "Unable to bind request body."}
	ErrInternalServer   = &ErrResponse{HTTPStatusCode: http.StatusInternalServerError, MessageText: "Internal server error."}
)

// NotFound Method to render Json respose, used by middleware
func NotFound(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, ErrNotFound)
}

// MethodNotAllowed Method to render Json respose, used by middleware
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, ErrMethodNotAllowed)
}

func (a *AppServer) renderError(w http.ResponseWriter, r *http.Request, err error, msg ...string) {
	log := a.logger
	errResp := NewErrResponse(err)
	// Logging error at different levels depending on status
	switch code := errResp.HTTPStatusCode; {
	case code < http.StatusInternalServerError:
		log.Warn(err.Error())
	default:
		log.Error(
			"Internal server error",
			zap.Error(err),
		)
	}
	_ = render.Render(w, r, errResp)
}
