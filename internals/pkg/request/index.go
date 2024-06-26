package request

import (
	"context"
	"net/http"
)

type Service struct {
	ctx context.Context
}

func New(ctx context.Context) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) Get(url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(s.ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
