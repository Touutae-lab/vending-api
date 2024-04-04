package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, _ := HashPassword(password)

	assert.True(t, CheckPasswordHash(password, hash))
	assert.False(t, CheckPasswordHash("wrongpassword", hash))
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, _ := HashPassword(password)

	assert.True(t, CheckPasswordHash(password, hash))
	assert.False(t, CheckPasswordHash("wrongpassword", hash))
}
