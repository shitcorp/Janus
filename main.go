package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shitcorp/janus/cmds"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
	"github.com/zekrotja/ken/store"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Info("Starting Janus")

	session, err := discordgo.New("Bot " + viper.GetString("TOKEN"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infof("Janus is connected as %s", r.User.Username)
	})

	// setup katana
	k, err := ken.New(session, ken.Options{
		CommandStore: store.NewDefault(),
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer k.Unregister()

	// register cmds
	err = k.RegisterCommands(new(cmds.PingCommand))
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
