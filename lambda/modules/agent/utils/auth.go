package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func AuthUser(c echo.Context) *models.User {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)


	Id := claims["id"]

	User := models.User{}

	DB.DB.Where("id = ?",Id).First(&User)



	//User.Password = ""
	return &User
}

func AuthUserObject(c echo.Context) map[string]interface{} {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	Id := claims["id"]

	rows, _ := DB.DB.DB().Query("SELECT * FROM users WHERE id = ?", Id)

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	userData := map[string]interface{}{}
	result_id := 0
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if (ok) {
				v = string(b)
			} else {
				v = val
			}
			userData[col] = v
		}


		result_id++
	}

	delete(userData, "password")

	return userData
}

func AuthUserObjectByLogin(login string) map[string]interface{} {
	userData := map[string]interface{}{}
	rows, _ := DB.DB.DB().Query("SELECT * FROM users WHERE login = ?", login)

	//fmt.Println(login)
	//fmt.Println(errorDB)

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)


	result_id := 0
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if (ok) {
				v = string(b)
			} else {
				v = val
			}
			userData[col] = v
		}


		result_id++
	}


	return userData
}
func AuthUserObjectByEmail(login string) map[string]interface{} {

	rows, _ := DB.DB.DB().Query("SELECT * FROM users WHERE email = ? limit 1", login)

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	userData := map[string]interface{}{}
	result_id := 0
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if (ok) {
				v = string(b)
			} else {
				v = val
			}
			userData[col] = v
		}


		result_id++
	}


	return userData
}
func Hash(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed), err
}
func IsSame(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}

type passwordPost struct {
	Password string   `json:"password"`
}
func CheckCurrentPassword(c echo.Context) error {

	post := new(passwordPost)
	if err := c.Bind(post); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false from json",
		})
	}

	user := AuthUser(c)


	if IsSame(post.Password, user.Password) {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "true",
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
			"msg": "Нууц үг буруу байна !!!",
		})

	}

}