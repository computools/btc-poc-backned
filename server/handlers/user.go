package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"btc-backend/internal/user"
)

func (h *Handler) GetUserHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	u, err := h.user.GetUser(ctx, id)
	if err != nil {
		slog.Error("Failed to get user", "err", err, "function", "GetUserHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (h *Handler) GetUsersByCompanyIDHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Query("company_id"))

	users, err := h.user.GetUsersByCompanyID(ctx, id)
	if err != nil {
		slog.Error("Failed to get companies", "err", err, "function", "GetUsersByCompanyIDHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) CreateUserHandler(ctx *gin.Context) {
	var u user.CreateUser
	if err := ctx.BindJSON(&u); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "CreateUserHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&u); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "CreateUserHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	us, err := h.user.CreateUser(ctx, u)
	if err != nil {
		slog.Error("Failed to create user", "err", err, "function", "CreateUserHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, us)
}

func (h *Handler) UpdateUserHandler(ctx *gin.Context) {
	var u user.User
	if err := ctx.BindJSON(&u); err != nil {
		slog.Error("Failed to bind request body", "err", err, "function", "UpdateUserHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&u); err != nil {
		slog.Error("Failed to validate request body", "err", err, "function", "UpdateUserHandler")
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := cast.ToInt64(ctx.Param("id"))
	u.ID = id

	u, err := h.user.UpdateUser(ctx, u)
	if err != nil {
		slog.Error("Failed to update user", "err", err, "function", "UpdateUserHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (h *Handler) DeleteUserHandler(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	err := h.user.DeleteUser(ctx, id)
	if err != nil {
		slog.Error("Failed to delete user", "err", err, "function", "DeleteUserHandler")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
