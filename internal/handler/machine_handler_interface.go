package handler

import (
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"golang.org/x/net/context"
)

type MachineHandlerInterface interface {
	GetMachineTypeByID(ctx context.Context, id int32) (model.Machinetype, error)
	GetMachineByID(ctx context.Context, id uuid.UUID) (model.Machine, error)
	GetAllMachine(ctx context.Context) ([]model.Machine, error)
	CreateMachine(ctx context.Context, machine model.Machine) (int64, error)
	UpdateMachine(ctx context.Context, machine model.Machine) (int64, error)
	DeleteMachine(ctx context.Context, id uuid.UUID) (int64, error)
}
