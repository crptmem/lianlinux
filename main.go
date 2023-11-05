package main

import (
	"github.com/charmbracelet/log"
	"os"
)

func main() {
	environment, _ := os.LookupEnv("ENVIRONMENT")

	if environment == "dev" {
		log.SetLevel(log.DebugLevel)
	}

	deviceInitialize()
}
