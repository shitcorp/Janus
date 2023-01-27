package starcitizenCmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/shitcorp/janus/utils"
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

type GameInfoCommand struct{}

var (
	_ ken.SlashCommand = (*GameInfoCommand)(nil)
	_ ken.DmCapable    = (*GameInfoCommand)(nil)
)

func (c *GameInfoCommand) Name() string {
	return "gameinfo"
}

func (c *GameInfoCommand) Description() string {
	return "Get general information about Star Citizen."
}

func (c *GameInfoCommand) Version() string {
	return "1.0.0"
}

func (c *GameInfoCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *GameInfoCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *GameInfoCommand) IsDmCapable() bool {
	return true
}

func (c *GameInfoCommand) Guild() string {
	return "492075852071174144"
}

func (c *GameInfoCommand) Run(ctx ken.Context) (err error) {
	ctx.Defer()

	_, stats, _ := utils.Api.Website.Stats()

	//log.WithFields(log.Fields{
	//  "stats": stats,
	//  "req":   req,
	//}).Info("Stats response")

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Game Info",
					Description: "General information about Star Citizen",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Live",
							Value:  stats.Data.CurrentLive,
							Inline: true,
						},
						{
							Name:   "PTU",
							Value:  stats.Data.CurrentPtu,
							Inline: true,
						},
						{
							Name:   "Evocati",
							Value:  stats.Data.CurrentEtf,
							Inline: true,
						},
						{
							Name:   "Citizens",
							Value:  fmt.Sprintf("%d", stats.Data.Fans),
							Inline: true,
						},
						{
							Name:   "Fleet",
							Value:  fmt.Sprintf("%d", stats.Data.Fleet),
							Inline: true,
						},
						{
							Name:   "Funds",
							Value:  fmt.Sprintf("%d", stats.Data.Funds),
							Inline: true,
						},
					},
				},
			},
			Content: fmt.Sprintf("current live: %s", stats.Data.CurrentLive),
		},
	})
	return
}
