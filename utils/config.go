package utils

import (
	"os"

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

func init() {
	//	// Log as JSON instead of the default ASCII formatter.
	//	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	//	// Only log the warning severity or above.
	//	log.SetLevel(log.WarnLevel)

	// if utils.Config.Debug {
	// 	log.SetLevel(log.DebugLevel)
	// }
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config ConfigOptions) {
	// define defaults
	viper.SetDefault("SENTRY_DSN", "")
	viper.SetDefault("DEBUG", false)

	// Read file path
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.SetEnvPrefix("")

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

	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	// log.WithField("Config", config).Debug("Loaded config")
	log.Debugf("Loaded config:\n%+v\n", config)
	return
}
