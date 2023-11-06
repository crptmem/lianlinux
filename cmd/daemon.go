package cmd

import (
	"github.com/spf13/cobra"
	"lianlinux/server"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run lianlinux as a daemon HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		server.Listen(port)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	daemonCmd.Flags().Int("port", 9341, "Port to listen")
}
