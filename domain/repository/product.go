package repository

import (
	"stock/domain/model"
)

// infra層、usecase層がこのinterfaceに依存する
type ProductRepository interface {
	GetAll() ([]*model.Product, error)
	GetAllByBrand(brandName string) ([]*model.Product, error)
	Insert(productName string, brandName string, imagePath []string) error
}
