package models

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Rating int     `json:"rating"`
	Price  float64 `json:"price"`
}
