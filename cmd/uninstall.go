/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"upgo/service"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall golang",
	Long:  `卸载Golang程序`,
	Run: func(cmd *cobra.Command, args []string) {
		goInstallDir := service.GetGoRootDir()
		red := color.New(color.FgRed).SprintFunc()

		if goInstallDir == "" {
			fmt.Println(red("The installed golang is not found, please check whether there are environment variables such as goroot configured."))
		} else {
			version := service.GetCurrentGoVersion()

			if service.Command("sudo rm -rf " + goInstallDir) {
				green := color.New(color.FgGreen).SprintFunc()
				fmt.Println(green("==> Uninstall succeeded: " + version))
			} else {
				fmt.Println(red("==> Uninstall fail!"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
