package app

import (
	"context"
)

type hook interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type databaseHook interface {
	hook
}

type authHook interface {
	hook
}
