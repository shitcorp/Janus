package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
	"github.com/rotisserie/eris"
	"github.com/shitcorp/janus/cmds"
	"github.com/shitcorp/janus/utils"
	log "github.com/sirupsen/logrus"
	"github.com/zekroTJA/shinpuru/pkg/rediscmdstore"
	"github.com/zekrotja/ken"
)

func oldMain() {
	log.Info("Starting Janus")

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
