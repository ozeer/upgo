/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ozeer/upgo/service"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show UpGo version",
	Long:  `All software has versions. This is UpGo's`,
	Run: func(cmd *cobra.Command, args []string) {
		goVersion := service.GetCurrentGoVersion()
		upGoVersion := viper.GetString("app.version")
		color.Cyan(fmt.Sprintf("UpGo version %s （Golang: %s）", upGoVersion, goVersion))
	},
	Aliases: []string{"v", "V"},
	Example: "upgo v",
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
