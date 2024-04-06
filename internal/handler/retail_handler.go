package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/touutae-lab/vending-api/internal/dto"
	"github.com/touutae-lab/vending-api/internal/lib"
	"github.com/touutae-lab/vending-api/internal/service"
	"golang.org/x/net/context"
	"net/http"
)

type RetailHandler struct {
	Context *context.Context
	Service service.RetailService
}

func NewRetailHandler(ctx *context.Context, retailService service.RetailService) *RetailHandler {
	return &RetailHandler{
		Context: ctx,
		Service: retailService,
	}
}

func (h *RetailHandler) RegisterRoute(router *gin.Engine) {
	api := router.Group("/retail")
	api.Use(lib.JWTAuthMiddleWare())
}

func (h *RetailHandler) BuyItem(c *gin.Context) {
	var request dto.BuyItemRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := request.GetUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid machine id"})
		return
	}

	_, err = h.Service.BuyProduct(*h.Context, id, request.ProductID, request.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully buy item"})
}
