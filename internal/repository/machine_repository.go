package repository

import (
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"golang.org/x/net/context"
)

type MachineRepository interface {
	GetMachineByID(ctx context.Context, id uuid.UUID) (model.Machine, error)
	GetAllMachine(ctx context.Context) ([]model.Machine, error)
	CreateMachine(ctx context.Context, machine model.Machine) (uuid.UUID, error)
	UpdateMachine(ctx context.Context, machine model.Machine) (uuid.UUID, error)
	DeleteMachine(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}
