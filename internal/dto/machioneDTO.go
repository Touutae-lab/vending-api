package dto

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
)

type MachineLoginRequest struct {
	ID uuid.UUID `json:"id"`
}

type MachineLoginResponse struct {
	Token string `json:"token"`
}

type CreateMachineRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

type Storage struct {
	ProductID int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

func (request *CreateMachineRequest) ToDomain() *model.Machine {
	storageDetails, err := json.Marshal([]Storage{})
	uuid := uuid.New()
	if err != nil {
		return nil
	}
	return &model.Machine{
		UUID:           uuid,
		Name:           request.Name,
		Location:       request.Location,
		Status:         request.Status,
		StorageDetails: string(storageDetails),
	}
}
