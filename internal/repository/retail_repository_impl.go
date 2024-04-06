package repository

import (
	"database/sql"
	"github.com/google/uuid"
)

type RetailRepositoryImpl struct {
	db *sql.DB
}

func NewRetailRepositoryImpl(db *sql.DB) RetailRepository {
	return &RetailRepositoryImpl{
		db: db,
	}
}

func (r *RetailRepositoryImpl) BuyProduct(machineID uuid.UUID, productID uuid.UUID, quantity int) (uuid.UUID, error) {
	return uuid.New(), nil
}
