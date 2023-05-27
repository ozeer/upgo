package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ozeer/upgo/service"
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
	service.Command("ls")

	if err != nil {
		panic(fmt.Sprintf("Load config file error: %s", err.Error()))
	}
}
