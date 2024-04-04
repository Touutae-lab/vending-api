package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"github.com/touutae-lab/vending-api/internal/dto"
	"github.com/touutae-lab/vending-api/internal/service"
	mock_repository "github.com/touutae-lab/vending-api/test"
	"go.uber.org/mock/gomock"
	"golang.org/x/net/context"
	"testing"
)

func TestGetAllProduct(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockProductRepository(ctrl)

	// Sample data to return
	products := []model.Product{{ID: 1, Name: "Test Product"}}

	// Set expectation
	mockRepo.EXPECT().GetAllProduct(gomock.Any()).Return(products, nil)

	// Create service with mock
	productService := service.NewProductService(mockRepo)

	// Call the method under test
	result, err := productService.GetAllProduct(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Product", result[0].Name)
}

func TestGetProductByID(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockProductRepository(ctrl)

	// Sample data to return
	product := model.Product{ID: 1, Name: "Test Product"}

	// Set expectation
	mockRepo.EXPECT().GetProductByID(gomock.Any(), int32(1)).Return(product, nil)

	// Create service with mock
	productService := service.NewProductService(mockRepo)

	// Call the method under test
	result, err := productService.GetProductByID(context.Background(), 1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "Test Product", result.Name)
}

func TestCreateProduct(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockProductRepository(ctrl)

	// Set expectation
	mockRepo.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	// Create service with mock
	productService := service.NewProductService(mockRepo)

	// Call the method under test
	result, err := productService.CreateProduct(context.Background(), dto.ToProductDTO(model.Product{}))

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func TestUpdateProduct(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockProductRepository(ctrl)

	// Set expectation
	mockRepo.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	// Create service with mock
	productService := service.NewProductService(mockRepo)

	// Call the method under test
	result, err := productService.UpdateProduct(context.Background(), dto.ToProductDTO(model.Product{}))

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func TestDeleteProduct(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockProductRepository(ctrl)

	// Set expectation
	mockRepo.EXPECT().DeleteProduct(gomock.Any(), int32(1)).Return(int64(1), nil)

	// Create service with mock
	productService := service.NewProductService(mockRepo)

	// Call the method under test
	result, err := productService.DeleteProduct(context.Background(), 1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}
