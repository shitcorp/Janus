package cmds

import (
	starcitizenCmds "github.com/shitcorp/janus/cmds/starcitizen"
	systemcmds "github.com/shitcorp/janus/cmds/system"
	"github.com/zekrotja/ken"
)

var Commands = []ken.Command{
	// star citizen
	new(starcitizenCmds.GameInfoCommand),
	new(starcitizenCmds.PlayerCommand),
	new(starcitizenCmds.OrgCommand),
	// new(starcitizenCmds.ShipCommand),

	// system
	new(systemcmds.PingCommand),
	new(systemcmds.SupportCommand),
	new(systemcmds.VoteCommand),
}
