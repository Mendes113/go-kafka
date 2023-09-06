package entity

import "github.com/google/uuid"

type Product struct {
	Id    string
	Name  string
	Price float64
}


type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
	FindById(id string) (*Product, error)
	Update(product *Product) error
	Delete(id string) error
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}




