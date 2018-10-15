package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hoop33/entrevista"
	"github.com/spf13/cobra"
)

var insecure bool
var timeout int

func init() {
	flags := rootCmd.PersistentFlags()
	flags.BoolVarP(&insecure, "insecure", "i", false, "skip certificate verification")
	flags.IntVarP(&timeout, "timeout", "t", 30, "http timeout on request")
}

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecure,
			},
		},
	}
}

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
