package main

import "github.com/charmbracelet/log"

const BaseByte byte = 224
const DefaultChannelByte byte = 16

const RgbSyncCmd byte = 48
const PwmCmd byte = 49
const SpeedCmd byte = 0

const (
	RGB   = 0
	PWM   = 1
	SPEED = 2
)

func packetGenerate(cmd int, arg byte, channel ...byte) []byte {
	var packet []byte
	packet = append(packet, BaseByte)
	switch cmd {
	case RGB:
		packet = append(packet, DefaultChannelByte, RgbSyncCmd, arg, 0, 0, 0)
	case PWM:
		packet = append(packet, DefaultChannelByte, PwmCmd, arg)
	case SPEED:
		packet = append(packet, channel[0], SpeedCmd, arg)
	}

	log.Info("packetGenerate:", "packet", packet)

	return packet
}
