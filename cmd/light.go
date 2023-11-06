package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"lianlinux/core"
	"os"
)

// lightCmd represents the light command
var lightCmd = &cobra.Command{
	Use:   "light",
	Short: "Control RGB lights",
	Run: func(cmd *cobra.Command, args []string) {
		mode, _ := cmd.Flags().GetString("mode")
		switch mode {
		case "rainbow", "morph", "static":
			log.Infof("Setting mode to %s", mode)
			core.SetLightMode(*core.Devs[0], mode)
		default:
			log.Errorf("Unknown mode: %s", mode)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(lightCmd)

	lightCmd.PersistentFlags().String("mode", "rainbow", "RGB light mode")
}
