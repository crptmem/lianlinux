/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"lianlinux/cmd"
	"lianlinux/core"
)

func main() {
	core.DeviceEnumerate()
	cmd.Execute()
}
