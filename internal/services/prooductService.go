package services

import (
	"ginTraining/internal/dtos"
	"ginTraining/internal/entity"
	"ginTraining/internal/repository"
	"time"
)

type ProductService struct {
	ProductRepo repository.IProductRepository
}

func (pr *ProductService) SaveProduct(productDto dtos.ProductDto) error {

	// We will convert dto to entity and pass it in our repo

	product := entity.Product{
		Name:         productDto.Name,
		ProductImage: productDto.Image,
		Description:  productDto.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// We can call our repository

	_, err := pr.ProductRepo.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
}
