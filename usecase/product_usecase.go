package usecase

import (
	"GoApi/model"
	"GoApi/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository

}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
	// Lógica de negócio para obter os produtos
}