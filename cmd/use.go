/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/mingosnunes/k8config/models"
	"github.com/mingosnunes/k8config/utils"

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

		settings := models.GetSettings()

		options := make([]string, 0)

		for _, c := range settings.ConfigList {
			options = append(options, c.Name)
		}

		config := ""
		prompt := &survey.Select{
			Message: "Choose a kube config:",
			Options: options,
		}

		survey.AskOne(prompt, &config)

		settings.UseConfig(config)

		utils.PrintSuccess("All done!")
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
