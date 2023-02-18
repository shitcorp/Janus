package main

import (
	"time"

	"github.com/botblock/golist"
	"github.com/getsentry/sentry-go"
	"github.com/go-co-op/gocron"
	"github.com/rotisserie/eris"
	"github.com/shitcorp/janus/bot"
	"github.com/shitcorp/janus/utils"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// ldflags
// app version #
// OR commit sha
var Release = "dev"

// time of build
var BuildTime = "dev"

func main() {
	// because I plan to overhaul the config later
	config := utils.Config

	// sentry options
	sentryOptions := sentry.ClientOptions{
		Release: Release,

		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}
	// if we specified the sentry DSN
	if utils.Config.SentryDSN != "" {
		sentryOptions.Dsn = utils.Config.SentryDSN
	}
	// init sentry
	err := sentry.Init(sentryOptions)
	if err != nil {
		log.WithError(err).Fatal("sentry init")
	}
	// still allow reports to be sent if a panic happens
	defer sentry.Recover()

	b := bot.New(logrus.New(), Release, config)

	// init botblock.org client
	botBlock := golist.NewClient()
	// set bot list tokens
	botBlock.AddToken("top.gg", config.TopGGToken)
	botBlock.AddToken("discordbotlist.com", config.DBotListToken)

	// init task scheduler
	scheduler := gocron.NewScheduler(time.UTC)

	_, err = scheduler.Every(1).Hour().Do(postBotStats)
	if err != nil {
		log.WithError(eris.Wrap(err, "Error in scheduler")).Error()
	}
}

func postBotStats() error {
	if utils.Config.AppEnv != "production" {
		return nil
	}

	// _, err := botBlock.PostStats(session.State.User.ID, golist.Stats{
	// 	ServerCount: int64(len(session.State.Guilds)),
	// 	ShardID:     int64(session.ShardID),
	// 	ShardCount:  int64(session.ShardCount),
	// })
	// return eris.Wrap(err, "Error with botblock")
	return nil
}
