package service

import (
	"context"
	"github.com/google/uuid"
)

type RetailService interface {
	BuyProduct(ctx context.Context, machineID uuid.UUID, productID int32, quantity int, totalPrice int) (int32, error)
}
