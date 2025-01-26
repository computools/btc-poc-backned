package auth

import (
	"context"

	"btc-backend/pkg/auth"
)

type Auth interface {
	Login(ctx context.Context, user, password string) (auth.JWT, error)
	RefreshToken(ctx context.Context, refreshToken string) (auth.JWT, error)
	Logout(ctx context.Context, refreshToken string) error
}
