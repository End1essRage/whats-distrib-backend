package main

import (
	"fmt"
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

	// Преобразуем
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer func() {
		// Закрываем таблицу
		if err := file.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	rows, err := file.GetRows(viper.GetString("sheet_name"))
	if err != nil {
		logrus.Error("can't get rows")
		return err
	}

	// Создаем хранение значений ячеек с указанием их адресов
	filledCells := make(map[string]string)

	for id, row := range rows {
		if id > 0 { // Начинаем со второй строки

			sb := strings.Builder{}
			for i := 0; i < len(s.Columns); i++ {
				cellRef := fmt.Sprintf("%s%d", string('A'+i), id+1) // Определяем ссылку на ячейку

				// Проверяем, есть ли значение в строке или нет
				var cellValue string
				if i < len(row) {
					cellValue = row[i]
				}

				// Заполняем мапу значением (пустая ячейка сохраняется как пустая строка)
				filledCells[cellRef] = cellValue

				// Добавляем значение в строку
				if cellValue != "" {
					sb.WriteString(cellValue + "|")
				} else {
					sb.WriteString("EMPTY|") // Если ячейка пустая
				}
			}

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
