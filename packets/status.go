package packets

import (
	"github.com/zohymc/zohymc/protocol"
	"github.com/zohymc/zohymc/server/types"
	"github.com/zohymc/zohymc/utils"
)

func HandleStatusRequest(player *types.Player) {
	SendStatusResponse(player)
}

func SendStatusResponse(player *types.Player) {
	pkt := utils.Packet{}

	jsonString := `{"version":{"name":"1.19","protocol":759},"players":{"max":1,"online":0,"sample":[]},"description":{"text":"§fZohyMC"},"favicon":""}`

	pkt.Data = protocol.EncodeString(jsonString)

	utils.WritePacket(*player.Conn, pkt)
}

func HandlePingRequest(player *types.Player, packet utils.Packet) {
	SendPongResponse(player, packet)
}

// Send the same VarLong that the client sent
func SendPongResponse(player *types.Player, packet utils.Packet) {
	pkt := utils.Packet{
		Id:   1,
		Data: packet.Data,
	}

	defer player.ForceDisconnect()
	utils.WritePacket(*player.Conn, pkt)
}
