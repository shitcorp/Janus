package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
	"github.com/rotisserie/eris"
	log "github.com/sirupsen/logrus"
	"github.com/zekrotja/ken"
)

func readyHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Infof("Janus is connected as %s#%s", r.User.Username, r.User.Discriminator)

	_ = s.UpdateStreamingStatus(0, "data from the verse", "https://www.youtube.com/watch?v=BbfsX9aFvYw")

	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category: "discordgo",
		Message:  "bot is ready",
		Level:    sentry.LevelInfo,
	})

	// start scheduled tasks
	scheduler.StartAsync()

	// deletes all cmds
	//cmds, err := session.ApplicationCommands(r.Application.ID, "492075852071174144")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for i := range cmds {
	//	//log.Infof("Command: %s", cmds[i].ID)
	//	session.ApplicationCommandDelete(r.Application.ID, "492075852071174144", cmds[i].ID)
	//}
}

// handle ken sys errors
func kenSystemError(context string, err error, args ...interface{}) {
	err = eris.Wrap(err, "error in ken")
	log.WithFields(log.Fields{
		"ctx":  context,
		"args": args,
	}).WithError(err).Error("ken")
	sentry.CaptureException(err)
}

// handle ken cmd errors
func kenCmdError(err error, ctx *ken.Ctx) {
	_ = ctx.Defer()

	if eris.Is(err, ken.ErrNotDMCapable) {
		ctx.FollowUpError("This command cannot be used in dms", "").Send()
		return
	}

	if eris.Is(err, eris.New("SC API Error: no data")) {
		ctx.FollowUpError("Couldn't find one under that name", "").Send()
		return
	}

	ctx.FollowUpError("An error has occurred in Janus, if this continues, please contact Janus's developers.", "").Send()
	log.WithError(err).Error("error in cmd")
	sentry.WithScope(func(scope *sentry.Scope) {
		// so we know the affected user count
		scope.SetUser(sentry.User{
			ID: ctx.User().ID,
		})
		sentry.CaptureException(err)
	})
}
