/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package cmd

import (
	"github.com/mingosnunes/k8config/models"
	"github.com/mingosnunes/k8config/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	useGetSettings  = models.GetSettings
	useSurveyAskOne = survey.AskOne
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Change the active Kubernetes configuration file",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		settings, err := useGetSettings()

		if err != nil {
			return err
		}

		options := make([]string, 0)

		for _, c := range settings.GetConfigList() {
			options = append(options, c.Name)
		}

		config := ""
		prompt := &survey.Select{
			Message: "Choose a kube config:",
			Options: options,
		}

		err = useSurveyAskOne(prompt, &config)

		if err != nil {
			return err
		}

		err = settings.UseConfig(config)

		if err != nil {
			return err
		}

		utils.PrintSuccess("All done!")

		return nil
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
