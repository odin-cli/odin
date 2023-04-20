package config

import (
	"github.com/lbernardo/fx-webserver/pkg/model/config"
	"log"
)
import "github.com/spf13/viper"

func NewConfig() *config.Config {
	var configData *config.Config
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error to load config:", err.Error())
		return nil
	}
	if err := viper.Unmarshal(&configData); err != nil {
		log.Fatal("Error to unmarshal config:", err.Error())
		return nil
	}
	return configData
}
