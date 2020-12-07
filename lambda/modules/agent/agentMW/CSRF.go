package agentMW

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"net/http"
	"strings"
	"time"
)



func CSRF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		config := middleware.CSRFConfig{
			Skipper:      middleware.DefaultSkipper,
			TokenLength:  32,
			TokenLookup:  "header:" + echo.HeaderXCSRFToken,
			ContextKey:   "csrf",
			CookieName:   "_csrf",
			CookieMaxAge: 86400,
		}
		req := c.Request()
		k, err := c.Cookie(config.CookieName)
		token := ""

		// Generate token
		if err != nil {
			token = random.String(config.TokenLength)
		} else {
			// Reuse token
			token = k.Value
		}

		parts := strings.Split(config.TokenLookup, ":")
		extractor := csrfTokenFromHeader(parts[1])

		switch req.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
		default:
			// Validate token only for requests which are not defined as 'safe' by RFC7231
			clientToken, err := extractor(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			if !validateCSRFToken(token, clientToken) {
				return echo.NewHTTPError(http.StatusForbidden, "invalid csrf token")
			}
		}

		// Set CSRF cookie
		cookie := new(http.Cookie)
		cookie.Name = config.CookieName
		cookie.Value = token
		if config.CookiePath != "" {
			cookie.Path = config.CookiePath
		}
		if config.CookieDomain != "" {
			cookie.Domain = config.CookieDomain
		}
		cookie.Expires = time.Now().Add(time.Duration(config.CookieMaxAge) * time.Second)
		cookie.Secure = config.CookieSecure
		cookie.HttpOnly = config.CookieHTTPOnly
		c.SetCookie(cookie)

		// Store token in the context
		c.Set(config.ContextKey, token)

		// Protect clients from caching the response
		c.Response().Header().Add(echo.HeaderVary, echo.HeaderCookie)


		return next(c)
	}
}
type (


	// csrfTokenExtractor defines a function that takes `echo.Context` and returns
	// either a token or an error.
	csrfTokenExtractor func(echo.Context) (string, error)
)

func csrfTokenFromHeader(header string) csrfTokenExtractor {
	return func(c echo.Context) (string, error) {
		return c.Request().Header.Get(header), nil
	}
}
func validateCSRFToken(token, clientToken string) bool {
	return subtle.ConstantTimeCompare([]byte(token), []byte(clientToken)) == 1
}