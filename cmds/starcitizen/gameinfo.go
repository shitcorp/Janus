package starcitizenCmds

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/cache/v8"
	"github.com/rotisserie/eris"
	scapiWebsite "github.com/shitcorp/janus/scapi/website"
	"github.com/shitcorp/janus/utils"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type GameInfoCommand struct{}

var (
	_ ken.SlashCommand     = (*GameInfoCommand)(nil)
	_ ken.DmCapable        = (*GameInfoCommand)(nil)
	_ cmdhelp.HelpProvider = (*GameInfoCommand)(nil)
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

func (c *GameInfoCommand) Help(ctx ken.SubCommandContext) (emb *discordgo.MessageEmbed, err error) {
	emb = &discordgo.MessageEmbed{
		Color:       0x00ff00,
		Description: c.Description(),
	}
	return
}

// func (c *GameInfoCommand) Guild() string {
// 	return "492075852071174144"
// }

func (c *GameInfoCommand) Run(ctx ken.Context) (err error) {
	ctx.Defer()

	stats := new(scapiWebsite.StatsData)
	err = utils.Cache.Once(&cache.Item{
		Key:   "api:gameinfo",
		Value: stats, // destination
		TTL:   time.Hour * 24,
		Do: func(*cache.Item) (interface{}, error) {
			_, res, err := utils.Api.Website.Stats()
			if res.Success == 0 {
				return nil, err
			}
			return res.Data, err
		},
	})
	if err != nil {
		err = eris.Wrap(err, "Error getting cached stats obj")
		return
	}

	p := message.NewPrinter(language.English)

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Title:       "Game Info",
		Description: "General information about Star Citizen",
		Timestamp:   time.Now().Format(time.RFC3339),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Live",
				Value:  stats.CurrentLive,
				Inline: true,
			},
			{
				Name:   "PTU",
				Value:  stats.CurrentPtu,
				Inline: true,
			},
			{
				Name:   "Evocati",
				Value:  stats.CurrentEtf,
				Inline: true,
			},
			{
				Name:   "Citizens",
				Value:  p.Sprintf("%d", stats.Fans),
				Inline: true,
			},
			{
				Name:   "Fleet",
				Value:  p.Sprintf("%d", stats.Fleet),
				Inline: true,
			},
			{
				Name:   "Funds",
				Value:  p.Sprintf("%.2f", stats.Funds),
				Inline: true,
			},
		},
	}).Send().Error
	err = eris.Wrap(err, "Gameinfo cmd response")
	return
}
