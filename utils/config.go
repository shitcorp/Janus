package utils

import (
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/iamolegga/enviper"
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
	// for the poor soul wondering wtf is going on in this config
	// there is a really stupid bug in viper https://github.com/spf13/viper/issues/761
	// and enviper is the only decent solution, but it requires some stupid shit to work

	e := enviper.New(viper.New())

	// define defaults
	e.SetDefault("SENTRY_DSN", "")
	e.SetDefault("DEBUG", false)

	// so all env vars start like JANUS_
	e.SetEnvPrefix("JANUS")

	// pull in env vars
	e.AutomaticEnv()

	// so we only use it in docker
	if e.GetString("APP_ENV") == "production" {
		log.Info("Using docker specific config")
		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "config",
			Message:  "Using docker specific config",
			Level:    sentry.LevelInfo,
		})

		// hack so viper can pull env vars
		e.AddConfigPath("/my/config/path")
		e.SetConfigName("config")
	} else {
		// Because I want have a decent dev experiance
		// so this section when "not in prod", uses env files

		log.Info("Not using docker specific config")
		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "config",
			Message:  "Not using docker specific config",
			Level:    sentry.LevelInfo,
		})

		// Read file path
		e.AddConfigPath(path)
		e.SetConfigFile(".env")
		e.SetConfigType("env")

		// read the config file
		err := e.ReadInConfig()
		if err != nil {
			log.Debug("Didn't load config from a file")
			// err = eris.Wrap(err, "failed to read viper config")
			return
		}
	}

	// load config into object
	err := e.Unmarshal(&config)
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
