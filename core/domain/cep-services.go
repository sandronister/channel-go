package domain

import (
	"context"
	"net/http"
)

type CDNResult struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type CepServices interface {
	Get(ctx context.Context, data chan<- string, cep string)
}

type RequestService interface {
	Get(ctx context.Context, url string) *http.Request
}
