/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/mingosnunes/k8config/models"
	"github.com/mingosnunes/k8config/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete kubernetes configuration file",
	Long: `Delete kubernetes configuration file.

Kubernetes configuratio file will be removed from your system`,
	RunE: func(cmd *cobra.Command, args []string) error {

		settings, err := models.GetSettings()

		if err != nil {
			return err
		}

		options := make([]string, 0)

		for _, c := range settings.GetConfigList() {
			options = append(options, c.Name)
		}

		configs2remove := make([]string, 0)
		prompt := &survey.MultiSelect{
			Message: "Choose a kube config:",
			Options: options,
		}

		err = survey.AskOne(prompt, &configs2remove)

		if err != nil {
			return nil
		}

		if len(configs2remove) == 0 {
			fmt.Println()
			utils.PrintWaring("No config selected... 🙄")
			return nil
		}

		settings.DelConfigs(configs2remove)

		utils.PrintSuccess("All configs removed")

		return nil
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
