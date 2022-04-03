/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add kubernetes configuration file",
	Long: `Add kubernetes configuration file.

Kubernetes configuratio file will be added to the list of available configs`,
	Run: func(cmd *cobra.Command, args []string) {
		//check if config already exists

		//copy config to app's dir
		appPath := os.Getenv("HOME") + "/.k8config/"

		srcSplit := strings.Split(args[0], "/")

		src := args[0]
		dest := appPath + srcSplit[len(srcSplit)-1]

		bytesRead, err := os.ReadFile(src)

		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(dest, bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
		}

		println("😈 This configuration file is mine now!")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
