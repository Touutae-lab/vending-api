package handler

import (
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"golang.org/x/net/context"
)

type ProductHandlerInterface interface {
	GetProductByID(ctx context.Context, id int32) (model.Product, error)
	GetAllProduct(ctx context.Context) ([]model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) (int64, error)
	UpdateProduct(ctx context.Context, product model.Product) (int64, error)
	DeleteProduct(ctx context.Context, id int32) (int64, error)
}
