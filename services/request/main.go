package request

import (
	"context"
	"net/http"

	"github.com/sandronister/channel-go/core/domain"
)

type service struct {
}

func New() domain.RequestService {
	return &service{}
}

func (s *service) Get(ctx context.Context, url string) *http.Request {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}
	return req
}
