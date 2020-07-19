package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func get(key string) string {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error while get [%v] config file: %s", key, err))
	}

	return viper.GetString(key)
}

func GetMongoConnStr() string {
	return get("MONGODB_CONNSTR")
}

func GetToken() string {
	if(os.Getenv("K8S") == "1") {
		return os.Getenv("TOKEN")
	}
	return get("TOKEN")
}

func GetRabbitDNS() string {
	if(os.Getenv("K8S") == "1") {
		return os.Getenv("RABBITMQ_DNS")
	}
	return get("RABBITMQ_DNS")
}

func GetESURL() string {
	return get("ELASTICSEARCH_URL")
}

func GetESIndex() string {
	return get("ELASTICSEARCH_INDEX")
}

func GetLogHook() string {
	return get("ELASTICSEARCH_HOOK")
}

func GetLogPath() string {
	return get("LOG_PATH")
}
