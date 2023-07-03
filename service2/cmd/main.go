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

	service := internal.NewService()
	internal.CreateRoutes(router, service)

	fmt.Println("Server running on port 8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}
}
