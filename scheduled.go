package main

import (
	"github.com/botblock/golist"
	"github.com/bwmarrin/discordgo"
	"github.com/rotisserie/eris"
	"github.com/shitcorp/janus/utils"
)

func postBotStats(session *discordgo.Session) error {
	if utils.Config.AppEnv != "production" {
		return nil
	}

	_, err := botBlock.PostStats(session.State.User.ID, golist.Stats{
		ServerCount: int64(len(session.State.Guilds)),
		ShardID:     int64(session.ShardID),
		ShardCount:  int64(session.ShardCount),
	})
	return eris.Wrap(err, "Error with botblock")
}
