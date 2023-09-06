package main

import (
	"database/sql"
	"encoding/json"
	"mendes/kafka/internal/infra/akafka"
	"mendes/kafka/internal/infra/repository"
	"mendes/kafka/internal/infra/web"
	"mendes/kafka/internal/usecase"
	"net/http"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
    
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host/docker/internal:3036/products)")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	

	productRepository := repository.NewProductRepositoryMySql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(productRepository)
	listProductUseCase := usecase.NewCreateProductUseCase(productRepository)

	productHandler := web.NewProductoHandler(createProductUseCase, (*usecase.ListProductUseCases)(listProductUseCase) )

	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProductHandler)
	r.Get("/products", productHandler.ListProductsHandler)

	go http.ListenAndServe(":8080", r)
	msgChan := make(chan *kafka.Message)
	
	go akafka.Consume([]string{"products"}, "host:9092", msgChan)

	for msg := range msgChan {
	 		dto := usecase.CreateProductInputDTO{}
			err := json.Unmarshal(msg.Value, &dto)
			if err != nil {
				
			}
			_, err = createProductUseCase.Execute(dto)

	}


	
}