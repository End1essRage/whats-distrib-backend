package main

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
)

type Scanner interface {
	Scan(filePath string) error
	//строки таблицы и перечень значений
	GetResult() [][]string
}

type ExcelScanner struct {
	// номер колонки - название свойства
	Columns map[int]string
	//каждый элемент предсавляет свою колонку, в каждой колонке набор значений
	result []string
}

func NewExcelScanner(c map[int]string) *ExcelScanner {
	return &ExcelScanner{Columns: c}
}

func (s *ExcelScanner) Scan(filePath string) error {
	result := make([]string, 0)

	//преобразовываем
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	rows, err := file.GetRows(viper.GetString("sheet_name"))
	if err != nil {
		logrus.Error("cant get rows")
		return err
	}

	for id, row := range rows {
		if id > 0 {
			sb := strings.Builder{}
			for i, _ := range s.Columns {
				sb.WriteString(row[i] + "|")
			}
			logrus.Info(sb.String())
			result = append(result, sb.String())
		}
	}

	s.result = result

	return nil
}

func (s *ExcelScanner) GetResult() [][]string {
	result := make([][]string, 0)

	//Проходимся по массиву результатов считывания строки
	for _, v := range s.result {
		rowData := make([]string, 0)
		//режем на ячейки и проходимся по нима
		data := strings.Split(v, "|")

		for _, val := range data {
			if val != "" {
				rowData = append(rowData, val)
			}
		}

		result = append(result, rowData)
	}

	return result
}
