package http

import (
	"io"
	"net/http"

	"github.com/sandronister/channel-go/internals/pkg/interfaces"
)

type Service struct {
	request interfaces.RequestService
}

func New(request interfaces.RequestService) *Service {
	return &Service{request: request}
}

func (s *Service) Do(url string) ([]byte, error) {
	req, err := s.request.Get(url)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
