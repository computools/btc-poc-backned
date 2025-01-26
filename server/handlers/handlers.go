package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

	"btc-backend/internal/auth"
	"btc-backend/internal/company"
	"btc-backend/internal/companyperformance"
	"btc-backend/internal/financialreports"
	"btc-backend/internal/heartbeat"
	"btc-backend/internal/user"
)

type Handler struct {
	heartbeat          *heartbeat.Service
	company            *company.Service
	user               *user.Service
	financialReports   *financialreports.Service
	auth               *auth.Service
	companyPerformance *companyperformance.Service

	validate *validator.Validate
	resty    *resty.Client
}

func (h *Handler) Heartbeat(c *gin.Context) {
	err := h.heartbeat.Ping(c)
	if err != nil {
		slog.Error("Failed to ping", "err", err, "function", "Heartbeat")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func NewHandler(
	company *company.Service,
	user *user.Service,
	financialReports *financialreports.Service,
	auth *auth.Service,
	heartbeat *heartbeat.Service,
	companyPerformance *companyperformance.Service,
	resty *resty.Client,
	validate *validator.Validate,
) *Handler {
	return &Handler{
		auth:               auth,
		company:            company,
		user:               user,
		financialReports:   financialReports,
		companyPerformance: companyPerformance,
		heartbeat:          heartbeat,
		resty:              resty,
		validate:           validate,
	}
}
