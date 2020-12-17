package krudMW

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud/models"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify/handlers"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"strconv"
)

func CrudLogger(next echo.HandlerFunc, useNotify bool) echo.HandlerFunc {
	return func(c echo.Context) (err error) {


		res := c.Response()
		action := c.Param("action")
		if(res.Status == 200 && action != "options"){


			if(action == "" && c.Path() == "/lambda/krud/delete/:schemaId/:id"){
				action = "delete"
			}



			c.Response().After(func() {

				req := c.Request()
				user := c.Get("user").(*jwt.Token)
				claims := user.Claims.(jwt.MapClaims)
				userID := claims["id"].(float64)
				schemaId, _ := strconv.ParseInt(c.Param("schemaId"), 10, 64)

				RowId := c.Param("id")


				var bodyBytes []byte

				bodyBytes, _ = ioutil.ReadAll(req.Body)

				Log := models.CrudLog{
					UserId: int64(userID),
					Ip: c.RealIP(),
					UserAgent: req.UserAgent(),
					Action: action,
					SchemaId: schemaId,
					RowId: RowId,
					Input: string(bodyBytes),
				}

				if(action == "store"){



					fmt.Println("HIHI")

				}



				DB.DB.Create(&Log)

				if(useNotify){
					if(action == "store" || action == "update" || action == "delete"){
						handlers.BuildNotification(bodyBytes, schemaId, action, int64(userID))
					}
				}

			})





			return next(c)
			// restore body bytes
		//	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))


		} else {
			return next(c)
		}

	}
}

