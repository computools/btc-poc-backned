package user

import (
	"context"

	"btc-backend/pkg/auth"
	"btc-backend/pkg/database"
)

type Database interface {
	GetUser(ctx context.Context, id int64) (database.User, error)
	GetUsersByCompanyID(ctx context.Context, id int64) ([]database.User, error)
	CreateUser(ctx context.Context, user database.User) (database.User, error)
	UpdateUser(ctx context.Context, user database.User) (database.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type Auth interface {
	CreateUser(ctx context.Context, user auth.User, password string) (string, error)
}
