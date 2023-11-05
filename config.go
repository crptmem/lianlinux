package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"github.com/sstallion/go-hid"
	"os"
	"path/filepath"
)

func getConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, "/.config")
}

// Check is config present in $HOME/.config/lianlinux, if not - create.
func isConfigPresent() string {
	// Set the config file name and path
	configFileName := "config.json"
	configFilePath := filepath.Join(getConfigDir(), "lianlinux")

	log.Debugf("configFilePath %s configFileName %s", configFilePath, configFileName)

	viper.SetConfigType("json")

	// Set the config file name and path
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFilePath)

	// Check if the config file already exists
	if _, err := os.Stat(filepath.Join(configFilePath, configFileName)); os.IsNotExist(err) {
		// Create the config directory if it doesn't exist
		if err := os.MkdirAll(configFilePath, 0755); err != nil {
			log.Fatalf("Failed to create config directory: %v", err)
			return ""
		}

		// Create a new config file with default values
		if err := viper.SafeWriteConfigAs(filepath.Join(configFilePath, configFileName)); err != nil {
			log.Fatalf("Failed to create config file: %v", err)
			return ""
		}

		log.Info(fmt.Sprintf("Config file created at %s", filepath.Join(configFilePath, configFileName)))
	} else {
		// Read the existing config file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Failed to read config file: %v", err)
			return ""
		}

	}

	return filepath.Join(configFilePath, configFileName)
}

func readConfig() {
	log.Info("Reading configuration...")
	viper.SetConfigFile(isConfigPresent())

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return
	}
}

func setLightMode(device hid.DeviceInfo) {
	currentMode := viper.Get("current")
	log.Infof("Setting mode to %s", currentMode)

	log.Debugf("Setting %s mode for device %s", currentMode, device.ProductStr)

	switch currentMode {
	case "rainbow":
		rainbow(device)
	case "rainbowMorph":
		rainbowMorph(device)
	default:
		rainbow(device)
		log.Warnf("Unknown mode %s, using fallback rainbow", currentMode)
	}
}
