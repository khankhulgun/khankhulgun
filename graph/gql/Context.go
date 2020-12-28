package gql

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)
import "github.com/labstack/echo/v4"
import "context"

func EchoContextFromContext(ctx context.Context) (echo.Context, error) {
	echoContext := ctx.Value("EchoContextKey")
	if echoContext == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		return nil, err
	}

	ec, ok := echoContext.(echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		return nil, err
	}
	return ec, nil
}
 func Auth(c echo.Context) (jwt.Claims, error ) {

	 token, err := JWTFromCookie("token", c)
	 if err != nil {
		 return nil, err
	 }

	 fmt.Println(token)


	return nil, nil
}
func CheckAuth(ctx context.Context) (jwt.MapClaims, error) {
	echoContext := ctx.Value("EchoContextKey")
	if echoContext == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		return  nil, err
	}

	ec, ok := echoContext.(echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		return  nil, err
	}
	userClaims, authError := IsLoggedIn(ec)
	if authError != nil {
		return  nil, authError
	}
	user := userClaims.(jwt.MapClaims)
	return user, nil
}