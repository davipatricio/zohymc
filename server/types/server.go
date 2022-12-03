package types

import (
	"net"
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

type Player struct {
	Conn  *net.Conn
	Nick  string
	State uint8
}

func (player *Player) ForceDisconnect() {
	(*player.Conn).Close()
}

func (player *Player) IsConnected() bool {
	return (*player.Conn).RemoteAddr() != nil
}
