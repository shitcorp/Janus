package utils

import (
	"github.com/bwmarrin/discordgo"
)

var GeneralErrorResponse = &discordgo.InteractionResponse{
	Type: discordgo.InteractionResponseChannelMessageWithSource,
	Data: &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			{
				Color:       0xF44336,
				Description: "An error has occurred in Janus, if this continues, please contact Janus's developers.",
			},
		},
	},
}
