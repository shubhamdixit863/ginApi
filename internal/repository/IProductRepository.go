package repository

import "ginTraining/internal/entity"

type IProductRepository interface {
	AddProduct(product entity.Product) (*entity.Product, error)
	GetProductById(id int) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
}
