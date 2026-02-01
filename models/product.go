package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Harga int    `json:"price"`
	Stok  int    `json:"stock"`
}
