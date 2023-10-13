package main

import (
	"context"
	"time"

	"github.com/sandronister/channel-go/core/domain"
	cdncep "github.com/sandronister/channel-go/services/cndcep"
	"github.com/sandronister/channel-go/services/request"
	"github.com/sandronister/channel-go/services/viacep"
)

var reqService domain.RequestService
var cdnService domain.CepServices
var viaService domain.CepServices

func init() {
	reqService = request.New()
	cdnService = cdncep.New(reqService)
	viaService = viacep.New(reqService)
}

func main() {
	cdnChannel := make(chan string)
	viaChannel := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go cdnService.Get(ctx, cdnChannel, "18050-605")
	go viaService.Get(ctx, viaChannel, "18050-605")

	select {
	case msg := <-cdnChannel:
		println("CDN Channel", msg)
	case msg := <-viaChannel:
		println("Via Cep Channel", msg)
	}

}
