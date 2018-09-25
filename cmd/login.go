package cmd

import (
	"fmt"

	"github.com/hoop33/entrevista"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log Into DocApp",
	Long:  "Log into DocApp through ApiMedic",
	Run: func(cmd *cobra.Command, args []string) {
		interview := createInterview()
		interview.Questions = []entrevista.Question{
			{
				Key:      "userId",
				Text:     "Enter your ApiMedic User ID",
				Required: true,
				Hidden:   false,
			},
			{
				Key:      "password",
				Text:     "Enter your ApiMedic Password",
				Required: true,
				Hidden:   true,
			},
		}

		answers, err := interview.Run()
		if err != nil {
			panic(err)
		}
		fmt.Println(answers["userId"])
		fmt.Println(answers["password"])
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
