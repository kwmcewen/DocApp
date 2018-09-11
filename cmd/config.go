package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from config")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
