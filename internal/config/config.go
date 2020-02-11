package config

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	DB *pg.DB
	DBURL string `mapstructure:"database_url"`
	ServerPort int `mapstructure:"port"`
}

func LoadConfig(configPath string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(configPath)
	v.BindEnv("PORT")
	v.BindEnv("DATABASE_URL")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found")
		} else {
			return fmt.Errorf("failed to read the configuration file: %s", err)
		}
	}
	return v.Unmarshal(&Config)
}
