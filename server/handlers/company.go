package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"btc-backend/internal/company"
)

func (h *Handler) GetCompanyHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	c, err := h.company.GetCompany(ctx, id)
	if err != nil {
		slog.Error("Failed to get company", "err", err, "function", "GetCompanyHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

func (h *Handler) GetCompaniesHandler(ctx *gin.Context) {
	companies, err := h.company.GetCompanies(ctx)
	if err != nil {
		slog.Error("Failed to get companies", "err", err, "function", "GetCompaniesHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

func (h *Handler) CreateCompanyHandler(ctx *gin.Context) {
	var c company.Company
	if err := ctx.BindJSON(&c); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "CreateCompanyHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&c); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "CreateCompanyHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	c, err := h.company.CreateCompany(ctx, c)
	if err != nil {
		slog.Error("Failed to create company", "err", err, "function", "CreateCompanyHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, c)
}

func (h *Handler) UpdateCompanyHandler(ctx *gin.Context) {
	var c company.Company
	if err := ctx.BindJSON(&c); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "UpdateCompanyHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&c); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "UpdateCompanyHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := cast.ToInt64(ctx.Param("id"))
	c.ID = id

	c, err := h.company.UpdateCompany(ctx, c)
	if err != nil {
		slog.Error("Failed to update company", "err", err, "function", "UpdateCompanyHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

func (h *Handler) DeleteCompanyHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	err := h.company.DeleteCompany(ctx, id)
	if err != nil {
		slog.Error("Failed to delete company", "err", err, "function", "DeleteCompanyHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
