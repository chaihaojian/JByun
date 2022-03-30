package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.Redconfig() faild, err:%v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event){
		fmt.Printf("Configuration file is modified...")
	})
	return
}
