package cmds

import (
  "context"
  "github.com/diamondburned/arikawa/v3/api"
  "github.com/diamondburned/arikawa/v3/api/cmdroute"
  "github.com/diamondburned/arikawa/v3/utils/json/option"
  log "github.com/sirupsen/logrus"
)

func (h *Handler) cmdStats(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {


  _, stats, err := h.api.Website.Stats()
  if err != nil {
    log.Error(err)
  }

  return &api.InteractionResponseData{
    Content:         option.NewNullableString(stats.Data.CurrentLive),
    AllowedMentions: &api.AllowedMentions{}, // don't mention anyone
  }
}