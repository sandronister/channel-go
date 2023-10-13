package cdncep

import (
	"context"
	"fmt"
	"io"
	"net/http"

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
}

func New(req domain.RequestService) domain.CepServices {
	return &cep{request: req}
}

func (c *cep) Get(ctx context.Context, data chan<- string, cep string) {
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

	data <- string(body)
}
