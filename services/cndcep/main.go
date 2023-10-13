package cdncep

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/sandronister/channel-go/core/domain"
)

type cep struct {
	request domain.RequestService
}

func New(req domain.RequestService) domain.CepServices {
	return &cep{request: req}
}

func (c *cep) Get(ctx context.Context, data chan<- interface{}, cep string) {
	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	req := c.request.Get(ctx, url)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	data <- body
}
