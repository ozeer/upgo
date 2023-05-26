/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ozeer/upgo/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initial installation of UpGo (if Golang is not installed locally)",
	Long:  `初始化安装UpGo（本地未安装Golang的情况下）`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		yellow := color.New(color.FgYellow).SprintFunc()

		defaultInstallDir := viper.GetString("app.install_dir")

		for {
			fmt.Print(yellow(fmt.Sprintf("请输入你想将UpGo安装到的目录（默认：%s）: ", defaultInstallDir)))

			// 获取用户输入的目录路径
			// fmt.Scanln(&installDir)
			installDir, _ := reader.ReadString('\n')
			installDir = strings.TrimSpace(installDir)

			if installDir == "" {
				installDir = defaultInstallDir
			}

			if service.CheckInputDirIsValid(installDir) {
				magenta := color.New(color.FgMagenta).SprintFunc()

				for {
					fmt.Print(yellow(fmt.Sprintf("请确认安装目录%s？(y/n)：", magenta(installDir))))
					yes, _ := reader.ReadString('\n')
					yes = strings.TrimSpace(yes)

					if yes == "y" {
						service.InitUpGo(installDir)
					}
					break
				}
				break
			} else {
				fmt.Print(yellow("输入不正确! "))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
