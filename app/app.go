package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"btc-backend/config"
	authservice "btc-backend/internal/auth"
	"btc-backend/internal/company"
	"btc-backend/internal/companyperformance"
	"btc-backend/internal/financialreports"
	"btc-backend/internal/heartbeat"
	"btc-backend/internal/user"
	"btc-backend/pkg/auth"
	"btc-backend/pkg/database"
	"btc-backend/server"
	"btc-backend/server/handlers"
	"btc-backend/server/middlewares"
)

func Exec(cfg *config.Config) fx.Option {
	return fx.Options(
		fx.Supply(cfg),
		fx.Provide(
			// ^ Libraries
			resty.New,
			validator.New,

			// ^ Repositories
			fx.Annotate(database.NewPostgres,
				fx.As(new(databaseHook)),
				fx.As(new(heartbeat.Database)),
				fx.As(new(company.Database)),
				fx.As(new(user.Database)),
				fx.As(new(financialreports.Database)),
				fx.As(new(companyperformance.Database)),
			),
			fx.Annotate(auth.NewKeycloak,
				fx.As(new(authHook)),
				fx.As(new(user.Auth)),
				fx.As(new(authservice.Auth)),
				fx.As(new(middlewares.Auth)),
			),

			// ^ Services
			heartbeat.NewService,
			company.NewService,
			user.NewService,
			financialreports.NewService,
			authservice.NewService,
			companyperformance.NewService,

			// ^ Server
			handlers.NewHandler,
			server.NewHTTPServer,
		),
		fx.Invoke(
			prepareHooks,
		),
	)
}

type hooks struct {
	fx.In

	Database databaseHook
	Auth     authHook
	Server   *server.HTTPServer
}

func prepareHooks(lc fx.Lifecycle, hooks hooks) {
	lc.Append(fx.Hook{OnStart: hooks.Database.Start, OnStop: hooks.Database.Stop})
	lc.Append(fx.Hook{OnStart: hooks.Auth.Start, OnStop: hooks.Auth.Stop})
	lc.Append(fx.Hook{OnStart: hooks.Server.Start, OnStop: hooks.Server.Stop})
}
