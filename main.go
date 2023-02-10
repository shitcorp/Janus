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

	scheduler := gocron.NewScheduler(time.UTC)
	botBlock := golist.NewClient()

	botBlock.AddToken("top.gg", utils.Config.TopGGToken)

	session, err := discordgo.New("Bot " + utils.Config.Token)
	if err != nil {
		err = eris.Wrap(err, "discordgo threw an error")
		log.WithError(err).Fatal("discordgo")
		sentry.CaptureException(err)
	}
	defer session.Close()

	//_, err = dgrs.New(dgrs.Options{
	//	DiscordSession: session,
	//	RedisClient:    utils.Redis,
	//	FetchAndStore:  true,
	//})

	_, err = scheduler.Every(1).Hour().Do(func() error {
		if utils.Config.AppEnv != "production" {
			return nil
		}

		_, err := botBlock.PostStats(session.State.User.ID, golist.Stats{
			ServerCount: int64(len(session.State.Guilds)),
			ShardID:     int64(session.ShardID),
			ShardCount:  int64(session.ShardCount),
		})
		return eris.Wrap(err, "Error with botblock")
	})
	if err != nil {
		log.WithError(eris.Wrap(err, "Error in scheduler")).Error()
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infof("Janus is connected as %s#%s", r.User.Username, r.User.Discriminator)

		_ = s.UpdateStreamingStatus(0, "data from the verse", "https://www.youtube.com/watch?v=BbfsX9aFvYw")

		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "discordgo",
			Message:  "bot is ready",
			Level:    sentry.LevelInfo,
		})

		scheduler.StartAsync()

		// deletes all cmds
		//cmds, err := session.ApplicationCommands(r.Application.ID, "492075852071174144")
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//for i := range cmds {
		//	//log.Infof("Command: %s", cmds[i].ID)
		//	session.ApplicationCommandDelete(r.Application.ID, "492075852071174144", cmds[i].ID)
		//}
	})

	// setup ken
	k, err := ken.New(session, ken.Options{
		//CommandStore: store.NewDefault(),
		// use redis to cache cmd info
		CommandStore: rediscmdstore.New(utils.Redis),
		EmbedColors: ken.EmbedColors{
			Default: 0x228dcc,
			Error:   0xF44336,
		},
		OnSystemError: func(context string, err error, args ...interface{}) {
			err = eris.Wrap(err, "error in ken")
			log.WithFields(log.Fields{
				"ctx":  context,
				"args": args,
			}).WithError(err).Error("ken")
			sentry.CaptureException(err)
		},
		OnCommandError: func(err error, ctx *ken.Ctx) {
			_ = ctx.Defer()

			if eris.Is(err, ken.ErrNotDMCapable) {
				ctx.FollowUpError("This command cannot be used in dms", "").Send()
				return
			}

			if eris.Is(err, eris.New("SC API Error: no data")) {
				ctx.FollowUpError("Couldn't find one under that name", "").Send()
				return
			}

			ctx.FollowUpError("An error has occurred in Janus, if this continues, please contact Janus's developers.", "").Send()
			log.WithError(err).Error("error in cmd")
			sentry.CaptureException(err)
		},
	})
	if err != nil {
		log.WithError(eris.Wrap(err, "Error setting up ken")).Fatal("ken")
	}
	defer k.Unregister()

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
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
