package starcitizenCmds

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rotisserie/eris"
	scapiWebsite "github.com/shitcorp/janus/scapi/website"
	"github.com/shitcorp/janus/utils"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type ShipCommand struct{}

var (
	_ ken.SlashCommand     = (*ShipCommand)(nil)
	_ ken.DmCapable        = (*ShipCommand)(nil)
	_ cmdhelp.HelpProvider = (*ShipCommand)(nil)
)

func (c *ShipCommand) Name() string {
	return "ship"
}

func (c *ShipCommand) Description() string {
	return "Get information about a ship."
}

func (c *ShipCommand) Version() string {
	return "1.0.0"
}

func (c *ShipCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *ShipCommand) Options() []*discordgo.ApplicationCommandOption {
	var minLen *int
	minLen = new(int)
	*minLen = 2

	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "name",
			Required:    true,
			Description: "Name of a specific ship.",
			MinLength:   minLen,
			MaxLength:   200,
		},
	}
}

func (c *ShipCommand) IsDmCapable() bool {
	return true
}

func (c *ShipCommand) Help(ctx ken.SubCommandContext) (emb *discordgo.MessageEmbed, err error) {
	emb = &discordgo.MessageEmbed{
		Color:       0x00ff00,
		Description: c.Description(),
	}
	return
}

// func (c *ShipCommand) Guild() string {
// 	return "492075852071174144"
// }

func (c *ShipCommand) Run(ctx ken.Context) (err error) {
	ctx.Defer()

	name := strings.ToLower(ctx.Options().GetByName("name").StringValue())

	// ships := new([]scapiWebsite.ShipsData)
	// err = utils.Cache.Once(&cache.Item{
	// 	Key:   fmt.Sprintf("api:ship:%s", name),
	// 	Value: ships, // destination
	// 	TTL:   time.Hour * 24,
	// 	Do: func(*cache.Item) (interface{}, error) {
	// 		_, res, err := utils.Api.Website.Ships(scapiWebsite.ShipQuery{
	// 			Name: name,
	// 		})
	// 		if res.Success == 0 {
	// 			return nil, err
	// 		}
	// 		return res.Data, err
	// 	},
	// })
	// if err != nil {
	// 	err = eris.Wrap(err, "Error getting cached player obj")
	// 	return
	// }

	// p := message.NewPrinter(language.English)

	_, res, err := utils.Api.Website.Ships(scapiWebsite.ShipQuery{
		Name: name,
	})
	if res.Success == 0 {
		return
	}

	ship := res.Data[0]

	// utc, err := time.LoadLocation("UTC")
	// if err != nil {
	// 	err = eris.Wrap(err, "Error setting time to UTC")
	// 	return
	// }
	// updated, err := time.ParseInLocation("2006-01-02T15:04:05.000000", ship.TimeModifiedUnfiltered, utc)
	// if err != nil {
	// 	err = eris.Wrap(err, "Error parsing time in player cmd")
	// 	return
	// }

	p := message.NewPrinter(language.English)

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Title: ship.Name,
		// Timestamp: updated.Format("Mon, 02 Jan 2006"),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Last Updated At",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: fmt.Sprintf("http://robertsspaceindustries.com%s", ship.Media),
		},
		URL: fmt.Sprintf("http://robertsspaceindustries.com%s", ship.Url),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Size",
				Value:  ship.Size,
				Inline: true,
			},
			{
				Name:   "Focus",
				Value:  ship.Focus,
				Inline: true,
			},
			{
				Name: "Price",
				// "2016-08-19T00:00:00.000000"
				Value:  p.Sprintf("%d", ship.Price),
				Inline: true,
			},
			{
				Name:   "Max Crew",
				Value:  ship.MaxCrew,
				Inline: true,
			},
			{
				Name:   "Min Crew",
				Value:  ship.MinCrew,
				Inline: true,
			},
			{
				Name:   "Prod Status",
				Value:  ship.ProductionStatus,
				Inline: true,
			},
		},
	}).Send().Error
	err = eris.Wrap(err, "Player cmd response")
	return
}
