package gql

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"errors"
	"reflect"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/labstack/echo/v4"
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
)

const (
	AlgorithmHS256 = "HS256"
)

var JWTconfig = JWTConfig{
	Skipper:       DefaultSkipper,
	SigningMethod: AlgorithmHS256,
	ContextKey:    "user",
	TokenLookup:   "header:" + echo.HeaderAuthorization,
	AuthScheme:    "Bearer",
	Claims:        jwt.MapClaims{},
}

func IsLoggedIn(c echo.Context) (jwt.Claims, error) {
	JWTconfig.SigningKey = []byte(config.Config.JWT.Secret)
	JWTconfig.keyFunc = func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != JWTconfig.SigningMethod {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		if len(JWTconfig.SigningKeys) > 0 {
			if kid, ok := t.Header["kid"].(string); ok {
				if key, ok := JWTconfig.SigningKeys[kid]; ok {
					return key, nil
				}
			}
			return nil, fmt.Errorf("unexpected jwt key id=%v", t.Header["kid"])
		}

		return JWTconfig.SigningKey, nil
	}
	auth, err := JWTFromCookie("token", c)
	if err != nil {
		authHeader, headerErr := JWTFromHeader("Authorization", JWTconfig.AuthScheme, c)
		if headerErr != nil {
			return nil, errors.New("invalid or expired jwt")
		} else {
			return ParseToken(authHeader, c)
		}
	} else {
		return ParseToken(auth, c)
	}

}
func ParseToken(auth string, c echo.Context) (jwt.Claims, error) {
	var err error = nil
	token := new(jwt.Token)
	// Issue #647, #656
	if _, ok := JWTconfig.Claims.(jwt.MapClaims); ok {
		token, err = jwt.Parse(auth, JWTconfig.keyFunc)
	} else {
		t := reflect.ValueOf(JWTconfig.Claims).Type().Elem()
		claims := reflect.New(t).Interface().(jwt.Claims)
		token, err = jwt.ParseWithClaims(auth, claims, JWTconfig.keyFunc)
	}
	if err == nil && token.Valid {
		// Store user information from token into context.
		c.Set(JWTconfig.ContextKey, token)
		if JWTconfig.SuccessHandler != nil {
			JWTconfig.SuccessHandler(c)
		}
		return token.Claims, nil
	} else {
		return nil, errors.New("invalid or expired jwt")
	}
}
func JWTFromCookie(name string, c echo.Context) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", ErrJWTMissing
	}
	return cookie.Value, nil
}
func JWTFromHeader(header string, authScheme string, c echo.Context) (string, error) {
	auth := c.Request().Header.Get(header)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}
	return "", ErrJWTMissing
}

type (
	// JWTConfig defines the config for JWT middleware.
	JWTConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper Skipper

		// BeforeFunc defines a function which is executed just before the middleware.
		BeforeFunc BeforeFunc

		// SuccessHandler defines a function which is executed for a valid token.
		SuccessHandler JWTSuccessHandler

		// ErrorHandler defines a function which is executed for an invalid token.
		// It may be used to define a custom JWT error.
		ErrorHandler JWTErrorHandler

		// ErrorHandlerWithContext is almost identical to ErrorHandler, but it's passed the current context.
		ErrorHandlerWithContext JWTErrorHandlerWithContext

		// Signing key to validate token. Used as fallback if SigningKeys has length 0.
		// Required. This or SigningKeys.
		SigningKey interface{}

		// Map of signing keys to validate token with kid field usage.
		// Required. This or SigningKey.
		SigningKeys map[string]interface{}

		// Signing method, used to check token signing method.
		// Optional. Default value HS256.
		SigningMethod string

		// Context key to store user information from the token into context.
		// Optional. Default value "user".
		ContextKey string

		// Claims are extendable claims data defining token content.
		// Optional. Default value jwt.MapClaims
		Claims jwt.Claims

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "param:<name>"
		// - "cookie:<name>"
		TokenLookup string

		// AuthScheme to be used in the Authorization header.
		// Optional. Default value "Bearer".
		AuthScheme string

		keyFunc jwt.Keyfunc
	}

	// JWTSuccessHandler defines a function which is executed for a valid token.
	JWTSuccessHandler func(echo.Context)

	// JWTErrorHandler defines a function which is executed for an invalid token.
	JWTErrorHandler func(error) error

	// JWTErrorHandlerWithContext is almost identical to JWTErrorHandler, but it's passed the current context.
	JWTErrorHandlerWithContext func(error, echo.Context) error

	jwtExtractor func(echo.Context) (string, error)
)
type (
	// Skipper defines a function to skip middleware. Returning true skips processing
	// the middleware.
	Skipper func(echo.Context) bool

	// BeforeFunc defines a function which is executed just before the middleware.
	BeforeFunc func(echo.Context)
)

func DefaultSkipper(echo.Context) bool {
	return false
}
