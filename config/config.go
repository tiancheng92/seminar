package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string `toml:"log_level"`
	Server   struct {
		Host string `toml:"host"`
		Mode string `toml:"mode"`
	} `toml:"server"`
	MySQL struct {
		Dsn string `toml:"dsn"`
	} `toml:"mysql"`
}

var cfg *Config

func GetConf() *Config {
	return cfg
}

func Init() {
	viper.SetConfigFile("./config_file/config.toml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "toml"
	}); err != nil {
		panic(err)
	}
}
