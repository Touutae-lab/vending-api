package service

import (
	"github.com/touutae-lab/vending-api/internal/dto"
	"github.com/touutae-lab/vending-api/internal/repository"
	"golang.org/x/net/context"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: repo,
	}
}

func (ps *ProductService) GetAllProduct(ctx context.Context) ([]dto.ProductDTO, error) {
	products, err := ps.productRepo.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	var dtos []dto.ProductDTO
	for _, p := range products {
		dtos = append(dtos, dto.ToProductDTO(p))
	}
	return dtos, nil
}

func (ps *ProductService) GetProductByID(ctx context.Context, id int32) (dto.ProductDTO, error) {
	product, err := ps.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return dto.ProductDTO{}, err
	}

	return dto.ToProductDTO(product), nil
}

func (ps *ProductService) CreateProduct(ctx context.Context, product dto.ProductDTO) (int64, error) {
	return ps.productRepo.CreateProduct(ctx, product.ToDomain())
}

func (ps *ProductService) UpdateProduct(ctx context.Context, product dto.ProductDTO) (int64, error) {
	return ps.productRepo.UpdateProduct(ctx, product.ToDomain())
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id int32) (int64, error) {
	return ps.productRepo.DeleteProduct(ctx, id)
}
