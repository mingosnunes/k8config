/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"k8config/models"
	"k8config/utils"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add kubernetes configuration file",
	Long: `Add kubernetes configuration file.

Kubernetes configuratio file will be added to the list of available configs`,
	Run: func(cmd *cobra.Command, args []string) {
		checks := utils.CheckInstallation()

		if len(checks) > 0 {
			utils.PrintRed.Println("\n⚠️ k8config is not installed correctly. Run ➡️ k8config install")
			os.Exit(1)
		}

		if len(args) != 1 {
			utils.PrintWaring("Add the config path as argument: k8config add <path>")
			os.Exit(1)
		}

		//copy config to configs' dir

		srcSplit := strings.Split(args[0], "/")
		fileName := srcSplit[len(srcSplit)-1]

		src := args[0]
		// tempDest := utils.ConfigsPath + fileName

		bytesRead, err := os.ReadFile(src)

		if err != nil {
			log.Fatal(err)
		}

		settings := models.GetSettings()

		fileName2Save := ""
		dest := ""
		for {
			prompt := &survey.Input{
				Message: "Name to save:",
				// Suggest: func (toComplete string) []string {
				// 	files, _ := filepath.Glob(toComplete + "*")
				// 	return files
				// },
				Default: fileName,
			}
			err := survey.AskOne(prompt, &fileName2Save)

			if err != nil {

				switch err.Error() {
				case "interrupt":
					utils.PrintWaring("Cancel by Mr. Noob")
					os.Exit(0)
				default:
					log.Fatal(err)
				}

			}

			dest = utils.ConfigsPath + "/" + fileName2Save

			checkName := settings.CheckConfigName(fileName2Save)

			if checkName {
				// create config copy
				err = os.WriteFile(dest, bytesRead, 0644)
				if err != nil {
					log.Fatal(err)
				}

				break
			}

			utils.PrintError("Noob, pay attention! We already have that config name 🙄\n")

			fileName2Save = ""

		}

		config := models.NewK8sConfig(fileName2Save, dest)

		settings.AddConfig(config)

		utils.PrintSuccess(" This configuration file is mine now! 😎")

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
