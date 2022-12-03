package utils

import (
	"net"

	"github.com/zohymc/zohymc/protocol"
	"github.com/zohymc/zohymc/server/types"
)

type Packet struct {
	Id     int32
	Length int32
	Data   []byte
}

func ReadNextPacket(player *types.Player) (packet Packet, err error) {
	// Read the packet length
	length, _, err := protocol.ReadVarInt(*player.Conn)
	if err != nil {
		return
	}

	length -= 1
	// Read the packet id
	id, _, err := protocol.ReadVarInt(*player.Conn)
	if err != nil {
		return
	}

	// Read the packet data
	data, err := readBytesFromConn(*player.Conn, length)
	if err != nil {
		return
	}

	packet.Id = id
	packet.Length = length
	packet.Data = data

	return
}

func WritePacket(conn net.Conn, packet Packet) {
	var bytes []byte

	packet.Length = int32(len(packet.Data)) + 1

	// Write the packet length
	bytes = append(bytes, protocol.EncodeVarInt(packet.Length)...)
	// Write the packet id
	bytes = append(bytes, protocol.EncodeVarInt(packet.Id)...)
	// Write the packet data
	bytes = append(bytes, packet.Data...)

	// Write the packet to the connection
	conn.Write(bytes)
}

func readBytesFromConn(conn net.Conn, length int32) (data []byte, err error) {
	if length <= 1 {
		return
	}

	data = make([]byte, length)
	_, err = conn.Read(data)

	return
}
