package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func InitConfig() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()

	dir, errWd := os.Getwd()
	if errWd != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}

	fmt.Println("当前所在目录:", dir)
	fmt.Println("InitConfig dir: ", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("读取目录失败:", err)
		return
	}

	fmt.Println("当前目录下的文件列表:")
	for _, file := range files {
		if file.IsDir() {
			// 如果需要只打印文件而不包括目录，可以在此处添加过滤逻辑
			fmt.Println(file.Name(), "(目录)")
		} else {
			fmt.Println(file.Name())
		}
	}

	if err != nil {
		panic(fmt.Sprintf("Load config file error: %s", err.Error()))
	}
}
