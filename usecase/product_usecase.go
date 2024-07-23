package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		productRepository: repository,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}
