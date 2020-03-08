package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func get(key string) string {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error while get [%v] config file: %s", key, err))
	}

	return viper.GetString(key)
}

func getMongoConnStr() string {
	return get("MONGODB_CONNSTR")
}

func getToken() string {
	return get("TOKEN")
}
