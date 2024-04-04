package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/touutae-lab/vending-api/internal/dto"
	"github.com/touutae-lab/vending-api/internal/service"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Context *context.Context
	Service *service.ProductService
}

func NewProductHandler(ctx *context.Context, productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Context: ctx,
		Service: productService,
	}
}

func (h *ProductHandler) RegisterRoute(router *gin.Engine) {
	router.GET("/product", h.GetAllProduct)
	router.GET("/product/:id", h.GetProductByID)
	router.POST("/product", h.CreateProduct)
	router.PUT("/product", h.UpdateProduct)
	router.DELETE("/product/:id", h.DeleteProduct)
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	products, err := h.Service.GetAllProduct(*h.Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)

}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	product, err := h.Service.GetProductByID(*h.Context, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productDTO dto.ProductDTO
	if err := c.BindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Service.CreateProduct(*h.Context, productDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var productDTO dto.ProductDTO
	if err := c.BindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Service.UpdateProduct(*h.Context, productDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	rowsAffected, err := h.Service.DeleteProduct(*h.Context, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rows_affected": rowsAffected})
}
