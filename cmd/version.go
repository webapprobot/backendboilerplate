package cmd

import (
	"fmt"

	"github.com/webappbot/backendboilerplate/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "backendboilerplate API version",
	Long:  `backendboilerplate API version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
