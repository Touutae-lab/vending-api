package repository

import "github.com/google/uuid"

type RetailRepository interface {
	BuyProduct(machineID uuid.UUID, productID uuid.UUID, quantity int) (uuid.UUID, error)
}
