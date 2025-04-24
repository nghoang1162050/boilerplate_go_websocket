package middleware

import (
	"boilerplate_go_websocket/internal/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // If the request path is in the ignored list, simply pass to the next handler.
            if utils.ShouldIgnoreRequest(c.Path()) {
                return next(c)
            }

            tokenString, err := utils.ExtractTokenFromHeader(c.Request().Header.Get("Authorization"))
            if err != nil {
                return echo.ErrUnauthorized
            }

            secret := os.Getenv("JWT_SECRET")
            claims := &jwt.RegisteredClaims{}
            token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
                return []byte(secret), nil
            })
            
            if err != nil || !token.Valid {
                return echo.ErrUnauthorized
            }

            return next(c)
        }
    }
}
