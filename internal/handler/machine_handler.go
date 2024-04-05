package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/dto"
	"github.com/touutae-lab/vending-api/internal/lib"
	"github.com/touutae-lab/vending-api/internal/service"
	"golang.org/x/net/context"
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

	api.POST("/login", h.MachineLogin)
	api.POST("/", h.CreateMachine)
	api.Use(lib.JWTMachineAuthMiddleWare())
	{
		api.GET("/", h.GetAllMachine)
		api.GET("/:id", h.GetMachineByID)
	}
}

func (h *MachineHandler) MachineLogin(c *gin.Context) {
	var request dto.MachineLoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	_, err := h.Service.GetMachineByID(*h.Context, request.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(404, gin.H{"error": "Machine not found"})
			return
		}
		c.JSON(500, gin.H{"error": "There is somethings wrong when request Machine data"})
		return
	}

	token, err := lib.GenerateMachineJWT(request.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "There is somethings wrong when generate token"})
		return
	}

	var response = dto.MachineLoginResponse{
		Token: token,
	}
	c.JSON(200, response)
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

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	machine, err := h.Service.GetMachineByID(*h.Context, id)
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

func (h *MachineHandler) CreateMachine(c *gin.Context) {
	var request dto.CreateMachineRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	machine := request.ToDomain()
	if machine == nil {
		c.JSON(400, gin.H{"error": "Unable to parse request body"})
	}
	id, err := h.Service.CreateMachine(*h.Context, *machine)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}
