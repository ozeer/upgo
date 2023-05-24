/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ozeer/upgo/service"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install the specified version of Golang",
	Long:    `安装指定版本Golang.`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		yellow := color.New(color.FgYellow).SprintFunc()

		for {
			fmt.Print(yellow("请输入你想安装的Golang版本号(如：1.20.4): "))
			version, _ := reader.ReadString('\n')
			version = strings.TrimSpace(version)

			if service.IsValidVersion(version) {
				// go1.20.4.darwin-amd64.tar.gz
				fileName := "go" + version + ".darwin-amd64.tar.gz"
				// https://go.dev/dl/go1.20.4.darwin-amd64.tar.gz
				fileUrl := service.GO_DOWNLOAD_WEB + fileName
				service.Install(fileUrl, fileName)
				break
			} else {
				fmt.Print(yellow("输入不正确! "))
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
