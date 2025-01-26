package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	tokens, err := h.auth.Login(ctx, username, password)
	if err != nil {
		slog.Error("Failed to login", "err", err, "function", "Login")
		ctx.Status(http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (h *Handler) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.Query("refresh_token")

	tokens, err := h.auth.RefreshToken(ctx, refreshToken)
	if err != nil {
		slog.Error("Failed to refresh token", "err", err, "function", "RefreshToken")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (h *Handler) Logout(ctx *gin.Context) {
	refreshToken := ctx.Query("refresh_token")

	err := h.auth.Logout(ctx, refreshToken)
	if err != nil {
		slog.Error("Failed to logout", "err", err, "function", "Logout")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
