package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/sstallion/go-hid"
)

func deviceEnumerate() []*hid.DeviceInfo {
	var devs []*hid.DeviceInfo
	hid.Enumerate(0x0cf2, 0xa100, func(info *hid.DeviceInfo) error {
		log.Infof("Found a %s", info.ProductStr)
		devs = append(devs, info)
		return nil
	})

	return devs
}

func deviceWrite(device hid.DeviceInfo, packet []byte) {
	open, err := hid.Open(device.VendorID, device.ProductID, device.SerialNbr)
	if err != nil {
		return
	}
	open.Write(packet)
}

func deviceInitialize() {
	log.Info("Looking for Lian Li devices...")

	// Enumerate HID devices
	devicesInfo := deviceEnumerate()
	if len(devicesInfo) == 0 {
		log.Error("No supported devices found")
		return
	}

	log.Info(fmt.Sprintf("Found %d device(s)", len(devicesInfo)))

	deviceWrite(*devicesInfo[0], packetGenerate(RGB, 0))
	changeMode(0x44, 0xff, *devicesInfo[0])
	//rainbow(*devicesInfo[0])
	//rainbowMorph(*devicesInfo[0])
}
