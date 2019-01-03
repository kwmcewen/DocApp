package cmd

import (
	"fmt"

	"github.com/hoop33/entrevista"
	"github.com/kwmcewen/docapp/apimedic"
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
				Key:      "username",
				Text:     "Enter your username",
				Required: true,
				Hidden:   false,
			},
			{
				Key:      "password",
				Text:     "Enter your password",
				Required: true,
				Hidden:   true,
			},
		}

		answers, err := interview.Run()
		if err != nil {
			panic(err)
		}

		ac := apimedic.NewClient(apimedic.Sandbox, newHttpClient())
		resp, err := ac.LogIn(answers["username"].(string), answers["password"].(string))
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
