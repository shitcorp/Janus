package bot

import (
	"context"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"github.com/shitcorp/janus/utils"
)

func New(logger log.Logger, version string, config utils.ConfigOptions) *Bot {
	return &Bot{
		Logger:  logger,
		Config:  config,
		Version: version,
	}
}

type Bot struct {
	Logger  log.Logger
	Client  bot.Client
	Config  utils.ConfigOptions
	Version string
}

func (b *Bot) SetupBot(listeners ...bot.EventListener) {
	var err error
	b.Client, err = disgo.New(b.Config.Token,
		bot.WithLogger(b.Logger),
		bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentGuilds, gateway.IntentGuildMessages, gateway.IntentMessageContent)),
		bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagGuilds)),
		bot.WithEventListeners(listeners...),
	)
	if err != nil {
		b.Logger.Fatal("Failed to setup b: ", err)
	}
}

func (b *Bot) OnReady(_ *events.Ready) {
	b.Logger.Infof("Butler ready")
	if err := b.Client.SetPresence(context.TODO(), gateway.WithListeningActivity("you"), gateway.WithOnlineStatus(discord.OnlineStatusOnline)); err != nil {
		b.Logger.Errorf("Failed to set presence: %s", err)
	}
}
