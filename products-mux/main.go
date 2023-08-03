package main

import (
	"fmt"
	"log"
	"net/http"
	"products-mux/controllers"
	"products-mux/database"
	"products-mux/initializers"

	"github.com/gorilla/mux"
)

func main() {
	// Load Configurations from config.json using Viper
	initializers.LoadAppConfig()

	// Initialize Database
	database.Connect((initializers.AppConfig).ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Printf(fmt.Sprintf("Starting Server on port %s", (initializers.AppConfig).Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", (initializers.AppConfig).Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
