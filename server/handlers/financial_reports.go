package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"btc-backend/internal/financialreports"
)

func (h *Handler) GetFinancialReportHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	u, err := h.financialReports.GetFinancialReport(ctx, id)
	if err != nil {
		slog.Error("Failed to get financial report", "err", err, "function", "GetFinancialReportHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (h *Handler) GetFinancialReportsByCompanyIDHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Query("company_id"))

	fs, err := h.financialReports.GetFinancialReportsByCompanyID(ctx, id)
	if err != nil {
		slog.Error("Failed to get financial report by company id", "err", err, "function", "GetFinancialReportsByCompanyIDHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, fs)
}

func (h *Handler) CreateFinancialReportHandler(ctx *gin.Context) {
	var fr financialreports.FinancialReports
	if err := ctx.BindJSON(&fr); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "CreateFinancialReportHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&fr); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "CreateFinancialReportHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	fr, err := h.financialReports.CreateFinancialReport(ctx, fr)
	if err != nil {
		slog.Error("Failed to create financial report", "err", err, "function", "CreateFinancialReportHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, fr)
}

func (h *Handler) UpdateFinancialReportHandler(ctx *gin.Context) {
	var fr financialreports.FinancialReports
	if err := ctx.BindJSON(&fr); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "UpdateFinancialReportHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&fr); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "UpdateFinancialReportHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := cast.ToInt64(ctx.Param("id"))
	fr.ID = id

	fr, err := h.financialReports.UpdateFinancialReport(ctx, fr)
	if err != nil {
		slog.Error("Failed to update financial report", "err", err, "function", "UpdateFinancialReportHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, fr)
}

func (h *Handler) DeleteFinancialReportHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	err := h.financialReports.DeleteFinancialReport(ctx, id)
	if err != nil {
		slog.Error("Failed to delete f", "err", err, "function", "DeleteFinancialReportHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
