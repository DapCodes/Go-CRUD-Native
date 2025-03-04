package main

import (
	"go-web-native/config"
	categorycontroller "go-web-native/controllers/CategoryController"
	homecontroller "go-web-native/controllers/HomeController"
	productcontroller "go-web-native/controllers/ProductController"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
