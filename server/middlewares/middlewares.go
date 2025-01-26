package middlewares

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CorsMiddleware(maxAge int) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           time.Duration(maxAge) * time.Second,
	})
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		sep := " | "

		slog.Info(
			"| " + c.Request.Method + sep +
				strings.Split(c.Request.RequestURI, "?")[0] + sep +
				cast.ToString(c.Writer.Status()) + sep +
				cast.ToString(duration) + "s",
		)
	}
}

func AuthMiddleware(auth Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateToken(c, strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "))
		if err != nil {
			slog.Error("Invalid token", "err", err, "function", "AuthMiddleware")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
