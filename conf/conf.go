package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Load config file error: %s", err.Error()))
	}
}
