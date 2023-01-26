package cmds

import (
  "context"
  "github.com/bwmarrin/discordgo"
  "github.com/diamondburned/arikawa/v3/api"
  "github.com/diamondburned/arikawa/v3/api/cmdroute"
  "github.com/diamondburned/arikawa/v3/utils/json/option"
  log "github.com/sirupsen/logrus"
  "github.com/zekrotja/ken"
)

//func (h *Handler) cmdStats(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
//
//	_, stats, err := h.api.Website.Stats()
//	if err != nil {
//		log.Error(err)
//	}
//
//	return &api.InteractionResponseData{
//		Content:         option.NewNullableString(stats.Data.CurrentLive),
//		AllowedMentions: &api.AllowedMentions{}, // don't mention anyone
//	}
//}

type StatsCommand struct{}

var (
  _ ken.SlashCommand = (*StatsCommand)(nil)
  _ ken.DmCapable    = (*StatsCommand)(nil)
)

func (c *StatsCommand) Name() string {
  return "stats"
}

func (c *StatsCommand) Description() string {
  return "Basic ping command"
}

func (c *StatsCommand) Version() string {
  return "1.0.0"
}

func (c *StatsCommand) Type() discordgo.ApplicationCommandType {
  return discordgo.ChatApplicationCommand
}

func (c *StatsCommand) Options() []*discordgo.ApplicationCommandOption {
  return []*discordgo.ApplicationCommandOption{}
}

func (c *StatsCommand) IsDmCapable() bool {
  return true
}

func (c *StatsCommand) Run(ctx ken.Context) (err error) {

  err = ctx.Respond(&discordgo.InteractionResponse{
    Type: discordgo.InteractionResponseChannelMessageWithSource,
    Data: &discordgo.InteractionResponseData{
      Content: "pong",
    },
  })
  return
}
