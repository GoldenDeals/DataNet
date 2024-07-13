package main

import (
	"os"
	"os/signal"

	"github.com/spf13/viper"

	"github.com/GoldenDeals/DataNet/network"
	"github.com/GoldenDeals/DataNet/utils"
)

func main() {
	utils.Configure()
	lf := utils.InitLoggers()

	lg := lf.NewLogger("main")
	lg.Infof("Starting DataNet Node (%s)", viper.GetString("version"))

	_, err := network.SetupNode(lf.NewLogger("net"))
	if err != nil {
		lg.Fatalw("Error while starting node", "err", err)
	}

	//nolint:mnd  // Т.к. не надо это число выностить в константу
	exit := make(chan os.Signal, 2)
	signal.Notify(exit, os.Interrupt)
	<-exit

	lg.Warn("Got sinal. Shutting down")
	lf.SyncLoggers()
}
