package product

import (
	"encoding/json"
	"my-project-mux/main/internal/model"
	"my-project-mux/main/internal/repository"
	"net/http"
	"time"
)

type (
	postRequestProduct struct {
		Name   string  `json:"name"`
		Price  float64 `json:"price"`
		Amount int     `json:"amount"`
	}

	postResponseProduct struct{
		Status string `json:"status"`
		ProductList []model.Product `json:"product_list"`
	}
)

func PostProduct(w http.ResponseWriter, r *http.Request){
	//parse request
	var rqProduct postRequestProduct
	statusMessage := "success"
	err := json.NewDecoder(r.Body).Decode(&rqProduct)
	if err != nil{
		statusMessage = "failed :" + err.Error()
	}
	//service logic
	modelProduct := model.NewProduct{
		Name: rqProduct.Name,
		Price: rqProduct.Price,
		Amount: rqProduct.Amount,
	}
	newProduct, err  := repository.DB().CreateNewProduct(modelProduct, time.Now())
	if err != nil{
		statusMessage = "failed :" + err.Error()
	}

	//write response
	res := 	postResponseProduct{
		Status: statusMessage,
		ProductList: []model.Product{newProduct},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}