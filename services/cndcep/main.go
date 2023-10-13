package cdncep

import (
	"context"
	"fmt"

	"github.com/sandronister/channel-go/core/domain"
)

type Result struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type cep struct {
	request domain.RequestService
	http    domain.HttpService
}

func New(req domain.RequestService, http domain.HttpService) domain.CepServices {
	return &cep{request: req, http: http}
}

func (c *cep) Get(ctx context.Context, data chan<- string, cep string) {
	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	req := c.request.Get(ctx, url)
	body := c.http.Do(req)
	data <- string(body)
}
