package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"btc-backend/config"
	"btc-backend/server/handlers"
	"btc-backend/server/middlewares"
)

type HTTPServer struct {
	server *http.Server

	cfg      *config.ServerConfig
	handlers *handlers.Handler
	auth     middlewares.Auth
}

func (s *HTTPServer) Start(_ context.Context) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middlewares.CorsMiddleware(s.cfg.CorsMaxAge),
		middlewares.LoggerMiddleware(),
		gin.Recovery(),
	)

	apiv1 := router.Group("/api/v1")
	apiv1.GET("/heartbeat", s.handlers.Heartbeat)
	apiv1.POST("/users", s.handlers.CreateUserHandler)
	apiv1.POST("/login", s.handlers.Login)

	apiv1Private := apiv1.Group("", middlewares.AuthMiddleware(s.auth))

	// Company
	apiv1Private.GET("/companies/:id", s.handlers.GetCompanyHandler)
	apiv1Private.GET("/companies", s.handlers.GetCompaniesHandler)
	apiv1Private.POST("/companies", s.handlers.CreateCompanyHandler)
	apiv1Private.PUT("/companies/:id", s.handlers.UpdateCompanyHandler)
	apiv1Private.DELETE("/companies/:id", s.handlers.DeleteCompanyHandler)

	// User
	apiv1Private.GET("/users/:id", s.handlers.GetUserHandler)
	apiv1Private.GET("/users", s.handlers.GetUsersByCompanyIDHandler)
	apiv1Private.PUT("/users/:id", s.handlers.UpdateUserHandler)
	apiv1Private.DELETE("/users/:id", s.handlers.DeleteUserHandler)

	// Financial report
	apiv1Private.GET("/reports/:id", s.handlers.GetFinancialReportHandler)
	apiv1Private.GET("/reports", s.handlers.GetFinancialReportsByCompanyIDHandler)
	apiv1Private.POST("/reports", s.handlers.CreateFinancialReportHandler)
	apiv1Private.PUT("/reports/:id", s.handlers.UpdateFinancialReportHandler)
	apiv1Private.DELETE("/reports/:id", s.handlers.DeleteFinancialReportHandler)

	// Auth
	apiv1Private.POST("/refresh", s.handlers.RefreshToken)
	apiv1Private.POST("/logout", s.handlers.Logout)

	// Company performance
	apiv1Private.GET("/company-reports/:id", s.handlers.GetCompanyPerformanceHandler)
	apiv1Private.GET("/company-reports", s.handlers.GetCompanyPerformancesByCompanyIDHandler)
	apiv1Private.POST("/company-reports", s.handlers.CreateCompanyPerformanceHandler)
	apiv1Private.PUT("/company-reports/:id", s.handlers.UpdateCompanyPerformanceHandler)
	apiv1Private.DELETE("/company-reports/:id", s.handlers.DeleteCompanyPerformanceHandler)

	s.server = &http.Server{
		Addr:              s.cfg.Addr,
		Handler:           router,
		ReadHeaderTimeout: s.cfg.ReadHeaderTimeout,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server start error", "err", err)
			panic(err)
		}
	}()

	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func NewHTTPServer(cfg *config.Config, handlers *handlers.Handler, auth middlewares.Auth) *HTTPServer {
	return &HTTPServer{
		cfg:      &cfg.ServerConfig,
		handlers: handlers,
		auth:     auth,
	}
}
