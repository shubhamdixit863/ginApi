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

	tx := my.Db.Create(&product)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &product, nil
}
func (my *MysqlProductRepository) GetProductById(id int) (*entity.Product, error) {
	return nil, nil

}
func (my *MysqlProductRepository) GetProducts() ([]entity.Product, error) {
	return nil, nil

}
