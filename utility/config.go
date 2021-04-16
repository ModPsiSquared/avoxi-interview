package utility

import (
	"github.com/spf13/viper"
)

type Config struct {
	ListeningAddress string `mapstructure:"LISTENING_ADDRESS" default:"127.0.0.1:8080"`
}

func LoadConfig(path string) (config Config, err error) {

	if path == "" {
		path = GetConfigFolder()
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
