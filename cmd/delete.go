/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"k8config/models"
	"k8config/utils"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete kubernetes configuration file",
	Long: `Delete kubernetes configuration file.

Kubernetes configuratio file will be removed from your system`,
	Run: func(cmd *cobra.Command, args []string) {
		checks := utils.CheckInstallation()

		if len(checks) > 0 {
			utils.PrintRed.Println("\n⚠️ k8config is not installed correctly. Run ➡️ k8config install")
			os.Exit(1)
		}

		settings := models.GetSettings()

		options := []string{"all"}

		for _, c := range settings.ConfigList {
			options = append(options, c.Name)
		}

		configs2remove := ""
		prompt := &survey.MultiSelect{
			Message: "Choose a kube config:",
			Options: options,
		}

		survey.AskOne(prompt, &configs2remove)

		// for _, config := range configs2remove {

		// }

		// err := os.Remove()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
