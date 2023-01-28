package systemcmds

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rotisserie/eris"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/middlewares/cmdhelp"
)

type VoteCommand struct{}

var (
	_ ken.SlashCommand     = (*VoteCommand)(nil)
	_ ken.DmCapable        = (*VoteCommand)(nil)
	_ cmdhelp.HelpProvider = (*VoteCommand)(nil)
)

func (c *VoteCommand) Name() string {
	return "vote"
}

func (c *VoteCommand) Description() string {
	return "Vote for Janus."
}

func (c *VoteCommand) Version() string {
	return "1.0.0"
}

func (c *VoteCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *VoteCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *VoteCommand) IsDmCapable() bool {
	return true
}

func (c *VoteCommand) Help(ctx ken.SubCommandContext) (emb *discordgo.MessageEmbed, err error) {
	emb = &discordgo.MessageEmbed{
		Color:       0x00ff00,
		Description: c.Description(),
	}
	return
}

func (c *VoteCommand) Guild() string {
	return "492075852071174144"
}

func (c *VoteCommand) Run(ctx ken.Context) (err error) {
	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Title:       "Vote",
		Description: "You can help support Janus by [voting on Top.gg](https://top.gg/bot/776100457260384266/vote), and leaving us a [review](https://top.gg/bot/776100457260384266#reviews).",
		Timestamp:   time.Now().Format(time.RFC3339),
	}).Send().Error
	err = eris.Wrap(err, "Vote cmd response")
	return
}
