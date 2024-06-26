package di

import (
	"context"

	"github.com/sandronister/channel-go/internals/pkg/http"
	"github.com/sandronister/channel-go/internals/pkg/interfaces"
	"github.com/sandronister/channel-go/internals/pkg/request"
	"github.com/sandronister/channel-go/internals/usecase"
)

func NewCDNCep(ctx context.Context, path string) interfaces.CepServices {
	request := request.New(ctx)
	httpService := http.New(request)
	return usecase.NewCDNCep(path, httpService)
}
