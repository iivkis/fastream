package config

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

var Env struct {
	APP_HOST string `mapstructure:"APP_HOST"`
	API_PORT string `mapstructure:"API_PORT"`
	SPA_PORT string `mapstructure:"SPA_PORT"`
}

func LoadEnv(filename string) {
	once.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigFile(filename)
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&Env); err != nil {
			panic(err)
		}
	})
}
