package cmds

import (
  "context"
  "github.com/diamondburned/arikawa/v3/api"
  "github.com/diamondburned/arikawa/v3/api/cmdroute"
  "github.com/diamondburned/arikawa/v3/utils/json/option"
  "math/rand"
  "time"
)

//func init() {
//  c := append(commands, api.CreateCommandData{})
//}

func (h *Handler) cmdThonk(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
  time.Sleep(time.Duration(3+rand.Intn(5)) * time.Second)
  return &api.InteractionResponseData{
    Content: option.NewNullableString("https://tenor.com/view/thonk-thinking-sun-thonk-sun-thinking-sun-gif-14999983"),
  }
}
