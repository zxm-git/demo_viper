package vip

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Conf struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
	Student `mapstructure:"student"`
}
type Student struct {
	Score int
	Name string
}

var C Conf

func Init() {
	//viper.SetDefault("con", "const")
	str, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(str)
	er := viper.ReadInConfig()
	if er != nil {
		fmt.Println(er)
	}
	//var c Conf
	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Println(err)
	}

}
