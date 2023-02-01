package utils

import (
	"github.com/rotisserie/eris"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config = LoadConfig(".")

type ConfigOptions struct {
	AppEnv       string `mapstructure:"APP_ENV"`
	Token        string `mapstructure:"TOKEN"`
	ScApiToken   string `mapstructure:"SCAPI_TOKEN"`
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
	SentryDSN    string `mapstructure:"SENTRY_DSN"`
	TopGGToken   string `mapstructure:"TOPGG_TOKEN"`

	// only if using in http mode
	WebhookAddress string `mapstructure:"WEBHOOK_ADDR"`
	WebhookPubkey  string `mapstructure:"WEBHOOK_PUBKEY"`

	// when appenv != production
	TestGuildId string `mapstructure:"TEST_GUILD_ID"`

	// debug  flag
	Debug bool `mapstructure:"DEBUG"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config ConfigOptions) {
	// define defaults
	viper.SetDefault("SENTRY_DSN", "")
	viper.SetDefault("DEBUG", false)

	// Read file path
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	// pull in env vars
	viper.AutomaticEnv()

	// read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Debug("Didn't load config from a file")
		// err = eris.Wrap(err, "failed to read viper config")
		return
	}

	// load config into object
	err = viper.Unmarshal(&config)
	if err != nil {
		log.WithError(eris.Wrap(err, "failed to load the config into an object")).Fatal("config")
	}

	log.Debug("Loaded config")
	return
}
