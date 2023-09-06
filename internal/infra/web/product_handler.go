package web

import (
	"encoding/json"
	"mendes/kafka/internal/usecase"
	"net/http"
)

type ProductHandler struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductUseCase   *usecase.ListProductUseCases
}


func NewProductoHandler(createProductUseCase *usecase.CreateProductUseCase, listProductUseCase *usecase.ListProductUseCases) *ProductHandler {
	return &ProductHandler{
		CreateProductUseCase: createProductUseCase,
		ListProductUseCase:   listProductUseCase,
	}
}


func (p *ProductHandler) CreateProductHandler(w http.ResponseWriter,r *http.Request){
	var input usecase.CreateProductInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(output)

}


func (p *ProductHandler) ListProductsHandler(w http.ResponseWriter,r *http.Request){
	output, err := p.ListProductUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}