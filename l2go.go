package main

import (
	"flag"
	"log"
	"runtime"

	"./config"
	"./gameserver"
	_ "./loginserver"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mode, gameServerID int
	flag.IntVar(&mode, "mode", 0, "Set to 0 to run the Login Server or 1 to run the Game Server")
	flag.IntVar(&gameServerID, "server", 1, "Set the ID of the Game Server you want to run")
	flag.Parse()

	// Load the global configuration object
	globalConfig := config.Read()

	if mode == 0 {
		server := loginserver.New(globalConfig)
		server.Init()
		server.Start()
	} else {
		// Try to load the Game Server configuration
		if gameServerID >= 1 && len(globalConfig.GameServers) >= gameServerID {
			config := config.GameServerConfigObject{}
			config.LoginServer = globalConfig.LoginServer
			config.GameServer = globalConfig.GameServers[gameServerID-1]
			server := gameserver.New(config)
			server.Init()
			server.Start()
		} else {
			log.Print("No configuration found for the specified server.")
		}

	}

	log.Print("Server stopped.")
}
