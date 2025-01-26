package main

import (
	"go.uber.org/fx"

	"btc-backend/app"
	"btc-backend/config"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	fx.New(app.Exec(cfg)).Run()
}
