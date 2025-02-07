package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func NewConfig(path string) *viper.Viper {
	envConf := path
	if envConf == "" {
		envConf = "config/dev.yml"
	}
	fmt.Println("load conf file:", envConf)
	return getConfig(envConf)

}
func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 监听配置变化
	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		_ = conf.ReadInConfig()
	})

	return conf
}
