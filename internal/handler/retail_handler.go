package handler

import (
	"fmt"
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
	service.ProductService
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
	{
		api.POST("/buy", h.BuyItem)
	}
}

func (h *RetailHandler) BuyItem(c *gin.Context) {
	var request dto.BuyItemRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := h.Service.BuyProduct(*h.Context, request.MachineID, request.ProductID, request.Quantity, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// This should have validation step, but since this for example project, so I will skip it
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully buy item get your exchange back in the tray for $%d", request.Payment-request.Amount*request.Quantity)})
}
