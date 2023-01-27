package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config = LoadConfig(".")

type ConfigOptions struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	Token      string `mapstructure:"TOKEN"`
	ScApiToken string `mapstructure:"SCAPI_TOKEN"`

	// only if using in http mode
	WebhookAddress string `mapstructure:"WEBHOOK_ADDR"`
	WebhookPubkey  string `mapstructure:"WEBHOOK_PUBKEY"`

	// when appenv != production
	TestGuildId string `mapstructure:"TEST_GUILD_ID"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config ConfigOptions) {
	// define defaults

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

	// load config into object
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

	log.Info("Loaded config")
	return
}
