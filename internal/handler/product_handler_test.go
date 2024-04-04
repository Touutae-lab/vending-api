package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/touutae-lab/vending-api/internal/dto"
	"github.com/touutae-lab/vending-api/internal/lib"
	mocks "github.com/touutae-lab/vending-api/test"
	"go.uber.org/mock/gomock"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductHandler_CreateProduct(t *testing.T) {
	// Initialize Gin
	gin.SetMode(gin.TestMode)

	// Create a new GoMock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // Assert that all expected calls were made

	// Create a mock ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Add JWT
	token, err := lib.GenerateJWT("any", "any")
	if err != nil {
		t.Error(err)
	}

	// Expected data
	productDTO := dto.ProductDTO{
		Name:  "Test Product",
		Price: 100,
	}
	expectedID := int64(1)

	// Setup expectation
	mockProductService.EXPECT().CreateProduct(gomock.Any(), productDTO).Return(expectedID, nil)

	// Create ProductHandler with mocked service
	ctx := context.Background()
	handler := NewProductHandler(&ctx, mockProductService)

	// Setup Gin context for HTTP request
	router := gin.Default()
	handler.RegisterRoute(router)

	// Marshal body data
	bodyData, _ := json.Marshal(productDTO)
	req, _ := http.NewRequest(http.MethodPost, "/product/", bytes.NewBuffer(bodyData))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.Code)
	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)
	assert.Equal(t, float64(expectedID), respBody["id"])
}

func TestProductHandler_GetAllProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mocks.NewMockProductService(ctrl)

	// Add JWT
	token, err := lib.GenerateJWT("any", "any")
	if err != nil {
		t.Error(err)
	}

	// Expected data
	expectedProducts := []dto.ProductDTO{
		{
			ID:    1,
			Name:  "Test Product",
			Price: 100,
		},
	}

	// Setup expectation
	mockProductService.EXPECT().GetAllProduct(gomock.Any()).Return(expectedProducts, nil)

	// Create ProductHandler with mocked service
	ctx := context.Background()
	handler := NewProductHandler(&ctx, mockProductService)

	// Setup Gin context for HTTP request
	router := gin.Default()
	handler.RegisterRoute(router)

	// Create HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/product/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var respBody []dto.ProductDTO
	json.Unmarshal(resp.Body.Bytes(), &respBody)
	assert.Equal(t, expectedProducts, respBody)
}
