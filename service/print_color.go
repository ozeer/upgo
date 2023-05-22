package service

import (
	"fmt"

	"github.com/fatih/color"
)

// 打印酒红色文字
func PrintMagenta(text string) {
	magenta := color.New(color.FgMagenta).SprintFunc()
	fmt.Println(magenta(text))
}

// 打印黄色文字
func PrintYellow(text string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(yellow(text))
}
