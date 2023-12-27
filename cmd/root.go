package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "backendboilerplate",
	Short: "Backend server for the backendboilerplate system.",
	Long:  `Backend server for the backendboilerplate system.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("prod", "p", false, "Run in production")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// config.Config(false)
	// don't recreate service file

}
