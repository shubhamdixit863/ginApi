package services

import "ginTraining/internal/repository"

type ProductService struct {
	ProductRepo repository.IProductRepository
}
