package middlewares

import "context"

type Auth interface {
	ValidateToken(ctx context.Context, accessToken string) error
}
