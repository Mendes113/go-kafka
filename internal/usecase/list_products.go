package usecase

import "mendes/kafka/internal/entity"

type ListProductOutputDTO struct {
	Id    string
	Name  string
	Price float64
}

type ListProductUseCases struct {
	ProductRepository entity.ProductRepository
}


func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase{
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func(u *ListProductUseCases) Execute() ([]*ListProductOutputDTO, error){
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var productsOutput []*ListProductOutputDTO
	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductOutputDTO{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}