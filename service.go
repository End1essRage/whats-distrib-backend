package main

import "github.com/sirupsen/logrus"

type Service interface {
	HandleScanRequest(fileName string) error
}

type TestService struct {
	scanner Scanner
	client  Client
}

func NewTestService(s Scanner, c Client) *TestService {
	return &TestService{scanner: s, client: c}
}

func (s *TestService) HandleScanRequest(fileName string) error {
	logrus.Info("Handling in service")

	if err := s.scanner.Scan("uploaded/" + fileName); err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info("Scan Data is ")
	logrus.Info(s.scanner.GetResult())

	return nil
}
