package starcitizenCmds

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/cache/v8"
	"github.com/rotisserie/eris"
	scapiWebsite "github.com/shitcorp/janus/scapi/website"
	"github.com/shitcorp/janus/utils"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
)

type PlayerCommand struct{}

var (
	_ ken.SlashCommand     = (*PlayerCommand)(nil)
	_ ken.DmCapable        = (*PlayerCommand)(nil)
	_ cmdhelp.HelpProvider = (*PlayerCommand)(nil)
)

func (c *PlayerCommand) Name() string {
	return "player"
}

func (c *PlayerCommand) Description() string {
	return "Get information about a specific player."
}

func (c *PlayerCommand) Version() string {
	return "1.0.0"
}

func (c *PlayerCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *PlayerCommand) Options() []*discordgo.ApplicationCommandOption {
	var minLen *int
	minLen = new(int)
	*minLen = 2

	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "handle",
			Required:    true,
			Description: "The player to get information about",
			MinLength:   minLen,
			MaxLength:   200,
		},
	}
}

func (c *PlayerCommand) IsDmCapable() bool {
	return true
}

func (c *PlayerCommand) Help(ctx ken.SubCommandContext) (emb *discordgo.MessageEmbed, err error) {
	emb = &discordgo.MessageEmbed{
		Color:       0x00ff00,
		Description: c.Description(),
	}
	return
}

func (c *PlayerCommand) Guild() string {
	return "492075852071174144"
}

func (c *PlayerCommand) Run(ctx ken.Context) (err error) {
	ctx.Defer()

	handle := strings.ToLower(ctx.Options().GetByName("handle").StringValue())

	player := new(scapiWebsite.UserData)
	err = utils.Cache.Once(&cache.Item{
		Key:   fmt.Sprintf("player:%s", handle),
		Value: player, // destination
		TTL:   time.Hour * 24,
		Do: func(*cache.Item) (interface{}, error) {
			_, res, err := utils.Api.Website.User(handle)
			if res.Success == 0 {
				err = eris.New(fmt.Sprintf("SC API Error: %s", res.Message))
			}
			return res.Data, err
		},
	})
	if err != nil {
		err = eris.Wrap(err, "Error getting cached player obj")
		return
	}

	// p := message.NewPrinter(language.English)

	utc, err := time.LoadLocation("UTC")
	if err != nil {
		err = eris.Wrap(err, "Error setting time to UTC")
		return
	}
	enlisted, err := time.ParseInLocation("2006-01-02T15:04:05.000000", player.Profile.Enlisted, utc)
	if err != nil {
		err = eris.Wrap(err, "Error parsing time in player cmd")
		return
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Username",
			Value:  player.Profile.Handle,
			Inline: true,
		},
		{
			Name:   "Badge",
			Value:  player.Profile.Badge,
			Inline: true,
		},
		{
			Name: "Enlisted",
			// "2016-08-19T00:00:00.000000"
			Value:  enlisted.Format(time.RFC1123),
			Inline: true,
		},
		{
			Name:   "Fluent",
			Value:  strings.Join(player.Profile.Fluency, ", "),
			Inline: true,
		},
	}

	var noOrg scapiWebsite.UserDataOrg
	if player.Organization != noOrg {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Main Organization",
			Value:  fmt.Sprintf("**Name**: %s\n**Rank**: %s", player.Organization.Name, player.Organization.Rank),
			Inline: true,
		})
	}

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Title: player.Profile.Display,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: player.Profile.Image,
		},
		URL:    player.Profile.Page.Url,
		Fields: fields,
	}).Send().Error
	err = eris.Wrap(err, "Player cmd response")
	return
}
