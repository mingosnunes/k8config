/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		basePath := os.Getenv("HOME") + "/.kube/configs"

		files, err := ioutil.ReadDir(basePath)

		if err != nil {
			log.Fatal(err)
		}

		options := make([]string, 0)

		for _, f := range files {
			options = append(options, f.Name())
		}

		config := ""
		prompt := &survey.Select{
			Message: "Choose a kube config:",
			Options: options,
		}

		survey.AskOne(prompt, &config)

		fmt.Println("Run this command: export KUBECONFIG=" + basePath + "/" + config)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
