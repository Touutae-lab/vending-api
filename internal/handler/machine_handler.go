package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/touutae-lab/vending-api/internal/lib"
	"github.com/touutae-lab/vending-api/internal/service"
	"golang.org/x/net/context"
	"strconv"
)

type MachineHandler struct {
	Context *context.Context
	Service service.MachineService
}

func NewMachineHandler(ctx *context.Context, service service.MachineService) *MachineHandler {
	return &MachineHandler{
		Context: ctx,
		Service: service,
	}
}

func (h *MachineHandler) RegisterRoute(router *gin.Engine) {
	api := router.Group("/machine")
	api.Use(lib.JWTMachineAuthMiddleWare())
	{
		api.GET("/", h.GetAllMachine)
		api.GET("/:id", h.GetMachineByID)
	}
}

func (h *MachineHandler) GetAllMachine(c *gin.Context) {
	machines, err := h.Service.GetAllMachine(*h.Context)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, machines)
}

func (h *MachineHandler) GetMachineByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	machine, err := h.Service.GetMachineByID(*h.Context, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(404, gin.H{"error": "Machine not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, machine)
}
