package gql

import (
	"fmt"
	"strconv"
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
func CheckAuth(ctx context.Context, roles []int) (jwt.MapClaims, error) {
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
	if(len(roles) >= 1){
		userRole := GetRole(user["role"])
		for _, role :=range roles{
			if(role ==userRole){
				return user, nil
			}
		}
		return user, fmt.Errorf("Permission denied by User role")
	} else {
		return user, nil
	}

}
func GetRole(role interface{}) int{
	statusID := 1

	switch v := role.(type) {
	case int:
		statusID = role.(int)
	case float64:
		statusID = int(role.(float64))
	case float32:
		statusID = int(role.(float32))
	case string:
		i, _ := strconv.Atoi(role.(string))
		statusID = i
	default:
		fmt.Println(v)
	}
	return statusID
}