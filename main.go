package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		//non fatal
		logrus.Error("error while reading environment %s", err.Error())
	}

	//настраиваем сканер

	cells := make(map[int]string)
	cells[0] = "Name"
	cells[1] = "PhoneNumber"

	scanner := NewExcelScanner(cells)
	client := NewWClient()
	service := NewTestService(scanner, client)
	api := NewApi(service)

	api.gin.Run(viper.GetString("host"))

	logrus.Info("app started!")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
