package packets

import (
	"github.com/zohymc/zohymc/server/types"
)

func HandleHandshake(player *types.Player) {
	player.State = 1
}
