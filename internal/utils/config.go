package utils

import (
	"log"

	"github.com/spf13/viper"
)

func Getenv(key string) string {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		return ""
	}

	return value
}

func Setenv(key string, value string) {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	viper.Set(key, value)

	err = viper.WriteConfig()
	if err != nil {
		log.Fatalf("Error while writing config file %s", err)
	}
}
