package mem

import (
	"context"
	_ "fmt"

	"github.com/go-kit/kit/endpoint"
)

var ()

type loginRequest struct{}
type loginResponse struct {
	Err error `json:"error,omitempty"`
}

func (r loginResponse) error() error { return r.Err }
func makeLoginEndpoint(s Member) endpoint.Endpoint {
	var ch chan int = make(chan int)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return loginResponse{Err: nil}, nil
	}
}

type logoutRequest struct{}
type logoutResponse struct {
	Err error `json:"error,omitempty"`
}

func (r logoutResponse) error() error { return r.Err }
func makeLogoutEndpoint(s Member) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return logoutResponse{Err: nil}, nil
	}
}

type remarkRequest struct{}
type remarkResponse struct {
	Err error `json:"error,omitempty"`
}

func (r remarkResponse) error() error { return r.Err }
func makeRemarkEndpoint(s Member) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return remarkResponse{Err: nil}, nil
	}
}
