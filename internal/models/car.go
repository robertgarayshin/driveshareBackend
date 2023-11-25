package models

type Car struct {
	Id       int    `json:"id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Produced string `json:"produced"`
	Price    int    `json:"price"`
	Booked   bool   `json:"booked"`
}
