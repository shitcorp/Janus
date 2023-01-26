package cmds

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/shitcorp/janus/scapi"
	"github.com/spf13/viper"
	"log"
)

var commands = []api.CreateCommandData{
	{
		Name:        "ping",
		Description: "ping pong!",
	},
	{
		Name:        "echo",
		Description: "echo back the argument",
		Options: []discord.CommandOption{
			&discord.StringOption{
				OptionName:  "argument",
				Description: "what's echoed back",
				Required:    true,
			},
		},
	},
	{
		Name:        "thonk",
		Description: "biiiig thonk",
	},
	{
		Name:        "stats",
		Description: "Gets star citizen stats",
	},
}

func OverwriteCommands(s *state.State) error {

	if viper.GetString("APP_ENV") == "production" {
		return cmdroute.OverwriteCommands(s, commands)
	} else {
		app, err := s.CurrentApplication()
		if err != nil {
			log.Fatal(err)
		}

		snowflake, err := discord.ParseSnowflake(viper.GetString("TEST_GUILD_ID"))
		if err != nil {
			log.Fatal(err)
		}

		guildID := discord.GuildID(snowflake)

		_, err = s.BulkOverwriteGuildCommands(app.ID, guildID, commands)
		return err
	}
}

type Handler struct {
	*cmdroute.Router
	State *state.State
	api   *scapi.Api
}

func NewHandler(s *state.State) *Handler {
	h := &Handler{State: s}

	h.api = scapi.CreateApiClient(viper.GetString("SCAPI_TOKEN"), "Janus Bot")

	h.Router = cmdroute.NewRouter()
	// Automatically defer handles if they're slow.
	h.Use(cmdroute.Deferrable(s, cmdroute.DeferOpts{}))
	h.AddFunc("echo", h.cmdEcho)
	h.AddFunc("thonk", h.cmdThonk)
	h.AddFunc("stats", h.cmdStats)

	return h
}

func errorResponse(err error) *api.InteractionResponseData {
	return &api.InteractionResponseData{
		Content:         option.NewNullableString("**Error:** " + err.Error()),
		Flags:           discord.EphemeralMessage,
		AllowedMentions: &api.AllowedMentions{ /* none */ },
	}
}
