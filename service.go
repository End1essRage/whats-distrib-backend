package main

import (
	"github.com/sirupsen/logrus"
)

type Service interface {
	HandleScanRequest(fileName string) error
}

type TestService struct {
	//scanner Scanner
	client Client
}

func NewTestService(c Client) *TestService {
	return &TestService{client: c}
}

func (s *TestService) HandleScanRequest(fileName string) error {
	//настраиваем сканер
	cells := make(map[int]string)
	cells[0] = "Name"
	cells[1] = "PhoneNumber"

	scanner := NewExcelScanner(cells)

	logrus.Info("Handling in service")

	logrus.Info("Начало сканирования")
	scanner.Scan("uploaded/" + fileName)

	logrus.Info("получение результата")
	result := scanner.GetResult()

	logrus.Info("Сериализация результата")
	data := make([]Track, 0)
	for id, res := range result {
		tr := Track{ID: id}
		for i, val := range res {
			row := Data{Column: scanner.Columns[i], Value: val}
			tr.Data = append(tr.Data, row)
		}
		data = append(data, tr)
	}

	logrus.Info("Сохранение в файл")
	if err := SaveToJsonFile(fileName, data); err != nil {
		logrus.Error("error while saving file " + err.Error())
	}

	return nil
}
