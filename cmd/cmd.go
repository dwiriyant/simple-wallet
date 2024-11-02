package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple-wallet",
	Short: "Simple Wallet REST API",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Run(startCmd, []string{})
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
