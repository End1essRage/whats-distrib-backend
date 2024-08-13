package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func GenerateFileName(fileName string) string {
	currentTime := time.Now().Format(time.RFC3339)

	data := fileName + currentTime

	hash := sha256.New()
	hash.Write([]byte(data))

	return hex.EncodeToString(hash.Sum(nil))
}

func SaveToJsonFile(fileName string, data any) error {
	if err := os.Mkdir("jsons", os.ModePerm); err != nil {
		logrus.Error("error while creating folder " + err.Error())
	}

	file, err := os.Create("jsons/" + strings.Split(fileName, ".")[0] + ".json")
	if err != nil {
		logrus.Error("error while creating file")
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Для удобочитаемого формата

	if err := encoder.Encode(data); err != nil {
		logrus.Error("Ошибка при записи в файл: ")
		return err
	}
	return nil
}
