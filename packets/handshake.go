package packets

import (
	"github.com/zohymc/zohymc/protocol"
	"github.com/zohymc/zohymc/server/types"
)

type HandshakeData struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       types.ConnectionState
}

func HandleHandshake(player *types.Player, data []byte) {
	handshakeData := DecodeStatusHandshake(data)
	player.State = handshakeData.NextState
}

func DecodeStatusHandshake(data []byte) (handshake HandshakeData) {
	// read protocol version (varint, up to 5 bytes)
	protocolVersion, endBytePosition, _ := protocol.DecodeVarInt(data)

	// read server address (string)
	serverAddress, endBytePosition := protocol.DecodeString(data[endBytePosition:])

	// read server port (unsigned short, 2 bytes)
	serverPort := protocol.DecodeUnsignedShort(data[endBytePosition : endBytePosition+2])

	// manually increment byte position
	endBytePosition += 2

	// read next state (varint, up to 5 bytes)
	nextState, _, _ := protocol.DecodeVarInt(data[endBytePosition:])

	// set handshake data
	handshake.ProtocolVersion = protocolVersion
	handshake.ServerAddress = serverAddress
	handshake.ServerPort = serverPort
	handshake.NextState = types.ConnectionState(nextState)

	return
}
