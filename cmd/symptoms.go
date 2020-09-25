package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/hoop33/entrevista"
	"github.com/kwmcewen/docapp/apimedic"
	"github.com/spf13/cobra"
)

var getSymptomsCmd = &cobra.Command{
	Use:   "getSymptoms",
	Short: "Get Symptoms from DocApp",
	Long:  "Get Symptoms from DocApp through ApiMedic",
	Run: func(cmd *cobra.Command, args []string) {
		interview := createInterview()
		interview.Questions = []entrevista.Question{
			{
				Key:      "symptoms",
				Text:     "Enter your symptom, separated by commas",
				Required: false,
				Hidden:   false,
			},
		}

		// answers, err := interview.Run()
		// if err != nil {
		// 	panic(err)
		// }

		tokenContentFromFile, err := ioutil.ReadFile("token.txt")
		if err != nil {
			panic(err)
		}

		// Convert []byte to string and print to screen
		tokenContentString := string(tokenContentFromFile)

		tokenDecoded, err := apimedic.DecodeToken(tokenContentString)
		if err != nil {
			panic(err)
		}

		ac := apimedic.NewClient(apimedic.Sandbox, newHttpClient())
		response, err := ac.RequestAndResponseService("symptoms", tokenDecoded.Value)
		if err != nil {
			panic(err)
		}
		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(getSymptomsCmd)
}
