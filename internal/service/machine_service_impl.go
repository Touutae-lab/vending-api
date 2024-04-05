package service

import (
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"github.com/touutae-lab/vending-api/internal/repository"
	"golang.org/x/net/context"
)

type MachineServiceImpl struct {
	machineRepo repository.MachineRepository
}

func NewMachineService(repo repository.MachineRepository) *MachineServiceImpl {
	return &MachineServiceImpl{
		machineRepo: repo,
	}
}

func (ms *MachineServiceImpl) GetMachineByID(ctx context.Context, id uuid.UUID) (model.Machine, error) {
	result, err := ms.machineRepo.GetMachineByID(ctx, id)
	if err != nil {
		return model.Machine{}, err
	}
	return result, nil
}

func (ms *MachineServiceImpl) GetAllMachine(ctx context.Context) ([]model.Machine, error) {
	result, err := ms.machineRepo.GetAllMachine(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ms *MachineServiceImpl) CreateMachine(ctx context.Context, machine model.Machine) (uuid.UUID, error) {
	return ms.machineRepo.CreateMachine(ctx, machine)
}

func (ms *MachineServiceImpl) UpdateMachine(ctx context.Context, machine model.Machine) (uuid.UUID, error) {
	return ms.machineRepo.UpdateMachine(ctx, machine)
}

func (ms *MachineServiceImpl) DeleteMachine(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	return ms.machineRepo.DeleteMachine(ctx, id)
}
