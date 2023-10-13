package domain

import (
	"context"
	"net/http"
)

type CepServices interface {
	Get(ctx context.Context, data chan<- string, cep string)
}

type RequestService interface {
	Get(ctx context.Context, url string) *http.Request
}

type HttpService interface {
	Do(req *http.Request) []byte
}
