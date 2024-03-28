package usecase

import (
	"errors"
	"fmt"
	"stock/domain/model"
	"stock/domain/repository"
)

type ProductUseCase interface {
	Insert(productName string, brandName string, imagePath []string) error
	GetAll() ([]*model.Product, error)
	GetAllByBrand(brandName string) ([]*model.Product, error)
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(pr repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepository: pr,
	}
}

// productName, brandName, imagePathを指定してデータベースにProductを追加する
func (pu productUseCase) Insert(productName string, brandName string, imagePath []string) error {
	if productName == "" {
		return errors.New("productName is empty")
	}
	if brandName == "" {
		return errors.New("brandName is empty")
	}
	if len(imagePath) == 0 {
		return errors.New("imagePath is empty")
	}
	// SQL文を実行する
	if err := pu.productRepository.Insert(productName, brandName, imagePath); err != nil {
		return fmt.Errorf("failed to insert product in Insert: %w", err)
	}
	return nil
}

// データベースの全てのProductを取得する
func (pu productUseCase) GetAll() ([]*model.Product, error) {
	// SQL文を実行する
	products, err := pu.productRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all products in GetAll: %w", err)
	}
	return products, nil
}

// brandNameを指定してデータベースのProductを取得する
func (pu productUseCase) GetAllByBrand(brandName string) ([]*model.Product, error) {
	if brandName == "" {
		return nil, errors.New("brandName is empty")
	}
	// SQL文を実行する
	products, err := pu.productRepository.GetAllByBrand(brandName)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products by brand in GetAllByBrand: %w", err)
	}
	return products, nil
}

