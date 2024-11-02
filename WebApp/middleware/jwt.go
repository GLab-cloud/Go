package middleware

import (
	"github-trend-BE/model"
	"github-trend-BE/security"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.SECRECT_KEY),
	}

	return middleware.JWTWithConfig(config)
}