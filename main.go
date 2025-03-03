package main

import (
	"go-web-native/config"
	categorycontroller "go-web-native/controllers/CategoryController"
	homecontroller "go-web-native/controllers/HomeController"
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

	log.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
