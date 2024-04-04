package service

import (
	"github.com/touutae-lab/vending-api/internal/dto"
	"golang.org/x/net/context"
)

type ProductService interface {
	GetAllProduct(ctx context.Context) ([]dto.ProductDTO, error)
	GetProductByID(ctx context.Context, id int32) (dto.ProductDTO, error)
	CreateProduct(ctx context.Context, product dto.ProductDTO) (int64, error)
	UpdateProduct(ctx context.Context, product dto.ProductDTO) (int64, error)
	DeleteProduct(ctx context.Context, id int32) (int64, error)
}
