package main

import (
	"github.com/sstallion/go-hid"
)

const BaseLightningByte = 0xE0

// I don't know how to properly name these bytes
func changeMode(first byte, second byte, device hid.DeviceInfo) {
	// Iterate through all ports (4) on hub and change their lightning
	for i := 0; i < 4; i++ {
		deviceWrite(device, []byte{BaseLightningByte, byte(0x10 + i), first, second})
	}
}

func rainbow(device hid.DeviceInfo) {
	//changeMode(0x15, 0x23, device)
	changeMode(0x05, 0xff, device)
}

func rainbowMorph(device hid.DeviceInfo) {
	changeMode(0x04, 0xff, device)
}
