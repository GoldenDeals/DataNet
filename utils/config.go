package utils

import (
	"github.com/spf13/viper"
)

/*
log.encoding
log.level
log.out
log.erros
*/

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
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.encoding", "json")
	viper.SetDefault("log.out", []string{"stdout"})
	viper.SetDefault("log.errors", []string{"stdout", "stderr"})
	viper.SetDefault("version", "undefined")
}
