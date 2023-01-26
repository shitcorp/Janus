package cmds

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
	//return []*discordgo.ApplicationCommandOption{
	//  {
	//    Type:        discordgo.ApplicationCommandOptionBoolean,
	//    Name:        "pog",
	//    Required:    true,s
	//    Description: "pog",
	//  },
	//}
}

func (c *PingCommand) IsDmCapable() bool {
	return true
}

func (c *PingCommand) Run(ctx ken.Context) (err error) {
	//val := ctx.Options().GetByName("pog").BoolValue()
	//
	//msg := "not poggers"
	//if val {
	//	msg = "poggers"
	//}
	//
	//err = ctx.Respond(&discordgo.InteractionResponse{
	//	Type: discordgo.InteractionResponseChannelMessageWithSource,
	//	Data: &discordgo.InteractionResponseData{
	//		Content: msg,
	//	},
	//})
	//return

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})
	return
}
