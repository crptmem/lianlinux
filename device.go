package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/gousb"
	"github.com/sstallion/go-hid"
)

func deviceEnumerate() []*hid.DeviceInfo {
	var devs []*hid.DeviceInfo
	readConfig()
	hid.Enumerate(0x0cf2, 0xa100, func(info *hid.DeviceInfo) error {
		log.Infof("Found a %s", info.ProductStr)
		devs = append(devs, info)
		setLightMode(*info)
		return nil
	})

	return devs
}

func hidWrite(vid uint16, pid uint16, serial string, packet []byte) {
	open, err := hid.Open(vid, pid, serial)
	if err != nil {
		return
	}
	open.Write(packet)
}

func controlWrite(vid uint16, pid uint16, packet []byte) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Find the USB device with the specified VID and PID
	dev, err := ctx.OpenDeviceWithVIDPID(gousb.ID(vid), gousb.ID(pid))
	if err != nil {
		log.Fatalf("Failed to open USB device: %v", err)
		return
	}
	defer dev.Close()

	// Perform control transfer to send a control packet
	_, err = dev.Control(gousb.ControlVendor|gousb.ControlOut|gousb.ControlDevice, 0x00, 0x00, 0x00, packet)
}

func deviceWrite(device hid.DeviceInfo, packet []byte) {
	hidWrite(device.VendorID, device.ProductID, device.SerialNbr, packet)
	//controlWrite(device.VendorID, device.ProductID, packet)
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
	//rainbowMorph(*devicesInfo[0])
}
