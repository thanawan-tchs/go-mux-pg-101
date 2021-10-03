package service

import (
	"my-project-mux/main/internal/service/product"

	"github.com/gorilla/mux"
)

func InitRoute(r *mux.Router){
	r.HandleFunc("/api/products", product.GetProduct).Methods("GET")
	r.HandleFunc("/api/products", product.PostProduct).Methods("POST")
}