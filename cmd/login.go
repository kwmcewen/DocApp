package cmd

import (
	"io"
	"os"
	"path/filepath"

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

		WriteTokenResponseToFile("token.txt", resp)

	},
}

func WriteTokenResponseToFile(filename, data string) error {
	file, err := os.Create(filepath.Join("/Users/kylemcewen/DocApp", filepath.Base(filename)))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
