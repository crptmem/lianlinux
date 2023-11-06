package core

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/sstallion/go-hid"
	"strings"
)

var Devs []*hid.DeviceInfo

// DeviceEnumerate enumerates HID devices and founds supported device by VID and PID
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

// Write a packet to USB HID device
func hidWrite(vid uint16, pid uint16, serial string, packet []byte) {
	open, err := hid.Open(vid, pid, serial)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%v", err), "Permission denied") {
			log.Fatalf("Please run lianlinux with elevated permissions")
			return
		}
		log.Fatal(fmt.Sprintf("Error during writing to device %x:%x (%s)", vid, pid, serial), "error", err)
		return
	}
	open.Write(packet)
}

// DeviceWrite Write a packet to supported Lian Li device
func DeviceWrite(device hid.DeviceInfo, packet []byte) {
	log.Debugf("Packet: %x", packet)
	hidWrite(device.VendorID, device.ProductID, device.SerialNbr, packet)
}

func DeviceInitialize() {
	if len(Devs) == 0 {
		log.Error("No supported devices found")
		return
	}

	log.Info(fmt.Sprintf("Found and configured %d device(s)", len(Devs)))
}
