package systemcmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type PingCommand struct{}

var (
	_ ken.SlashCommand = (*PingCommand)(nil)
	_ ken.DmCapable    = (*PingCommand)(nil)
)

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Basic ping command"
}

func (c *PingCommand) Version() string {
	return "1.0.0"
}

func (c *PingCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *PingCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *PingCommand) IsDmCapable() bool {
	return true
}

func (c *PingCommand) Help(ctx ken.SubCommandContext) (emb *discordgo.MessageEmbed, err error) {
	emb = &discordgo.MessageEmbed{
		Color:       0x00ff00,
		Description: c.Description(),
	}
	return
}

func (c *PingCommand) Run(ctx ken.Context) (err error) {
	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})
	return
}
