package repository

import (
	"context"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
)

type UserLoginRepository interface {
	InsertUserLogin(ctx context.Context, userLogin *model.UserLogin) (int64, error)
	GetUserLoginByUsername(ctx context.Context, username string) (model.UserLogin, error)
	UpdateUserLogin(ctx context.Context, userLogin *model.UserLogin) (int64, error)
	DeleteUserLogin(ctx context.Context, username string) (int64, error)
}
