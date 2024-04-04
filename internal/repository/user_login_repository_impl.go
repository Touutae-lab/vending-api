package repository

import (
	"context"
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/table"
)

type UserLoginRepositoryImpl struct {
	db *sql.DB
}

func NewUserLoginRepository(db *sql.DB) *UserLoginRepositoryImpl {
	return &UserLoginRepositoryImpl{
		db: db,
	}
}

// InsertUserLogin inserts a new user login record into the database.
func (ulr *UserLoginRepositoryImpl) InsertUserLogin(ctx context.Context, userLogin *model.UserLogin) (int64, error) {
	query := table.UserLogin.INSERT(
		table.UserLogin.Username,
		table.UserLogin.Password,
	).VALUES(
		userLogin.Username,
		userLogin.Password,
	).RETURNING(
		table.UserLogin.ID,
	)

	queryString, args := query.Sql()

	result, err := ulr.db.ExecContext(ctx, queryString, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (ulr *UserLoginRepositoryImpl) GetUserLoginByUsername(ctx context.Context, username string) (model.UserLogin, error) {
	query := table.UserLogin.SELECT(
		table.UserLogin.AllColumns,
	).WHERE(
		table.UserLogin.Username.EQ(postgres.String(username)),
	).LIMIT(1)

	queryString, args := query.Sql()
	row := ulr.db.QueryRowContext(ctx, queryString, args...)

	var userLogin model.UserLogin
	if err := row.Scan(&userLogin.ID, &userLogin.Username, &userLogin.Password); err != nil {
		return model.UserLogin{}, err
	}
	return userLogin, nil
}

func (ulr *UserLoginRepositoryImpl) UpdateUserLogin(ctx context.Context, userLogin *model.UserLogin) (int64, error) {
	query := table.UserLogin.UPDATE().SET(
		table.UserLogin.Username.SET(postgres.String(userLogin.Username)),
		table.UserLogin.Password.SET(postgres.String(userLogin.Password)),
	).WHERE(
		table.UserLogin.ID.EQ(postgres.Int32(userLogin.ID)),
	)

	queryString, args := query.Sql()
	result, err := ulr.db.ExecContext(ctx, queryString, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (ulr *UserLoginRepositoryImpl) DeleteUserLogin(ctx context.Context, username string) (int64, error) {
	query := table.UserLogin.DELETE().WHERE(
		table.UserLogin.Username.EQ(postgres.String(username)),
	)
	queryString, args := query.Sql()
	result, err := ulr.db.ExecContext(ctx, queryString, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
