package product

import (
	"encoding/json"
	"my-project-mux/main/internal/model"
	"my-project-mux/main/internal/repository"
	"net/http"
)

type (
	RequestProduct struct {
		Name   string  `json:"name"`
		Price  float64 `json:"price"`
		Amount int     `json:"amount"`
	}

	ResponseProduct struct{
		Status string `json:"status"`
		ProductList []model.Product `json:"product_list"`
	}
)

func GetProduct(w http.ResponseWriter, r *http.Request){

	statusMessage := "success"
	productList, err := repository.DB().ListProduct()
	if err != nil{
		statusMessage = "failed :" + err.Error()
	}
	//write response
	res := 	ResponseProduct{
		Status: statusMessage,
		ProductList: productList,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
