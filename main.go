package main

import (
	"os"
	"os/signal"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/GoldenDeals/DataNet/utils"
)

var lg *zap.SugaredLogger

func main() {
	utils.Configure()
	utils.InitLoggers()

	lg = utils.NewLogger("main")
	lg.Infof("Starting DataNet Node (%s)", viper.GetString("version"))

	lg.Error("Errors")
	lg.Warn("Warn")

	exit := make(chan os.Signal, 2)
	signal.Notify(exit, os.Interrupt)
	<-exit
	ShutdownApp()
}

func ShutdownApp() {

	utils.SyncLoggers()
}
