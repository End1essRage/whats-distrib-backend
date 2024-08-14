package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			logrus.Fatalf("error while reading environment %s", err.Error())
		}

	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config %s", err.Error())
	}

	client := NewWClient()
	service := NewTestService(client)
	api := NewApi(service)

	api.gin.Run(viper.GetString("host"))

	logrus.Info("app started!")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	//базовый
	viper.SetConfigName("config")
	viper.ReadInConfig()
	// из окружения
	viper.SetConfigName("config." + os.Getenv("ENV"))
	return viper.MergeInConfig()
}
