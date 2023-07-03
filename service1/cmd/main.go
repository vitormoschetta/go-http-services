package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitormoschetta/go-minimal-api/internal"
)

func main() {
	router := mux.NewRouter()
	router.Use(internal.AcceptJSON)

	repository := internal.NewRepository()
	internal.CreateRoutes(router, repository)
	seed(repository)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func seed(repository *internal.Repository) {
	repository.Save(internal.NewProduct("Product 1", 10.0))
	repository.Save(internal.NewProduct("Product 2", 20.0))
	repository.Save(internal.NewProduct("Product 3", 30.0))
}
