package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type ConfigIPForwarding struct {
	Enabled bool   `mapstructure:"enabled"`
	IP      string `mapstructure:"ip"`
	Port    string `mapstructure:"port"`
}

type AppConfig struct {
	AppPort        int    `mapstructure:"app_port"`
	AppEnvironment string `mapstructure:"app_environment"`
	DbDriver       string `mapstructure:"db_driver"`
	DbAddress      string `mapstructure:"db_address"`
	DbPort         int    `mapstructure:"db_port"`
	DbUsername     string `mapstructure:"db_username"`
	DbPassword     string `mapstructure:"db_password"`
	DbName         string `mapstructure:"db_name"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	var finalConfig AppConfig

	defaultConfig.AppPort = 8080
	defaultConfig.AppEnvironment = ""
	defaultConfig.DbDriver = "mysql"
	defaultConfig.DbAddress = "localhost"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = "root"
	defaultConfig.DbPassword = "1"
	defaultConfig.DbName = "db_name"

	viper.AutomaticEnv()
	viper.SetEnvPrefix("your")
	viper.BindEnv("app_port")
	viper.BindEnv("app_environment")
	viper.BindEnv("db_driver")
	viper.BindEnv("db_address")
	viper.BindEnv("db_port")
	viper.BindEnv("db_username")
	viper.BindEnv("db_password")
	viper.BindEnv("db_name")
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig
}
