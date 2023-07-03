package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router, service *Service) {

	router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		products, err := service.FindAll(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)
	}).Methods(http.MethodGet)
}
