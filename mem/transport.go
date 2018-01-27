package mem

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

var errBadRoute = errors.New("bad route")

func MakeHandler(bs Member, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r := mux.NewRouter()
	LoginHandler := kithttp.NewServer(
		makeLoginEndpoint(bs),
		decodeLoginRequest,
		encodeResponse,
		opts...,
	)
	LogoutHandler := kithttp.NewServer(
		makeLogoutEndpoint(bs),
		decodeLogoutRequest,
		encodeResponse,
		opts...,
	)
	RemarkHandler := kithttp.NewServer(
		makeRemarkEndpoint(bs),
		decodeRemarkRequest,
		encodeResponse,
		opts...,
	)
	r.Handle("/URL/v1/login", LoginHandler).Methods("GET")
	r.Handle("/URL/v1/logout", LogoutHandler).Methods("GET")
	r.Handle("/URL/v1/remark", RemarkHandler).Methods("GET")
	return r
}
func decodeLogoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return logoutRequest{}, nil
}
func decodeRemarkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return remarkRequest{}, nil
}
func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return loginRequest{}, nil
}
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case errBadRoute:
		w.WriteHeader(http.StatusNotFound)
	case errBadRoute:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
