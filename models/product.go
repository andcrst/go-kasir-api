package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	Price int    `json:"harga"`
	Stock int    `json:"stok"`
}