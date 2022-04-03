/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"k8config/utils"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install k8config system files on your system",
	Long: `Install k8config system files on your system.
The system files will be localted on the $HOME/.k8config directory (created if not present).`,
	Run: func(cmd *cobra.Command, args []string) {
		printRed := color.New(color.FgHiRed)
		// printGreen := color.New(color.FgHiGreen)
		printYellow := color.New(color.FgHiYellow)

		homePath := os.Getenv("HOME")

		dirInfo, err := os.Stat(homePath + "/.k8config")

		if err != nil {
			printYellow.Println("$HOME/.k8config not found. Installing...")

			os.Mkdir(homePath+"/.k8config", os.ModePerm)

			utils.CreateSettings(homePath)

			printYellow.Println("Installation Successful")

		} else {
			_, errSettings := os.Stat(homePath + "/.k8config/settings.json")

			if dirInfo.IsDir() && errSettings == nil {
				printRed.Println("Hey Noob. You already have everything installed!")
			} else if errSettings != nil {
				println("Settings file not found. Installing...")

				utils.CreateSettings(homePath)

				printYellow.Println("Installation Successful")
			}

		}

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
