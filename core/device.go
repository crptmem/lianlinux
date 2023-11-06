package core

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/sstallion/go-hid"
)

var Devs []*hid.DeviceInfo

func DeviceEnumerate() {
	readConfig()
	err := hid.Enumerate(0x0cf2, 0xa100, func(info *hid.DeviceInfo) error {
		log.Infof("Found a %s", info.ProductStr)
		Devs = append(Devs, info)
		setConfigLightMode(*info)
		return nil
	})
	if err != nil {
		log.Fatalf("Error during enumerating HID devices: %v", err)
		return
	}
}

func hidWrite(vid uint16, pid uint16, serial string, packet []byte) {
	open, err := hid.Open(vid, pid, serial)
	if err != nil {
		return
	}
	open.Write(packet)
}

func deviceWrite(device hid.DeviceInfo, packet []byte) {
	hidWrite(device.VendorID, device.ProductID, device.SerialNbr, packet)
}

func DeviceInitialize() {
	log.Info("Looking for Lian Li devices...")

	// Enumerate HID devices
	if len(Devs) == 0 {
		log.Error("No supported devices found")
		return
	}

	log.Info(fmt.Sprintf("Found and configured %d device(s)", len(Devs)))
	//rainbowMorph(*Devs[0])
	//static(*Devs[0])
}
