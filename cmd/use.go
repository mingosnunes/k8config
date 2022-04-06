/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"k8config/utils"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Change the active Kubernetes configuration file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		checks := utils.CheckInstallation()

		if len(checks) > 0 {
			utils.PrintRed.Println("\n⚠️ k8config is not installed correctly. Run ➡️ k8config install")
			os.Exit(1)
		}

		filesPath := os.Getenv("HOME") + "/.k8config/configs/"

		files, err := ioutil.ReadDir(filesPath)

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

		// fmt.Println("Run this command: export KUBECONFIG=" + basePath + "/" + config)
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
