package krudMW

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud/models"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"strconv"
)

func CrudLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {


		res := c.Response()
		action := c.Param("action")
		if(res.Status == 200 && action != "options"){

			req := c.Request()
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			userID := claims["id"].(float64)
			schemaId, _ := strconv.ParseInt(c.Param("schemaId"), 10, 64)

			RowId := c.Param("id")


			var bodyBytes []byte

			bodyBytes, _ = ioutil.ReadAll(req.Body)


			if(action == "" && c.Path() == "/lambda/krud/delete/:schemaId/:id"){
				action = "delete"
			}

			Log := models.CrudLog{
				UserId: int64(userID),
				Ip: c.RealIP(),
				UserAgent: req.UserAgent(),
				Action: action,
				SchemaId: schemaId,
				RowId: RowId,
				Input: string(bodyBytes),
			}

			DB.DB.Create(&Log)

			// restore body bytes
			req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		return next(c)
	}
}
