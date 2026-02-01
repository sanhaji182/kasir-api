package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// type product struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Harga int    `json:"price"`
// 	Stok  int    `json:"stock"`
// }

// type category struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// var products = []product{
// 	{ID: 1, Name: "Laptop", Harga: 1500, Stok: 10},
// 	{ID: 2, Name: "Smartphone", Harga: 800, Stok: 25},
// 	{ID: 3, Name: "Tablet", Harga: 400, Stok: 15},
// }

// var categories = []category{
// 	{ID: 1, Name: "Electronics", Description: "Gadgets and devices"},
// 	{ID: 2, Name: "Books", Description: "Printed and digital books"},
// 	{ID: 3, Name: "Clothing", Description: "Apparel and accessories"},
// 	{ID: 4, Name: "Home & Kitchen", Description: "Household items and kitchenware"},
// }

// func getcategories(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(categories)
// 	// return
// }

// func postcategory(w http.ResponseWriter, r *http.Request) {
// 	var newCategory category
// 	err := json.NewDecoder(r.Body).Decode(&newCategory)
// 	if err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}

// 	newCategory.ID = len(categories) + 1
// 	categories = append(categories, newCategory)

// 	w.WriteHeader(http.StatusCreated) //201
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(newCategory)
// }

// func updatecategoryByID(w http.ResponseWriter, r *http.Request) {
// 	//get id dari request
// 	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
// 	//ganti int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
// 		return
// 	}
// 	//get dat dari request
// 	var updatedCategory category
// 	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
// 	if err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}
// 	// loop produk, cari id, ganti sesuai data dari request
// 	for i, c := range categories {
// 		if c.ID == id {
// 			categories[i] = updatedCategory
// 			categories[i].ID = id // Ensure ID remains unchanged
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(categories[i])
// 			return
// 		}
// 	}
// 	http.Error(w, "Category Not Found", http.StatusNotFound)
// 	return
// }

// func getcategoriesByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, c := range categories {
// 		if c.ID == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(c)
// 			return
// 		}
// 	}

// 	http.Error(w, "Category Not Found", http.StatusNotFound)
// 	return
// }

// func deletecategoryByID(w http.ResponseWriter, r *http.Request) {
// 	//get id dari request
// 	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
// 	//ganti id ke int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
// 		return
// 	}
// 	//loop produk, cari id, hapus data
// 	for i, c := range categories {
// 		if c.ID == id {
// 			categories = append(categories[:i], categories[i+1:]...)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "Category Deleted",
// 			})
// 			return
// 		}
// 	}

// 	http.Error(w, "Category Not Found", http.StatusNotFound)
// 	return
// }

// func getproductByID(w http.ResponseWriter, r *http.Request) {

// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, p := range products {
// 		if p.ID == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(p)
// 			return
// 		}
// 	}

// 	http.Error(w, "Product Not Found", http.StatusNotFound)
// 	return
// }

// func updateProductByID(w http.ResponseWriter, r *http.Request) {
// 	//get id dari request
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
// 	//ganti int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
// 		return
// 	}
// 	//get dat dari request
// 	var updatedProduct product
// 	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
// 	if err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}
// 	// loop produk, cari id, ganti sesuai data dari request
// 	for i, p := range products {
// 		if p.ID == id {
// 			products[i] = updatedProduct
// 			products[i].ID = id // Ensure ID remains unchanged
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(products[i])
// 			return
// 		}
// 	}
// 	http.Error(w, "Product Not Found", http.StatusNotFound)
// 	return
// }

// func deleteProductByID(w http.ResponseWriter, r *http.Request) {
// 	//get id dari request
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
// 	//ganti id ke int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
// 		return
// 	}
// 	//loop produk, cari id, hapus data
// 	for i, p := range products {
// 		if p.ID == id {
// 			products = append(products[:i], products[i+1:]...)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "Product Deleted",
// 			})
// 			return
// 		}
// 	}

// 	//bikin slice baru dengan data sebelum dan sesudah data yang dihapus

// 	http.Error(w, "Product Not Found", http.StatusNotFound)
// 	return
// }

type Config struct {
	Port    string `mapstructure:"port"`
	DB_CONN string `mapstructure:"db_conn"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
		if err != nil {
			fmt.Println("Error reading config file", err)
			return
		}
	}

	config := Config{
		Port:    viper.GetString("port"),
		DB_CONN: viper.GetString("db_conn"),
	}

	db, err := database.InitDB(config.DB_CONN)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)
	// http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		if r.URL.Path == "/categories/" {
	// 			getcategories(w, r)
	// 			return
	// 		} else {
	// 			getcategoriesByID(w, r)
	// 			return
	// 		}
	// 		// getcategories(w, r)
	// 		// return
	// 	} else if r.Method == "POST" {
	// 		postcategory(w, r)
	// 		return
	// 	} else if r.Method == "PUT" {
	// 		updatecategoryByID(w, r)
	// 		return
	// 	} else if r.Method == "DELETE" {
	// 		deletecategoryByID(w, r)
	// 		return
	// 	}
	// 	//  else if r.Method == "GET" {
	// 	// 	getcategoriesByID(w, r)
	// 	// 	return
	// 	// }
	// })

	// GET api/products/{id}
	// PUT api/products/{id}
	// DELETE api/products/{id}
	// http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		getproductByID(w, r)
	// 		return
	// 	} else if r.Method == "PUT" {
	// 		updateProductByID(w, r)
	// 		return
	// 	} else if r.Method == "DELETE" {
	// 		deleteProductByID(w, r)
	// 		return
	// 	}
	// })

	// GET /api/products
	// POST /api/products
	// http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(products)
	// 		return
	// 	} else if r.Method == "POST" {
	// 		//baca data dari requst
	// 		var newProduct product
	// 		err := json.NewDecoder(r.Body).Decode(&newProduct)
	// 		if err != nil {
	// 			http.Error(w, "Invalid Request", http.StatusBadRequest)
	// 			return
	// 		}

	// 		//masukin data ke dalam variabel
	// 		newProduct.ID = len(products) + 1
	// 		products = append(products, newProduct)

	// 		w.WriteHeader(http.StatusCreated) //201
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(newProduct)
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(products)
	// })

	// GET /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Print("Server running at " + config.Port + "\n")
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
