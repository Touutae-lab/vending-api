package service

import (
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

func (ms *MachineServiceImpl) GetMachineTypeByID(ctx context.Context, id int32) (model.Machinetype, error) {
	result, err := ms.machineRepo.GetMachineTypeByID(ctx, id)
	if err != nil {
		return model.Machinetype{}, err
	}
	return result, nil
}

func (ms *MachineServiceImpl) GetMachineByID(ctx context.Context, id int32) (model.Machine, error) {
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

func (ms *MachineServiceImpl) CreateMachine(ctx context.Context, machine model.Machine) (int64, error) {
	return ms.machineRepo.CreateMachine(ctx, machine)
}

func (ms *MachineServiceImpl) UpdateMachine(ctx context.Context, machine model.Machine) (int64, error) {
	return ms.machineRepo.UpdateMachine(ctx, machine)
}

func (ms *MachineServiceImpl) DeleteMachine(ctx context.Context, id int32) (int64, error) {
	return ms.machineRepo.DeleteMachine(ctx, id)
}
