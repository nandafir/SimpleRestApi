package main

import (
	"anaconda/config"
	"anaconda/server"
)

func main() {
	config.Init()
	server.Start()
}
