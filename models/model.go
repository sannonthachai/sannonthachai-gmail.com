package models

type Model struct {
	Code  string `json:"code"`
	Price int    `json:"price"`
}

type DeleteModel struct {
	Code string `json:"code"`
}
