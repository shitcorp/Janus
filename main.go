package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/botblock/golist"
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
	"github.com/go-co-op/gocron"
	"github.com/rotisserie/eris"
	"github.com/shitcorp/janus/cmds"
	"github.com/shitcorp/janus/utils"
	log "github.com/sirupsen/logrus"
	"github.com/zekroTJA/shinpuru/pkg/rediscmdstore"
	"github.com/zekrotja/ken"
)

// ldflags
// app version #
// OR commit sha
var Release = "dev"

// time of build
var BuildTime = "dev"

// task scheduler
var scheduler = gocron.NewScheduler(time.UTC)

// botblock api client
var botBlock = golist.NewClient()

func main() {
	log.Info("Starting Janus")

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

	// set tog.gg token
	botBlock.AddToken("top.gg", utils.Config.TopGGToken)

	session, err := discordgo.New("Bot " + utils.Config.Token)
	if err != nil {
		err = eris.Wrap(err, "discordgo threw an error")
		log.WithError(err).Fatal("discordgo")
		sentry.CaptureException(err)
	}
	defer session.Close()

	session.AddHandler(readyHandler)

	//_, err = dgrs.New(dgrs.Options{
	//	DiscordSession: session,
	//	RedisClient:    utils.Redis,
	//	FetchAndStore:  true,
	//})

	_, err = scheduler.Every(1).Hour().Do(postBotStats, session)
	if err != nil {
		log.WithError(eris.Wrap(err, "Error in scheduler")).Error()
	}

	// setup ken
	k, err := ken.New(session, ken.Options{
		//CommandStore: store.NewDefault(),
		// use redis to cache cmd info
		CommandStore: rediscmdstore.New(utils.Redis),
		EmbedColors: ken.EmbedColors{
			Default: 0x228dcc,
			Error:   0xF44336,
		},
		OnSystemError:  kenSystemError,
		OnCommandError: kenCmdError,
	})
	if err != nil {
		log.WithError(eris.Wrap(err, "Error setting up ken")).Fatal("ken")
	}

	// register cmds
	err = k.RegisterCommands(cmds.Commands...)
	if err != nil {
		log.WithError(eris.Wrap(err, "Error registering cmds with ken")).Fatal("ken")
	}
	// err = k.RegisterMiddlewares(cmdhelp.New())
	// if err != nil {
	// 	log.WithError(eris.Wrap(err, "Error setting up ken")).Fatal("ken")
	// }

	// login
	err = session.Open()
	if err != nil {
		log.WithError(eris.Wrap(err, "discordgo threw an error while connecting to discord")).Fatal("discordgo")
	}

	// handle stop
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
