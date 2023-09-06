package repository

import (
	"database/sql"
	"mendes/kafka/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

// findById implements entity.ProductRepository.
func (*ProductRepository) FindById(id string) (*entity.Product, error) {
	panic("unimplemented")
}

// update implements entity.ProductRepository.
func (*ProductRepository) Update(product *entity.Product) error {
	panic("unimplemented")
}

func NewProductRepositoryMySql(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.Id, product.Name, product.Price)

	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() { // Itera sobre as linhas retornadas
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}


func (r *ProductRepository) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

