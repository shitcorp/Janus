package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv string `mapstructure:"APP_ENV"`
	Token  string `mapstructure:"TOKEN"`
	ScApi  string `mapstructure:"SCAPI_TOKEN"`

	// only if using in http mode
	WebhookAddress string `mapstructure:"WEBHOOK_ADDR"`
	WebhookPubkey  string `mapstructure:"WEBHOOK_PUBKEY"`

	// when appenv != production
	TestGuildId string `mapstructure:"TEST_GUILD_ID"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) {
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

	log.Info("Loaded the config")
}
