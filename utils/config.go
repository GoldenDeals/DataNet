package utils

import (
	"github.com/spf13/viper"
)

const configName = "config"
const configType = "yaml"
const mainConfigDir = "."

func Configure() {
	viper.SetConfigName(configName)
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
	viper.SetDefault("log.json", false)

	viper.SetDefault("node.adddr", "/ip4/127.0.0.1/tcp/2000")
	viper.SetDefault("node.initialPeers", []string{}) // TODO: Fill
}
