/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/mingosnunes/k8config/models"
	"github.com/mingosnunes/k8config/utils"

	"github.com/spf13/cobra"
)

var (
	listGetSettings = models.GetSettings
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available kubernetes configuration files",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		settings, err := listGetSettings()

		if err != nil {
			return err
		}

		utils.PrintInfo("Available Kubernetes configs:")

		for _, config := range settings.GetConfigList() {
			fmt.Println("\t" + config.Name + " (" + config.Location + ")")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
