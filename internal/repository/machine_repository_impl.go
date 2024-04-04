package repository

import (
	"context"
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/table"
)

type MachineRepositoryImpl struct {
	db *sql.DB
}

func NewMachineRepository(db *sql.DB) *MachineRepositoryImpl {
	return &MachineRepositoryImpl{
		db: db,
	}
}

func (mr *MachineRepositoryImpl) InsertMachine(ctx context.Context, machine *model.Machine) (int64, error) {
	query := table.Machine.INSERT(
		table.Machine.UUID,
		table.Machine.TypeID,
		table.Machine.Location,
		table.Machine.Status,
		table.Machine.StorageDetails,
	).VALUES(
		machine.UUID,
		machine.TypeID,
		machine.Location,
		machine.Status,
		machine.StorageDetails,
	).RETURNING(
		table.Machine.UUID,
	)

	queryString, args := query.Sql()

	result, err := mr.db.ExecContext(ctx, queryString, args...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (mr *MachineRepositoryImpl) GetMachineByUUID(ctx context.Context, uuid uuid.UUID) (model.Machine, error) {
	query := table.Machine.SELECT(
		table.Machine.AllColumns,
	).WHERE(
		table.Machine.UUID.EQ(postgres.UUID(uuid)),
	).LIMIT(1)

	queryString, args := query.Sql()
	row := mr.db.QueryRowContext(ctx, queryString, args...)

	var machine model.Machine
	if err := row.Scan(&machine.UUID, &machine.TypeID, &machine.Location, &machine.Status, &machine.StorageDetails); err != nil {
		return model.Machine{}, err
	}

	return machine, nil
}

func (mr *MachineRepositoryImpl) GetAllMachine(ctx context.Context) ([]model.Machine, error) {
	query := table.Machine.SELECT(
		table.Machine.AllColumns,
	)

	queryString, args := query.Sql()
	rows, err := mr.db.QueryContext(ctx, queryString, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var machines []model.Machine
	for rows.Next() {
		var machine model.Machine
		if err := rows.Scan(&machine.UUID, &machine.TypeID, &machine.Location, &machine.Status, &machine.StorageDetails); err != nil {
			return nil, err
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func (mr *MachineRepositoryImpl) UpdateMachine(ctx context.Context, machine *model.Machine) (int64, error) {
	query := table.Machine.UPDATE().SET(
		table.Machine.TypeID.SET(postgres.Int32(machine.TypeID)),
		table.Machine.Location.SET(postgres.String(machine.Location)),
		table.Machine.Status.SET(postgres.String(machine.Status)),
		table.Machine.StorageDetails.SET(postgres.String(machine.StorageDetails)),
	).WHERE(
		table.Machine.UUID.EQ(postgres.UUID(machine.UUID)),
	)

	queryString, args := query.Sql()
	result, err := mr.db.ExecContext(ctx, queryString, args...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (mr *MachineRepositoryImpl) DeleteMachine(ctx context.Context, uuid uuid.UUID) (int64, error) {
	query := table.Machine.DELETE().WHERE(
		table.Machine.UUID.EQ(postgres.UUID(uuid)),
	)

	queryString, args := query.Sql()

	result, err := mr.db.ExecContext(ctx, queryString, args...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
