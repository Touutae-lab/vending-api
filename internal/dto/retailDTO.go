package dto

import "github.com/google/uuid"

type BuyItemRequest struct {
	MachineID string `json:"machine_id"`
	ProductID int32  `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (btr *BuyItemRequest) GetUUID() (uuid.UUID, error) {
	id, err := uuid.Parse(btr.MachineID)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
