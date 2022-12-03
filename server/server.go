package server

import (
	"fmt"
	"net"

	"github.com/zohymc/zohymc/packets/decider"
	"github.com/zohymc/zohymc/server/types"
	"github.com/zohymc/zohymc/utils"
)

type Server struct {
	*types.BaseServer
	ShouldCompress bool
}

// Listen on port 25565 for incoming connections
func (s Server) CreateServer(options types.ServerOptions) Server {
	s.BaseServer = &types.BaseServer{}

	LoadConfig()

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", options.Address, options.Port))

	// TODO: throw better errors
	if err != nil {
		panic(err)
	}

	for {
		// Accept incoming connections
		conn, err := ln.Accept()

		// TODO: throw better errors
		if err != nil {
			if err.Error() == "EOF" {
				continue
			}
			panic(err)
		}

		// Handle the connection
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	player := &types.Player{
		Conn:  &conn,
		Nick:  "<unknown>",
		State: 0,
	}

	for {
		if player.IsConnected() {
			// Read the next packet
			packet, err := utils.ReadNextPacket(player)

			// check eof
			if err != nil {
				// check if error is "read tcp xxx:xxx use of closed network connection"
				if err.Error() == "EOF" {
					// disconnect the player
					player.ForceDisconnect()
				}
				break
			}

			// create a new player

			decider.Decide(packet, player)
		} else {
			player.ForceDisconnect()
			break
		}
	}
}
