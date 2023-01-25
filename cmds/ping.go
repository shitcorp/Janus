package cmds

import (
  "context"
  "github.com/diamondburned/arikawa/v3/api"
  "github.com/diamondburned/arikawa/v3/api/cmdroute"
  "github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (h *Handler) cmdPing(ctx context.Context, cmd cmdroute.CommandData) *api.InteractionResponseData {
	return &api.InteractionResponseData{
		Content: option.NewNullableString("Pong!"),
	}
}
