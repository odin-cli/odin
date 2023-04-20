package utils

import (
	"fmt"
	"github.com/odin-cli/odin/pkg/models"
	"github.com/spf13/viper"
	"os"
)

func ReadConfig() *models.Config {
	var config *models.Config
	home, _ := os.UserHomeDir()
	viper.SetConfigFile(fmt.Sprintf("%v/.odin.config.yaml", home))
	if err := viper.ReadInConfig(); err != nil {
		PrintError("Error to load config! Runner: odin config")
		os.Exit(1)
	}
	viper.Unmarshal(&config)
	return config
}
