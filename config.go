package main

import (
	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) () {
	// Read file path
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	// pull in env vars
	viper.AutomaticEnv()

	// read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
