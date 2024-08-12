package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Scanner interface {
	Scan(filePath string) (*ScanResult, error)
}

type ExcelScanner struct {
	cells map[int]string
}

type ScanResult struct {
	Data map[int]string
}

func NewExcelScanner(c map[int]string) *ExcelScanner {
	return &ExcelScanner{cells: c}
}

func (s *ExcelScanner) Scan(filePath string) (*ScanResult, error) {
	result := make(map[int]string)
	result[0] = "0"

	//преобразовываем в
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	return &ScanResult{Data: result}, nil
}
