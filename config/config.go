package config

import (
	"sync"

	"github.com/spf13/viper"
)

var defaultConfigFile = "./configs/config.yaml"

type Config struct {
	User    UserConfig
	OCRHost string
	Mysql   MysqlConfig
}
type UserConfig struct {
	Account  string
	Password string
}
type MysqlConfig struct {
	Host   string
	Port   string
	User   string
	Passwd string
	DBName string
}

var conf *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(defaultConfigFile)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&conf); err != nil {
			panic(err)
		}
	})
	return conf

}

func SetConfigPath(path string) {
	defaultConfigFile = path
}
