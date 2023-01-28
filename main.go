package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rotisserie/eris"
	"github.com/shitcorp/janus/cmds"
	"github.com/shitcorp/janus/utils"
	log "github.com/sirupsen/logrus"
	"github.com/zekroTJA/shinpuru/pkg/rediscmdstore"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
)

func init() {
	//	// Log as JSON instead of the default ASCII formatter.
	//	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	//	// Only log the warning severity or above.
	//	log.SetLevel(log.WarnLevel)
}

func main() {
	log.Info("Starting Janus")

	session, err := discordgo.New("Bot " + utils.Config.Token)
	if err != nil {
		log.WithError(eris.Wrap(err, "discordgo threw an error")).Fatal("discordgo")
	}
	defer session.Close()

	//_, err = dgrs.New(dgrs.Options{
	//	DiscordSession: session,
	//	RedisClient:    utils.Redis,
	//	FetchAndStore:  true,
	//})

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infof("Janus is connected as %s#%s", r.User.Username, r.User.Discriminator)

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
	err = k.RegisterMiddlewares(cmdhelp.New())
	if err != nil {
		log.WithError(eris.Wrap(err, "Error setting up ken")).Fatal("ken")
	}

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
