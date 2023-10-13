package httpservice

import (
	"io"
	"net/http"

	"github.com/sandronister/channel-go/core/domain"
)

type service struct {
}

func New() domain.HttpService {
	return &service{}
}

func (s *service) Do(req *http.Request) []byte {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}
