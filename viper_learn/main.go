package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	var configFile string

	// 设置命令行参数
	flag.StringVar(&configFile, "config", "config", "Configuration file in TOML format.")
	flag.Parse()

	// 初始化 viper
	conf := viper.New()
	conf.AddConfigPath(".")
	conf.SetConfigName(configFile)
	conf.AutomaticEnv()

	// 读取配置
	if err := conf.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// 从配置文件获取一个值
	host := conf.GetString("database.host")
	fmt.Printf("The host is : %s", host)

	// 另一种获取值的方式
	params := conf.Get("database").(map[string]interface{})
	password, exists := params["password"]
	if exists {
		fmt.Println("The password is : ", password.(string))
	}
}
