package core

import (
	"bytes"
	"github.com/charmbracelet/log"
	"github.com/sstallion/go-hid"
)

const BaseLightningByte = 0xE0

const Rainbow = 0x05
const RainbowMorph = 0x04
const RainbowSecond = 0xFF

const StaticFirst = 0x01
const StaticSecond = 0x02

// Change device light mode by sending mode bytes
func changeMode(color []byte, device hid.DeviceInfo) {
	// Iterate through all ports (4) on hub and change their lightning
	for i := 0; i < 4; i++ {
		DeviceWrite(device, append([]byte{BaseLightningByte, byte(0x10 + i)}, color...))
	}
}

// Set device light mode to Static
func static(device hid.DeviceInfo, color []byte) {
	for i := 0; i < 4; i++ {
		var rgb []byte
		for i := 0; i < 38; i++ {
			rgb = append(rgb, color...)
		}

		rgb = append(rgb, bytes.Repeat([]byte{0x00}, 200)...)
		colorPacket := append([]byte{0xe0, byte(0x30 + i)}, rgb...)

		DeviceWrite(*Devs[0], []byte{0xe0, byte(0x10 + i), 0x32, 0x03})
		DeviceWrite(*Devs[0], colorPacket)
	}
	changeMode([]byte{StaticFirst, StaticSecond}, device)

}

// Set device light mode to Rainbow
func rainbow(device hid.DeviceInfo) {
	changeMode([]byte{Rainbow, RainbowSecond}, device)
}

// Set device light mode to Morph
func rainbowMorph(device hid.DeviceInfo) {
	changeMode([]byte{RainbowMorph, RainbowSecond}, device)
}

func SetLightMode(device hid.DeviceInfo, mode string, rgb ...byte) {
	log.Debugf("Setting %s mode for device %s", mode, device.ProductStr)

	switch mode {
	case "rainbow":
		rainbow(device)
	case "morph":
		rainbowMorph(device)
	case "static":
		static(device, rgb)
	default:
		rainbow(device)
		log.Warnf("Unknown mode %s, using fallback rainbow", mode)
	}
}
