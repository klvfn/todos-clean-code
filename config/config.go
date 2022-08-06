package config

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort int   `mapstructure:"app_port"`
	Mysql   Mysql `mapstructure:"mysql"`
}

type Mysql struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	DB   string `mapstructure:"db"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

var AppConfig Config

func InitConfig() {
	AppConfig = Config{}
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.UnmarshalExact(&AppConfig, func(c *mapstructure.DecoderConfig) {
		c.ErrorUnset = true
	}); err != nil {
		log.Fatal(err)
	}
}
