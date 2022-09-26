package cfg

import (
	"log"

	"github.com/spf13/viper"
)

//Cfg is the global config
var Cfg Config

//Init initiates the Cfg
func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("Unable to decode into Config, %s", err)
	}
}

//Config is the general config
type Config struct {
	App AppConfig `mapstructure:"app"`
}

//AppConfig is the app config
type AppConfig struct {
	ENV        string `mapstructure:"env"`
	ServerPort string `mapstructure:"server_port"`
}
