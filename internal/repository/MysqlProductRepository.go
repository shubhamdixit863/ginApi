package repository

import (
	"ginTraining/internal/entity"
	"gorm.io/gorm"
)

type MysqlProductRepository struct {

	// We will define our database object
	Db *gorm.DB
}

func (my *MysqlProductRepository) AddProduct(product entity.Product) (*entity.Product, error) {
	return nil, nil
}
func (my *MysqlProductRepository) GetProductById(id int) (*entity.Product, error) {
	return nil, nil

}
func (my *MysqlProductRepository) GetProducts() ([]entity.Product, error) {
	return nil, nil

}
