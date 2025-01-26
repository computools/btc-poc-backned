package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"btc-backend/internal/companyperformance"
)

func (h *Handler) GetCompanyPerformanceHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	c, err := h.companyPerformance.GetCompanyPerformance(ctx, id)
	if err != nil {
		slog.Error("Failed to get company", "err", err, "function", "GetCompanyPerformanceHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

func (h *Handler) GetCompanyPerformancesByCompanyIDHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Query("company_id"))

	companies, err := h.companyPerformance.GetCompanyPerformancesByCompanyID(ctx, id)
	if err != nil {
		slog.Error("Failed to get companies", "err", err, "function", "GetCompanyPerformancesByCompanyIDHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

func (h *Handler) CreateCompanyPerformanceHandler(ctx *gin.Context) {
	var c companyperformance.CompanyPerformance
	if err := ctx.BindJSON(&c); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "CreateCompanyPerformanceHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&c); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "CreateCompanyPerformanceHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	c, err := h.companyPerformance.CreateCompanyPerformance(ctx, c)
	if err != nil {
		slog.Error("Failed to create company performance", "err", err, "function", "CreateCompanyPerformanceHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, c)
}

func (h *Handler) UpdateCompanyPerformanceHandler(ctx *gin.Context) {
	var c companyperformance.CompanyPerformance
	if err := ctx.BindJSON(&c); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "UpdateCompanyPerformanceHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&c); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "UpdateCompanyPerformanceHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := cast.ToInt64(ctx.Param("id"))
	c.ID = id

	c, err := h.companyPerformance.UpdateCompanyPerformance(ctx, c)
	if err != nil {
		slog.Error("Failed to update company performance", "err", err, "function", "UpdateCompanyPerformanceHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

func (h *Handler) DeleteCompanyPerformanceHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	err := h.companyPerformance.DeleteCompanyPerformance(ctx, id)
	if err != nil {
		slog.Error("Failed to delete company performance", "err", err, "function", "DeleteCompanyPerformanceHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
