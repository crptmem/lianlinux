package cmd

import (
	"github.com/spf13/cobra"
	"lianlinux/core"
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
		core.SetLightMode(*core.Devs[0], mode, []byte{byte(r), byte(b), byte(g)}...)
	},
}

func init() {
	rootCmd.AddCommand(lightCmd)

	lightCmd.PersistentFlags().String("mode", "", "RGB light mode")
	lightCmd.PersistentFlags().Int("red", 0, "Red color")
	lightCmd.PersistentFlags().Int("green", 0, "Green color")
	lightCmd.PersistentFlags().Int("blue", 0, "Blue color")
}
