package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	DBURL string `mapstructure:"database_url"`
	ServerPort int `mapstructure:"port"`
}

func LoadConfig(configPath string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.AutomaticEnv()
	v.AddConfigPath(configPath)
	
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}