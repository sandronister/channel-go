package usecase

import (
	"fmt"
	"log"

	"github.com/sandronister/channel-go/internals/pkg/interfaces"
)

type CDNCep struct {
	path string
	http interfaces.HttpService
}

func NewCDNCep(path string, http interfaces.HttpService) *CDNCep {
	return &CDNCep{path: path, http: http}
}

func (c *CDNCep) Get(data chan<- string, cep string) {
	url := fmt.Sprintf(c.path, cep)
	res, err := c.http.Do(url)
	if err != nil {
		fmt.Print("DEU Merda")
		log.Fatal(err)
	}
	fmt.Println("CDN CEP Response", string(res))
	data <- string(res)
}
