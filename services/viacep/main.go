package viacep

import (
	"context"
	"fmt"

	"github.com/sandronister/channel-go/core/domain"
)

type cep struct {
	request     domain.RequestService
	httpService domain.HttpService
}

func New(req domain.RequestService, httpservice domain.HttpService) domain.CepServices {
	return &cep{request: req, httpService: httpservice}
}

func (c *cep) Get(ctx context.Context, data chan<- string, cep string) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req := c.request.Get(ctx, url)
	body := c.httpService.Do(req)
	data <- string(body)
}
