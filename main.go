package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Harga int    `json:"price"`
	Stok  int    `json:"stock"`
}

var products = []product{
	{ID: 1, Name: "Laptop", Harga: 1500, Stok: 10},
	{ID: 2, Name: "Smartphone", Harga: 800, Stok: 25},
	{ID: 3, Name: "Tablet", Harga: 400, Stok: 15},
}

func getproductByID(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
	return
}

func updateProductByID(w http.ResponseWriter, r *http.Request) {
	//get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	//ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	//get dat dari request
	var updatedProduct product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	// loop produk, cari id, ganti sesuai data dari request
	for i, p := range products {
		if p.ID == id {
			products[i] = updatedProduct
			products[i].ID = id // Ensure ID remains unchanged
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}
	http.Error(w, "Product Not Found", http.StatusNotFound)
	return
}

func deleteProductByID(w http.ResponseWriter, r *http.Request) {
	//get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	//ganti id ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	//loop produk, cari id, hapus data
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Product Deleted",
			})
			return
		}
	}

	//bikin slice baru dengan data sebelum dan sesudah data yang dihapus

	http.Error(w, "Product Not Found", http.StatusNotFound)
	return
}

func main() {

	// GET api/products/{id}
	// PUT api/products/{id}
	// DELETE api/products/{id}
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getproductByID(w, r)
			return
		} else if r.Method == "PUT" {
			updateProductByID(w, r)
			return
		} else if r.Method == "DELETE" {
			deleteProductByID(w, r)
			return
		}
	})

	// GET /api/products
	// POST /api/products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
			return
		} else if r.Method == "POST" {
			//baca data dari requst
			var newProduct product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			//masukin data ke dalam variabel
			newProduct.ID = len(products) + 1
			products = append(products, newProduct)

			w.WriteHeader(http.StatusCreated) //201
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newProduct)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})

	// GET /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Print("Server running at 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
