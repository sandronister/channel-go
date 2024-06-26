package main

import (
	"context"
	"time"

	"github.com/sandronister/channel-go/configs"
	"github.com/sandronister/channel-go/internals/di"
)

func main() {

	cdnChannel := make(chan string)
	viaChannel := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	cdnService := di.NewCDNCep(ctx, config.CDNCepPath)
	viaService := di.NewVIACep(ctx, config.VIACepPath)

	go cdnService.Get(cdnChannel, "18050-605")
	go viaService.Get(viaChannel, "18050-605")

	select {
	case msg := <-cdnChannel:
		println("CDN Channel", msg)
	case msg := <-viaChannel:
		println("Via Cep Channel", msg)
	case <-time.After(time.Second):
		println("Timeout")
	}

}
