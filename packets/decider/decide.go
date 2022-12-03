package decider

import (
	"fmt"

	"github.com/zohymc/zohymc/packets"
	"github.com/zohymc/zohymc/server/types"
	"github.com/zohymc/zohymc/utils"
)

func Decide(packet utils.Packet, player *types.Player) {
	fmt.Println("Packet id:", packet.Id, "Packet length:", packet.Length, "Packet data:", packet.Data)

	if player.State == 0 {
		if packet.Id == 0 {
			packets.HandleHandshake(player)
			return
		}
	}

	if player.State == 1 {
		if packet.Id == 0 {
			packets.HandleStatusRequest(player)
			return
		}
		if packet.Id == 1 {
			packets.HandlePingRequest(player, packet)
			return
		}
	}
}
