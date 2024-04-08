package repository

import "github.com/google/uuid"

type OrderRepository interface {
	CreateOrder(machineID uuid.UUID, productID int32, quantity int, totalPrice int) (int32, error)
}
