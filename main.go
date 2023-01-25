package main

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"net/http"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"

	"github.com/shitcorp/janus/cmds"
)

func main() {
	LoadConfig(".")

	var (
		webhookAddr   = viper.GetString("WEBHOOK_ADDR")
		webhookPubkey = viper.GetString("WEBHOOK_PUBKEY")
	)

	if webhookAddr != "" {
		s := state.NewAPIOnlyState(viper.GetString("TOKEN"), nil)

		h := cmds.NewHandler(s)

		if err := cmds.OverwriteCommands(s); err != nil {
			log.Fatalln("cannot update commands:", err)
		}

		srv, err := webhook.NewInteractionServer(webhookPubkey, h)
		if err != nil {
			log.Fatalln("cannot create interaction server:", err)
		}

		log.Println("listening and serving at", webhookAddr+"/")
		log.Fatalln(http.ListenAndServe(webhookAddr, srv))
	} else {
		s := state.New("Bot " + viper.GetString("TOKEN"))
		s.AddIntents(gateway.IntentGuilds)
		s.AddHandler(func(*gateway.ReadyEvent) {
			me, _ := s.Me()
			log.Println("connected to the gateway as", me.Tag())
		})

		h := cmds.NewHandler(s)
		s.AddInteractionHandler(h)

		if err := cmds.OverwriteCommands(s); err != nil {
			log.Fatalln("cannot update commands:", err)
		}

		if err := h.State.Connect(context.Background()); err != nil {
			log.Fatalln("cannot connect:", err)
		}
	}
}
