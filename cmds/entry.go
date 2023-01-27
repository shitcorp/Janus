package cmds

import (
  "github.com/shitcorp/janus/cmds/starcitizen"
  "github.com/zekrotja/ken"
)

var Commands = []ken.Command{
  new(PingCommand),
  new(starcitizenCmds.GameInfoCommand),
}

//func errorResponse(err error) *api.InteractionResponseData {
//	return &api.InteractionResponseData{
//		Content:         option.NewNullableString("**Error:** " + err.Error()),
//		Flags:           discord.EphemeralMessage,
//		AllowedMentions: &api.AllowedMentions{ /* none */ },
//	}
//}
