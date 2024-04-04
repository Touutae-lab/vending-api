package dto

import "github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"

type ProductDTO struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func ToProductDTO(product model.Product) ProductDTO {
	return ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Details,
		Price:       product.Price,
	}
}

func (productDTO *ProductDTO) ToDomain() *model.Product {
	return &model.Product{
		ID:      productDTO.ID,
		Name:    productDTO.Name,
		Details: productDTO.Description,
		Price:   productDTO.Price,
	}
}
