package krudMW

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	agentModels "github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
	"github.com/khankhulgun/khankhulgun/DB"

	"github.com/dgrijalva/jwt-go"
)

type PermissionData struct {
	C      bool   `json:"c"`
	D      bool   `json:"d"`
	MenuID string `json:"menu_id"`
	R      bool   `json:"r"`
	Show   bool   `json:"show"`
	Title  string `json:"title"`
	U      bool   `json:"u"`
}
type Permissions struct {
	DefaultMenu string `json:"default_menu"`
	Extra       struct {
		Chart       bool `json:"chart"`
		Datasourcce bool `json:"datasourcce"`
		Datasource  bool `json:"datasource"`
		Moqup       bool `json:"moqup"`
	} `json:"extra"`
	MenuID      int                       `json:"menu_id"`
	Permissions map[string]PermissionData `json:"permissions"`
}

func PermissionEdit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		page_id := c.QueryParam("page_id")
		action := c.QueryParam("action")
		if page_id != "" {
			editPermission := GetPermission(c)

			if(action == "edit"){
				if editPermission.R {
					return next(c)
				} else {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error":  "Засах эрх олгогдоогүй байна",
						"status": false,
					})
				}
			}
			if(action == "update"){
				if editPermission.U {
					return next(c)
				} else {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error":  "Засах эрх олгогдоогүй байна",
						"status": false,
					})
				}
			}

		}
		return next(c)
	}
}
func PermissionCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		page_id := c.QueryParam("page_id")
		if page_id != "" {
			editPermission := GetPermission(c)
			if editPermission.C {
				return next(c)
			} else {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error":  "Нэмэх эрх олгогдоогүй байна",
					"status": false,
				})
			}
		}
		return next(c)
	}
}
func PermissionDelete(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		page_id := c.QueryParam("page_id")
		if page_id != "" {
			editPermission := GetPermission(c)

			if editPermission.D {
				return next(c)
			} else {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error":  "Устгах эрх олгогдоогүй байна",
					"status": false,
				})
			}
		}
		return next(c)
	}
}

func GetPermission(c echo.Context) PermissionData {

	page_id := c.QueryParam("page_id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"]

	Role := agentModels.Role{}
	DB.DB.Where("id = ?", role).Find(&Role)
	Permissions_ := Permissions{}
	json.Unmarshal([]byte(Role.Permissions), &Permissions_)

	return Permissions_.Permissions[page_id]

}
