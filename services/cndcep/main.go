package cdncep

import (
	"context"
	"fmt"

	"github.com/sandronister/channel-go/core/domain"
)

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
