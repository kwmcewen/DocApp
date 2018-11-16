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
				Key:      "apikey",
				Text:     "Enter your apikey",
				Required: true,
				Hidden:   false,
			},
			{
				Key:      "secretkey",
				Text:     "Enter your secretkey",
				Required: true,
				Hidden:   true,
			},
		}

		answers, err := interview.Run()
		if err != nil {
			panic(err)
		}

		ac := apimedic.NewClient(apimedic.Sandbox, newHttpClient())
		resp, err := ac.LogIn(answers["apikey"].(string), answers["secretkey"].(string))
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
