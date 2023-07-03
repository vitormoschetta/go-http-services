package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router, repository *Repository) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := repository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		product, err := repository.FindById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if product.ID == "" {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(product)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product = NewProduct(product.Name, product.Price)
		err := repository.Save(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}).Methods(http.MethodPost)

	router.HandleFunc("/api/v1/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		product, err := repository.FindById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if product.ID == "" {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newProduct.ID = product.ID
		err = repository.Update(newProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPut)

	router.HandleFunc("/api/v1/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		product, err := repository.FindById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if product.ID == "" {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		if err := repository.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodDelete)
}
