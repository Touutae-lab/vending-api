package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/table"
	"time"
)

type OrderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (or *OrderRepositoryImpl) CreateOrder(machineID uuid.UUID, productID int32, quantity int, totalPrice int) (int32, error) {

	query := table.Order.INSERT(
		table.Order.MachineID,
		table.Order.ProductID,
		table.Order.Quantity,
		table.Order.TotalPrice,
		table.Order.OrderDate,
	).VALUES(
		machineID,
		productID,
		quantity,
		totalPrice,
		time.Now().UTC(),
	).RETURNING(
		table.Order.ID,
	)

	queryString, args := query.Sql()

	var id int32
	err := or.db.QueryRow(queryString, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
