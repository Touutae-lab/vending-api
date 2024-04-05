package service_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"github.com/touutae-lab/vending-api/internal/service"
	mock_repository "github.com/touutae-lab/vending-api/test"
	"go.uber.org/mock/gomock"
	"golang.org/x/net/context"
	"testing"
)

func TestMachineService_GetAllMachine(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockMachineRepository(ctrl)

	// Set expectation
	mockRepo.EXPECT().GetAllMachine(gomock.Any()).Return([]model.Machine{}, nil)

	// Create service with mock
	machineService := service.NewMachineService(mockRepo)

	// Call the method under test
	result, err := machineService.GetAllMachine(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 0, len(result))
}

func TestMachineService_GetMachineByID(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockMachineRepository(ctrl)

	// Set expectation
	expect := uuid.New()

	mapMachine := make(map[uuid.UUID]model.Machine)
	mapMachine[expect] = model.Machine{
		UUID:           expect,
		Name:           "test",
		Location:       "test",
		Status:         "test",
		StorageDetails: "test",
	}
	mockRepo.EXPECT().GetMachineByID(gomock.Any(), expect).Return(mapMachine[expect], nil)

	// Create service with mock
	machineService := service.NewMachineService(mockRepo)

	// Call the method under test
	result, err := machineService.GetMachineByID(context.Background(), expect)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mapMachine[expect], result)
}

func TestMachineService_CreateMachine(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockMachineRepository(ctrl)

	// Set expectation
	mockRepo.EXPECT().CreateMachine(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

	// Create service with mock
	machineService := service.NewMachineService(mockRepo)

	// Call the method under test
	result, err := machineService.CreateMachine(context.Background(), model.Machine{})

	// Assertions
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, result)
}

func TestMachineService_UpdateMachine(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockMachineRepository(ctrl)

	// Set expectation
	mockRepo.EXPECT().UpdateMachine(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

	// Create service with mock
	machineService := service.NewMachineService(mockRepo)

	// Call the method under test
	result, err := machineService.UpdateMachine(context.Background(), model.Machine{})

	// Assertions
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, result)
}

func TestMachineService_DeleteMachine(t *testing.T) {
	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock_repository.NewMockMachineRepository(ctrl)

	// Set expectation
	mockRepo.EXPECT().DeleteMachine(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

	// Create service with mock
	machineService := service.NewMachineService(mockRepo)

	// Call the method under test
	result, err := machineService.DeleteMachine(context.Background(), uuid.New())

	// Assertions
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, result)
}
