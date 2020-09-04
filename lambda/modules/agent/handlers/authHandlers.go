package handlers

import (
	"encoding/json"
	config2 "github.com/khankhulgun/khankhulgun/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
	agentUtils "github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
	"github.com/khankhulgun/khankhulgun/tools"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type User struct {
	Login    string `json:"login" xml:"login" form:"login" query:"login"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}
type UserData struct {
	Id    int64
	Login    string
	Role int64
}
type jwtClaims struct {
	Id  int64 `json:"id"`
	Login string   `json:"login"`
	Role int64   `json:"role"`
	jwt.StandardClaims
}
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {

		u := new(User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusUnauthorized, models.Unauthorized{
				Error:  "Username & password required",
				Status: false,
			})
		}

		foundUser := agentUtils.AuthUserObjectByLogin(u.Login)

		if len(foundUser) == 0{

			return c.JSON(http.StatusUnauthorized, models.Unauthorized{
				Error:  "User not found",
				Status: false,
			})

		}


		//password, err := Hash(u.Password)
		//password_check1 := IsSame(password, foundUser.Password)
		if agentUtils.IsSame(u.Password, foundUser["password"].(string)) {

			// create jwt token
			token, err := createJwtToken(UserData{Id:foundUser["id"].(int64),Login:foundUser["login"].(string), Role:foundUser["role"].(int64)})
			if err != nil {
				//log.Println("Error Creating JWT token", err)
				return c.JSON(http.StatusUnauthorized, models.Unauthorized{
					Error:  "Unauthorized",
					Status: false,
				})
			}

			cookie := new(http.Cookie)
			cookie.Name = "token"
			cookie.Path = "/"
			cookie.Value = token
			cookie.Expires = time.Now().Add(time.Hour* time.Duration(config2.Config.JWT.Ttl))


			delete(foundUser, "password")

			foundUser["jwt"] = token

			c.SetCookie(cookie)
			return c.JSON(http.StatusOK, models.LoginData{
				Token:  token,
				Path:   checkRole(foundUser["role"].(int64)),
				Status: true,
				Data:  foundUser,
			})
		}

		return c.JSON(http.StatusUnauthorized, models.Unauthorized{
			Error:  "Unauthorized",
			Status: false,
		})

	}
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {

			cookie := new(http.Cookie)
			cookie.Name = "token"
			cookie.Path = "/"
			cookie.Value = ""
			cookie.Expires = time.Now()

			c.SetCookie(cookie)
			return c.JSON(http.StatusOK, map[string]string{
				"status": "true",
				"data":   "",
				"path":   "auth/login",
				"token":  "",
			})
		}


}


func LoginPage(c echo.Context) error {

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"title":         config.Config.Title,
		"favicon":       config.Config.Favicon,
		"lambda_config": config.Config,
		"mix":           tools.Mix,
	})
}

func createJwtToken(user UserData) (string, error) {
	// Set custom claims
	claims := &jwtClaims{
		user.Id,
		user.Login,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config2.Config.JWT.Ttl)).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config2.Config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
func checkRole(role int64) string {

	for _, r := range config.Config.RoleRedirects {
		if role == r.RoleID {
			return r.URL
		}
	}
	foundRole := models.Role{}
	DB.DB.Where("id = ?",role).First(&foundRole)
	if foundRole.Permissions != ""{

		Permissions := models.Permissions{}
		json.Unmarshal([]byte(foundRole.Permissions), &Permissions)
		if Permissions.DefaultMenu != ""{
			return config.Config.AppURL+Permissions.DefaultMenu
		}
	}
	return "/auth/login"
}
