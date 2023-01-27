package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shitcorp/janus/cmds"
	"github.com/shitcorp/janus/utils"
	log "github.com/sirupsen/logrus"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
	"github.com/zekrotja/ken/store"
	"os"
	"os/signal"
	"syscall"
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
		panic(err)
	}
	defer session.Close()

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infof("Janus is connected as %s#%s", r.User.Username, r.User.Discriminator)
	})

	// setup ken
	k, err := ken.New(session, ken.Options{
		CommandStore: store.NewDefault(),
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer k.Unregister()

	// register cmds
	err = k.RegisterCommands(cmds.Commands...)
	if err != nil {
		log.Fatalln(err)
	}
	err = k.RegisterMiddlewares(cmdhelp.New())
	if err != nil {
		log.Fatalln(err)
	}

	// login
	err = session.Open()
	if err != nil {
		log.Fatalln(err)
	}

	// handle stop
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
