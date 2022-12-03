package types

import (
	"net"
)

type ConnectionState uint8

const (
	// this state is only used for the handshake packet. it is not a valid state
	ConnectionStateNotConnected ConnectionState = iota
	ConnectionStateStatus       ConnectionState = iota
	ConnectionStateLogin        ConnectionState = iota
	ConnectionStatePlay         ConnectionState = iota
)

type ServerOptions struct {
	Address        string
	Port           uint16
	ShouldCompress bool
}

type BaseServer struct {
	Listener net.Listener
	Players  []Player
}

// Represents a Minecraft player.
// If the player's nick is <unknown>, then the player is not connected to the Minecraft server but has a TCP connection.
type Player struct {
	Conn  *net.Conn
	Nick  string
	State ConnectionState
}

func (player *Player) ForceDisconnect() {
	(*player.Conn).Close()
}

func (player *Player) IsConnected() bool {
	return (*player.Conn).RemoteAddr() != nil
}
