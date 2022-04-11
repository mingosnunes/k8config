/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// actualCmd represents the actual command
var actualCmd = &cobra.Command{
	Use:   "actual",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Actual kube config: " + os.Getenv("KUBECONFIG"))
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
