package usecase

import (
	"fmt"
	"log"

	"github.com/sandronister/channel-go/internals/pkg/interfaces"
)

type VIACep struct {
	path string
	http interfaces.HttpService
}

func NewVIACep(path string, http interfaces.HttpService) *VIACep {
	return &VIACep{path: path, http: http}
}

func (c *VIACep) Get(data chan<- string, cep string) {
	url := fmt.Sprintf(c.path, cep)
	res, err := c.http.Do(url)
	if err != nil {
		log.Fatal(err)
	}

	data <- string(res)
}
