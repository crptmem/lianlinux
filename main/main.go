package main

import (
	"github.com/charmbracelet/log"
	"lianlinux/core"
	"os"
)

func main() {
	environment, _ := os.LookupEnv("ENVIRONMENT")

	if environment == "dev" {
		log.SetLevel(log.DebugLevel)
	}

	core.DeviceInitialize()
	// server.Listen()
}
