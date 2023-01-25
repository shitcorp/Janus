package cmds

import (
  "context"
  "github.com/diamondburned/arikawa/v3/api"
  "github.com/diamondburned/arikawa/v3/api/cmdroute"
  "github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (h *Handler) cmdEcho(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
  var options struct {
    Arg string `discord:"argument"`
  }

  if err := data.Options.Unmarshal(&options); err != nil {
    return errorResponse(err)
  }

  return &api.InteractionResponseData{
    Content:         option.NewNullableString(options.Arg),
    AllowedMentions: &api.AllowedMentions{}, // don't mention anyone
  }
}