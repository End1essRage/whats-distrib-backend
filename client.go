package main

type Client interface {
	SendMessage(number string, message string)
}

type WClient struct {
}

func NewWClient() *WClient {
	return &WClient{}
}

func (c *WClient) SendMessage(number string, message string) {

}
