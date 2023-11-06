package main

import (
	"lianlinux/cmd"
	"lianlinux/core"
)

func main() {
	core.DeviceEnumerate()
	cmd.Execute()
}
