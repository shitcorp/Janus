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

type OrgCommand struct{}

var (
	_ ken.SlashCommand     = (*OrgCommand)(nil)
	_ ken.DmCapable        = (*OrgCommand)(nil)
	_ cmdhelp.HelpProvider = (*OrgCommand)(nil)
)

func (c *OrgCommand) Name() string {
	return "organization"
}

func (c *OrgCommand) Description() string {
	return "Get information about a specific organization."
}

func (c *OrgCommand) Version() string {
	return "1.0.0"
}

func (c *OrgCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *OrgCommand) Options() []*discordgo.ApplicationCommandOption {
	var minLen *int
	minLen = new(int)
	*minLen = 2

	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "sid",
			Required:    true,
			Description: "The org to get information about",
			MinLength:   minLen,
			MaxLength:   200,
		},
	}
}

func (c *OrgCommand) IsDmCapable() bool {
	return true
}

func (c *OrgCommand) Help(ctx ken.SubCommandContext) (emb *discordgo.MessageEmbed, err error) {
	emb = &discordgo.MessageEmbed{
		Color:       0x00ff00,
		Description: c.Description(),
	}
	return
}

func (c *OrgCommand) Guild() string {
	return "492075852071174144"
}

func (c *OrgCommand) Run(ctx ken.Context) (err error) {
	ctx.Defer()

	sid := strings.ToUpper(ctx.Options().GetByName("sid").StringValue())

	org := new(scapiWebsite.OrgData)
	err = utils.Cache.Once(&cache.Item{
		Key:   fmt.Sprintf("api:organization:%s", sid),
		Value: org, // destination
		TTL:   time.Hour * 24,
		Do: func(*cache.Item) (interface{}, error) {
			_, res, err := utils.Api.Website.Org(sid)
			if res.Success == 0 {
				return nil, err
			}

			return res.Data, err
		},
	})
	if err != nil {
		err = eris.Wrap(err, "Error getting cached org obj")
		return
	}

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Title:     org.Name,
		Timestamp: time.Now().Format(time.RFC3339),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: org.Logo,
		},
		Image: &discordgo.MessageEmbedImage{
			URL: org.Banner,
		},
		URL: org.Url,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "SID",
				Value:  org.Sid,
				Inline: true,
			},
			{
				Name:   "Archetype",
				Value:  org.Archetype,
				Inline: true,
			},
			{
				Name:   "Members",
				Value:  fmt.Sprintf("%d", org.Members),
				Inline: true,
			},
			{
				Name:   "Language",
				Value:  org.Lang,
				Inline: true,
			},
			{
				Name:   "Focus",
				Value:  fmt.Sprintf("**Primary**: %s\n**Secondary**: %s", org.Focus.Primary.Name, org.Focus.Secondary.Name),
				Inline: true,
			},
			{
				Name:   "Roleplay",
				Value:  fmt.Sprintf("%t", org.Roleplay),
				Inline: true,
			},
			{
				Name:   "Recruiting",
				Value:  fmt.Sprintf("%t", org.Recruiting),
				Inline: true,
			},
			{
				Name:   "Commitment",
				Value:  org.Commitment,
				Inline: true,
			},
		},
	}).Send().Error
	if err != nil {
		err = eris.Wrap(err, "Player cmd response")
	}
	return
}
