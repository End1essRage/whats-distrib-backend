package main

import (
	"github.com/end1essrage/xslxmerge"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	logrus.Info("Handling in service")

	f := xslxmerge.NewReadFascade("uploaded/"+fileName, viper.GetString("sheet_name"))
	command, err := f.NewReadFull()
	if err != nil {
		logrus.Error(err)
	}

	rows, err := command.ReadRowsSync()
	if err != nil {
		logrus.Error(err)
	}

	columns, err := xslxmerge.GetAllHeaders("uploaded/"+fileName, viper.GetString("sheet_name"))
	if err != nil {
		logrus.Error(err)
	}

	data := make([]Track, 0)

	for _, row := range rows {
		tr := Track{ID: row.Id}

		for _, cell := range row.Data {
			rd := Data{Column: columns[cell.ColumnId].Data, Value: cell.Data}
			tr.Data = append(tr.Data, rd)
		}

		data = append(data, tr)
	}

	logrus.Info("Сохранение в файл")
	if err := SaveToJsonFile(fileName, data); err != nil {
		logrus.Error("error while saving file " + err.Error())
	}

	return nil
}
