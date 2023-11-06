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
func changeMode(color []byte, device hid.DeviceInfo) {
	// Iterate through all ports (4) on hub and change their lightning
	for i := 0; i < 4; i++ {
		deviceWrite(device, append([]byte{BaseLightningByte, byte(0x10 + i)}, color...))
	}
}

func static(device hid.DeviceInfo) {
	for i := 0; i < 4; i++ {
		var rgb []byte
		for i := 0; i < 19; i++ {
			rgb = append(rgb, 0x00, 0xff, 0xaa)
		}

		rgb = append(rgb, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00}...)
		colorPacket := append([]byte{0xe0, byte(0x30 + i)}, rgb...)

		deviceWrite(*Devs[0], []byte{0xe0, byte(0x10 + i), 0x32, 0x03})
		deviceWrite(*Devs[0], colorPacket)

		log.Infof("%x", colorPacket)
	}
	changeMode([]byte{StaticFirst, StaticSecond}, device)

}

func rainbow(device hid.DeviceInfo) {
	//changeMode(0x15, 0x23, device)
	changeMode([]byte{Rainbow, RainbowSecond}, device)
}

func rainbowMorph(device hid.DeviceInfo) {
	changeMode([]byte{RainbowMorph, RainbowSecond}, device)
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
