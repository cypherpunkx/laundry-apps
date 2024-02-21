package model

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Uom string `json:"uom"`
}