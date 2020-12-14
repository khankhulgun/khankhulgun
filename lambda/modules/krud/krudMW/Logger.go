package krudMW

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
)

func CrudLogger(next echo.HandlerFunc, useNotify bool) echo.HandlerFunc {
	return func(c echo.Context) error {


		res := c.Response()
		action := c.Param("action")
		if(res.Status == 200 && action != "options"){

			//req := c.Request()
			//user := c.Get("user").(*jwt.Token)
			//claims := user.Claims.(jwt.MapClaims)
			//userID := claims["id"].(float64)
			//schemaId, _ := strconv.ParseInt(c.Param("schemaId"), 10, 64)
			//
			//RowId := c.Param("id")
			//
			//
			//var bodyBytes []byte
			//
			//bodyBytes, _ = ioutil.ReadAll(req.Body)
			//
			//
			//if(action == "" && c.Path() == "/lambda/krud/delete/:schemaId/:id"){
			//	action = "delete"
			//}
			//
			//if(action == "store"){
			//
			//
			//	c.Response().After(func() {
			//
			//		Log := models.CrudLog{
			//			UserId: int64(userID),
			//			Ip: c.RealIP(),
			//			UserAgent: req.UserAgent(),
			//			Action: action,
			//			SchemaId: schemaId,
			//			RowId: RowId,
			//			Input: string(bodyBytes),
			//		}
			//
			//		DB.DB.Create(&Log)
			//	})
			//} else {
			//	Log := models.CrudLog{
			//		UserId: int64(userID),
			//		Ip: c.RealIP(),
			//		UserAgent: req.UserAgent(),
			//		Action: action,
			//		SchemaId: schemaId,
			//		RowId: RowId,
			//		Input: string(bodyBytes),
			//	}
			//
			//	DB.DB.Create(&Log)
			//
			//}
			//
			//
			//
			//
			//if(useNotify){
			//	if(action == "store" || action == "update" || action == "delete"){
			//		handlers.BuildNotification(bodyBytes, schemaId, action, int64(userID))
			//	}
			//}
			//
			//// restore body bytes
			//req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

				c.Response().After(func() {
					Response, _ := json.Marshal(&c.Response().Writer)
					fmt.Println(string(Response))

				})
		}

		return next(c)
	}
}
