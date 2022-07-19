/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package cmd

import (
	"github.com/mingosnunes/k8config/models"
	"github.com/mingosnunes/k8config/utils"
	"github.com/spf13/cobra"
)

var (
	actualGetSettings = models.GetSettings
)

// actualCmd represents the actual command
var actualCmd = &cobra.Command{
	Use:   "actual",
	Short: "Show the active Kubernetes configuration file",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		settings, err := actualGetSettings()

		if err != nil {
			return err
		}

		currentConfig := settings.GetCurrentConfig()

		if currentConfig.Name == "" {
			if len(settings.GetConfigList()) > 0 {
				utils.PrintWaring("Wake up Noob 🙄 You don't have any active config! Run `k8config use`")
			} else {
				utils.PrintWaring("Wake up Noob 🙄 You don't have any configs saved! Run `k8config add <config-path>`")
			}

		} else {
			utils.PrintInfo("Actual kube config: " + currentConfig.Name + " (" + currentConfig.Location + ")")
		}

		return err
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
