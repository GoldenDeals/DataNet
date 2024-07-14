package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/GoldenDeals/DataNet/network"
	"github.com/GoldenDeals/DataNet/utils"
	"github.com/rs/zerolog/log"
)

var lg utils.LoggerGroup

func main() {
	utils.Configure()

	utils.InitLogger()
	lg = utils.NewLoggerGroup("main")

	lg(log.Info()).Msg("Starting")
	node, err := (&network.Node{}).Init()
	if err != nil {
		lg(log.Fatal()).Err(err).Msg("Error Starting node")
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	lg(log.Warn()).Msg("Got signal. Shutting down")
	node.Shutdown()
}
