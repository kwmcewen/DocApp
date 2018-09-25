package cmd

import (
	"fmt"
	"os"

	"github.com/hoop33/entrevista"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "docapp",
}

// Execute runs the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createInterview() *entrevista.Interview {
	interview := entrevista.NewInterview()
	/*
		interview.ShowOutput = func(message string) {
			fmt.Print(color.GreenString(message))
		}
		interview.ShowError = func(message string) {
			color.Red(message)
		}
	*/
	return interview
}
