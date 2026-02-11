package usecase

import "GoApi/model"

type ProductUseCase struct {
	// Repository
}

func NewProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
	// Lógica de negócio para obter os produtos
}