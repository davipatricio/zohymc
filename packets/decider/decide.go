package decider

import (
	"fmt"

	"github.com/zohymc/zohymc/packets"
	"github.com/zohymc/zohymc/server/types"
	"github.com/zohymc/zohymc/utils"
)

func Decide(packet utils.Packet, player *types.Player) {
	fmt.Println("Player name:", player.Nick, "State:", player.State, "|", "Packet id:", packet.Id, "Packet length:", packet.Length, "Packet data:", packet.Data)

	switch player.State {
	case types.ConnectionStateNotConnected:
		handleNotConnectedPackets(packet, player)
	case types.ConnectionStateStatus:
		handleStatusPackets(packet, player)
	}
}

func handleNotConnectedPackets(packet utils.Packet, player *types.Player) {
	switch packet.Id {
	case 0x00:
		packets.HandleHandshake(player, packet.Data)
	}
}

func handleStatusPackets(packet utils.Packet, player *types.Player) {
	switch packet.Id {
	case 0x00:
		packets.HandleStatusRequest(player)
	case 0x01:
		packets.HandlePingRequest(player, packet)
	}
}

func handleLoginPackets(packet utils.Packet, player *types.Player) {
	switch packet.Id {
	case 0x00:
		packets.HandleLoginStart(player, packet)
	}
}
