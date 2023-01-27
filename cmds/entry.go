package cmds

import (
	"github.com/shitcorp/janus/cmds/starcitizen"
	"github.com/shitcorp/janus/cmds/system"
	"github.com/zekrotja/ken"
)

var Commands = []ken.Command{
	new(systemcmds.PingCommand),
	new(starcitizenCmds.GameInfoCommand),
}
