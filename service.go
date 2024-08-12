package main

import "github.com/sirupsen/logrus"

type Service interface {
	HandleRequest()
}

type TestService struct {
	scanner Scanner
	client  Client
}

func NewTestService(s Scanner, c Client) *TestService {
	return &TestService{scanner: s, client: c}
}

func (s *TestService) HandleRequest() {
	logrus.Info("Handling in service")

}
