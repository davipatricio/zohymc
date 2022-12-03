package main

import (
	"github.com/zohymc/zohymc/server"
	"github.com/zohymc/zohymc/server/types"
)

func main() {
	mcserver := server.Server{}
	mcserver.CreateServer(types.ServerOptions{Port: 25565})
}
