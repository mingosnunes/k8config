/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package cmd

import (
	"os"

	"github.com/mingosnunes/k8config/models"
	"github.com/mingosnunes/k8config/utils"
	"github.com/spf13/cobra"
)

// actualCmd represents the actual command
var actualCmd = &cobra.Command{
	Use:   "actual",
	Short: "Show the active Kubernetes configuration file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		checks := utils.CheckInstallation()

		if len(checks) > 0 {
			utils.PrintRed.Println("\n⚠️ k8config is not installed correctly. Run ➡️ k8config install")
			os.Exit(1)
		}

		settings := models.GetSettings()

		if settings.CurrentConfig.Name == "" {
			if len(settings.ConfigList) > 0 {
				utils.PrintWaring("Wake up Noob 🙄 You don't have any active config! Run `k8config use`")
			} else {
				utils.PrintWaring("Wake up Noob 🙄 You don't have any configs saved! Run `k8config add <config-path>`")
			}

		} else {
			utils.PrintInfo("Actual kube config: " + settings.CurrentConfig.Name + " (" + settings.CurrentConfig.Location + ")")
		}
	},
}

func init() {
	rootCmd.AddCommand(actualCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actualCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actualCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
