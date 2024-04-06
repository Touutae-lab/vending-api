package service

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/dto"
	"github.com/touutae-lab/vending-api/internal/repository"
	"golang.org/x/net/context"
)

type RetailServiceImpl struct {
	productRepo repository.ProductRepository
	machineRepo repository.MachineRepository
	orderRepo   repository.OrderRepository
}

func NewRetailService(productRepo repository.ProductRepository, machineRepo repository.MachineRepository, orderRepo repository.OrderRepository) *RetailServiceImpl {
	return &RetailServiceImpl{
		productRepo: productRepo,
		machineRepo: machineRepo,
		orderRepo:   orderRepo,
	}
}

func (rs *RetailServiceImpl) BuyProduct(ctx context.Context, machineID uuid.UUID, productID int32, quantity int) (int32, error) {

	orderId, err := rs.orderRepo.CreateOrder(machineID, productID, quantity)

	if err != nil {
		return orderId, err
	}

	machine, err := rs.machineRepo.GetMachineByID(ctx, machineID)

	if err != nil {
		return orderId, err
	}

	var storage []dto.Storage

	err = json.Unmarshal([]byte(machine.StorageDetails), storage)

	if err != nil {
		return orderId, err
	}

	for i, s := range storage {
		if s.ProductID == productID {
			storage[i].Quantity -= int32(quantity)
			break
		}
	}

	storageDetails, err := json.Marshal(storage)

	if err != nil {
		return orderId, err
	}

	machine.StorageDetails = string(storageDetails)

	_, err = rs.machineRepo.UpdateMachine(ctx, machine)

	if err != nil {
		return orderId, err
	}

	return orderId, nil
}
