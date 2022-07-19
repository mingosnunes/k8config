/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mingosnunes/k8config/models"
	"github.com/mingosnunes/k8config/utils"

	"github.com/spf13/cobra"
)

// functions to mock for testing
var (
	getHomeDir     = os.UserHomeDir
	mkDirRoot      = os.Mkdir
	mkDirConfigs   = os.Mkdir
	createSettings = models.CreateSettings
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install k8config system files on your system",
	Long: `Install k8config system files on your system.
The system files will be localted on the $HOME/.k8config directory (created if not present).`,
	RunE: func(cmd *cobra.Command, args []string) error {

		utils.PrintInfo(fmt.Sprint("Checks: ", checks))

		if len(checks) == 0 {
			println("🤓 Hey Noob... You already have everything installed!")
			return nil
		}

		homePath, err := getHomeDir()

		if err != nil {
			return err
		}

		utils.PrintInfo(fmt.Sprintf("Home folder location: %s", homePath))

		for _, check := range checks {
			switch check {
			case 1:
				utils.PrintInfo("Installing...")

				err := mkDirRoot(homePath+"/.k8config", os.ModePerm)

				if err != nil {
					return err
				}

				err = createSettings()

				if err != nil {
					return err
				}

				err = mkDirConfigs(homePath+"/.k8config/configs", os.ModePerm)

				if err != nil {
					return err
				}

			case 2:
				utils.PrintInfo("Settings file not found. Installing...")

				err = createSettings()

				if err != nil {
					return err
				}
			case 3:
				utils.PrintInfo("Kubernetes configuration directory not found. Installing...")

				err := mkDirConfigs(homePath+"/.k8config/configs", os.ModePerm)

				if err != nil {
					return err
				}
			case 4:
				utils.PrintWaring("Add this line to your .profile/.bashrc/.zshrc file and source it:\n\texport KUBECONFIG=$HOME/.k8config/actual")
			}
		}

		utils.PrintSuccess("Installation Successful")

		return nil

	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
