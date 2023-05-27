package conf

import (
	"fmt"
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

	if err != nil {
		panic(fmt.Sprintf("Load config file error: %s", err.Error()))
	}
}
