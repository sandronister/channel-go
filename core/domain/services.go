package domain

import (
	"context"
	"net/http"
)

type CepServices interface {
	Get(ctx context.Context, data chan<- interface{}, cep string)
}

type RequestService interface {
	Get(ctx context.Context, url string) *http.Request
}
