package core

import (
	"github.com/charmbracelet/log"
	"github.com/sstallion/go-hid"
)

const BaseLightningByte = 0xE0

const Rainbow = 0x05
const RainbowMorph = 0x04
const RainbowSecond = 0xFF

const StaticFirst = 0x01
const StaticSecond = 0x02

// I don't know how to properly name these bytes
func changeMode(first byte, second byte, device hid.DeviceInfo) {
	// Iterate through all ports (4) on hub and change their lightning
	for i := 0; i < 4; i++ {
		deviceWrite(device, []byte{BaseLightningByte, byte(0x10 + i), first, second})
	}
}

func static(device hid.DeviceInfo) {
	changeMode(StaticFirst, StaticSecond, device)
}

func rainbow(device hid.DeviceInfo) {
	//changeMode(0x15, 0x23, device)
	changeMode(Rainbow, RainbowSecond, device)
}

func rainbowMorph(device hid.DeviceInfo) {
	changeMode(RainbowMorph, RainbowSecond, device)
}

func SetLightMode(device hid.DeviceInfo, mode string) {
	log.Debugf("Setting %s mode for device %s", mode, device.ProductStr)

	switch mode {
	case "rainbow":
		rainbow(device)
	case "morph":
		rainbowMorph(device)
	case "static":
		static(device)
	default:
		rainbow(device)
		log.Warnf("Unknown mode %s, using fallback rainbow", mode)
	}
}
