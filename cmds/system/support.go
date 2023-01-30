package systemcmds

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rotisserie/eris"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
)

type SupportCommand struct{}

var (
	_ ken.SlashCommand     = (*SupportCommand)(nil)
	_ ken.DmCapable        = (*SupportCommand)(nil)
	_ cmdhelp.HelpProvider = (*SupportCommand)(nil)
)

func (c *SupportCommand) Name() string {
	return "support"
}

func (c *SupportCommand) Description() string {
	return "Get support for Janus."
}

func (c *SupportCommand) Version() string {
	return "1.0.0"
}

func (c *SupportCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *SupportCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *SupportCommand) IsDmCapable() bool {
	return true
}

func (c *SupportCommand) Help(ctx ken.SubCommandContext) (emb *discordgo.MessageEmbed, err error) {
	emb = &discordgo.MessageEmbed{
		Color:       0x00ff00,
		Description: c.Description(),
	}
	return
}

// func (c *SupportCommand) Guild() string {
// 	return "492075852071174144"
// }

func (c *SupportCommand) Run(ctx ken.Context) (err error) {
	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Title:       "Support",
		Description: "You can get support for Janus via our [Discord server](https://discord.gg/RuEdX5T), or on our [Github](https://github.com/shitcorp/Janus/issues).",
		Timestamp:   time.Now().Format(time.RFC3339),
	}).Send().Error
	err = eris.Wrap(err, "Support cmd response")
	return
}
