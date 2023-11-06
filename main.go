package main

import (
	"github.com/charmbracelet/log"
	"lianlinux/cmd"
	"lianlinux/core"
	"os"
)

func main() {
	environment, _ := os.LookupEnv("ENVIRONMENT")

	if environment == "dev" {
		log.SetLevel(log.DebugLevel)
	}
	core.DeviceEnumerate()
	cmd.Execute()
}
