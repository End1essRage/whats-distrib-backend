package main

// представляет одну строку значений
type Track struct {
	ID   int    `json:"id"`
	Data []Data `json:"data"`
}

type Data struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}
