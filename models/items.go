package models

type Item struct {
	Id      int     `json: "id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Creator User    `json: "creator"`
}

var Items []Item
