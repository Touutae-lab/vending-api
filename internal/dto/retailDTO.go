package dto

import "github.com/google/uuid"

type BuyItemRequest struct {
	MachineID uuid.UUID `json:"machine_id"`
	ProductID int32     `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Amount    int       `json:"amount"`
	Payment   int       `json:"payment"`
}
