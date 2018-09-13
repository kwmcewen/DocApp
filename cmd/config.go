package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure docapp",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from config")
	},
}

func init() {
	//var Config string
	rootCmd.AddCommand(configCmd)
	rootCmd.PersistentFlags().StringP("config", "c", "", "config of docapp")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}
