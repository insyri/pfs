package main

import (
	"github.com/spf13/viper"
)

// Searches and loads configuration file
func LoadConfig(f string) (*Config, error) {
	viper.SetConfigName(f)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.UnmarshalExact(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
