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
		r, _ := cmd.Flags().GetInt("red")
		g, _ := cmd.Flags().GetInt("green")
		b, _ := cmd.Flags().GetInt("blue")
		switch mode {
		case "rainbow", "morph":
			log.Infof("Setting mode to %s", mode)
			core.SetLightMode(*core.Devs[0], mode)
		case "static", "breathing":
			log.Infof("Setting mode to %s", mode)
			core.SetLightMode(*core.Devs[0], mode, []byte{byte(r), byte(b), byte(g)}...)
		case "":
			log.Errorf("Please specify a mode")
			os.Exit(1)
		default:
			log.Errorf("Unknown mode: %s", mode)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(lightCmd)

	lightCmd.PersistentFlags().String("mode", "", "RGB light mode")
	lightCmd.PersistentFlags().Int("red", 0, "Red color")
	lightCmd.PersistentFlags().Int("green", 0, "Green color")
	lightCmd.PersistentFlags().Int("blue", 0, "Blue color")
}
