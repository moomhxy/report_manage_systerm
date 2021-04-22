package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Port  string `yaml:"port"`
	Mysql Mysql  `yaml:"mysql"`
}

type Mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

func InitConfig() error {
	viper.SetConfigName("config") //配置文件名称，不需要带后缀
	viper.AddConfigPath("conf")   //配置文件路径
	viper.SetConfigType("yaml")   //配置文件类型
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read config error: %+v\n", err)
		return err
	}

	// 读取配置文件
	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Printf("unmarshal conf error: %+v\n", err)
		return err
	}
	fmt.Printf("%+v", Conf)
	return nil
}
