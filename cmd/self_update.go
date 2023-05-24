/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ozeer/upgo/service"
	"github.com/spf13/cobra"
)

// selfUpdateCmd represents the selfUpdate command
var selfUpdateCmd = &cobra.Command{
	Use:     "self-update",
	Short:   "Updates UpGo to the latest version",
	Long:    `升级UpGo到最新版本`,
	Aliases: []string{"selfupdate", "s"},
	Run: func(cmd *cobra.Command, args []string) {
		service.SelfUpdate()
	},
}

func init() {
	rootCmd.AddCommand(selfUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selfUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selfUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
