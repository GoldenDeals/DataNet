package utils

import (
	"flag"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

const RUNTIME_EXTRA_WORKERS = 3

const configName = "config"
const configType = "yaml"
const mainConfigDir = "."

func Configure() {
	confname := flag.String("config", "", "Config name")
	flag.Parse()
	if *confname == "" {
		viper.SetConfigName(configName)
	} else {
		viper.SetConfigName(*confname)
	}

	viper.SetConfigType(configType)

	viper.AddConfigPath(mainConfigDir)

	configDefaults()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func configDefaults() {
	viper.SetDefault("version", "undefined")

	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.encoding", "json")
	viper.SetDefault("log.out", []string{"stdout"})
	viper.SetDefault("log.errors", []string{"stdout", "stderr"})

	viper.SetDefault("node.addr", "127.0.0.1:8080")
	viper.SetDefault("node.listen", "0.0.0.0")
	viper.SetDefault("node.port", 8080)

	viper.SetDefault("node.timeouts.ping", 15*time.Second)
	viper.SetDefault("node.timeouts.idle", 2*time.Minute)

	viper.SetDefault("node.max.dialAttempts", 5)
	viper.SetDefault("node.max.inConnections", 256)
	viper.SetDefault("node.max.outConnections", 256)
	viper.SetDefault("node.max.messageSize", 0)
	viper.SetDefault("node.max.workers", runtime.NumCPU()-RUNTIME_EXTRA_WORKERS)
	viper.SetDefault("node.privateKeySeed", "gen")

	viper.SetDefault("node.initialPeers", []string{}) // TODO: Fill
}
