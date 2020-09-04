package agentMW

import (

	"github.com/labstack/echo/v4"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"

	"github.com/khankhulgun/khankhulgun/tools"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(utils.Config.JWT.Secret),
})
var IsLoggedInCookie = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey:  []byte(utils.Config.JWT.Secret),
	TokenLookup: "cookie:token",
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		role := claims["role"]

		if role != 1.0 {
			return echo.ErrUnauthorized
		}


		return next(c)
	}
}
