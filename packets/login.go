package packets

import (
	"github.com/zohymc/zohymc/protocol"
	"github.com/zohymc/zohymc/server/types"
	"github.com/zohymc/zohymc/utils"
)

func HandleLoginStart(player *types.Player, packet utils.Packet) {
	// Decode the packet
	player.Nick, _ = protocol.DecodeString(packet.Data)

	// Send the response
	SendLoginSuccess(player)
	SendLoginPlay(player)
}

func SendLoginSuccess(player *types.Player) {
	pkt := utils.Packet{
		Id: 0x02,
	}

	// generate a random UUID
	uuid := "OfflinePlayer:1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p"
	numberOfProperties := protocol.EncodeVarInt(0)

	propertiesArray := []byte{}
	// append an empty string
	propertiesArray = append(propertiesArray, protocol.EncodeString("")...)
	propertiesArray = append(propertiesArray, protocol.EncodeString("")...)
	propertiesArray = append(propertiesArray, 0x00)

	finalData := []byte{}
	finalData = append(finalData, uuid...)
	finalData = append(finalData, protocol.EncodeString(player.Nick)...)
	finalData = append(finalData, numberOfProperties...)
	pkt.Data = finalData

	utils.WritePacket(*player.Conn, pkt)
}

func SendLoginPlay(player *types.Player) {
	pkt := utils.Packet{
		Id: 0x25,
	}

	// fake data (entity id, is hardcore, gamemode, previous gamemode, world count, world names, registry codec, dimension, world name, hashed seed, max players, view distance, reduced debug info, enable respawn screen, is debug, is flat)
	var data []byte

	// entity id
	data = append(data, 1)
	// is hardcore
	data = append(data, 0)
	// gamemode
	data = append(data, 0)
	// previous gamemode
	data = append(data, 0)
	// world count
	data = append(data, protocol.EncodeVarInt(1)...)
	// world names (array of strings)
	data = append(data, protocol.EncodeString("minecraft:overworld")...)

	utils.WritePacket(*player.Conn, pkt)
}
