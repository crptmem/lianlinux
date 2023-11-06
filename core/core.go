package core

import (
	"github.com/charmbracelet/log"
	"os"
)

func Run() {
	environment, _ := os.LookupEnv("ENVIRONMENT")

	if environment == "dev" {
		log.SetLevel(log.DebugLevel)
	}

	DeviceInitialize()
	// server.Listen()
}
