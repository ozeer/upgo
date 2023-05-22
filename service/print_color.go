package service

import (
	"fmt"
	"strings"

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

// 按每列固定显示n个版本
func PrintFixedColumnVersion(data []Version) {
	// 每列的宽度
	columnWidth := 20
	// 每行的元素个数
	columnCount := 5
	// 总行数
	rowCount := (len(data) + columnCount - 1) / columnCount

	cyan := color.New(color.FgCyan).SprintFunc()

	for i := 0; i < rowCount; i++ {
		startIndex := i * columnCount
		endIndex := (i + 1) * columnCount
		if endIndex > len(data) {
			endIndex = len(data)
		}

		// 每行的元素
		rowElements := data[startIndex:endIndex]
		rowString := ""

		for _, element := range rowElements {
			padding := strings.Repeat(" ", columnWidth-len(element.Version))
			rowString += element.Version + padding
		}

		fmt.Println(cyan(rowString))
	}
}
