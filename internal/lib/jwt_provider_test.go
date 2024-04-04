package lib

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMachineJWT(t *testing.T) {
	// Set up
	machineID := uuid.New()
	expectedLength := 100 // Rough estimate of token length, actual may vary

	// Execution
	token, err := GenerateMachineJWT(machineID)

	// Validation
	assert.Nil(t, err)
	assert.Greater(t, len(token), expectedLength)
}

func TestVerifyMachineJWT(t *testing.T) {
	// Set up
	machineID := uuid.New()
	token, _ := GenerateMachineJWT(machineID)

	// Execution
	machineClaims, err := VerifyMachineJWT(token)

	// Validation
	assert.Nil(t, err)
	assert.Equal(t, machineID, machineClaims.MachineID)
}

func TestGenerateJWT(t *testing.T) {
	// Set up
	username := "testuser"
	role := "admin"
	expectedLength := 100 // Rough estimate of token length

	// Execution
	token, err := GenerateJWT(username, role)

	// Validation
	assert.Nil(t, err)
	assert.Greater(t, len(token), expectedLength)
}

func TestVerifyJWT(t *testing.T) {
	// Set
	username := "testuser"
	role := "admin"
	token, _ := GenerateJWT(username, role)

	// Execution
	authorized, err := VerifyJWT(token)
	if err != nil {
		t.Error(err)
	}
	// Validation
	assert.Nil(t, err)
	assert.Equal(t, authorized.Username, username)
	assert.Equal(t, authorized.Role, role)
}
