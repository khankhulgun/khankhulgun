package handlers

import (
	"encoding/json"
	notify "github.com/khankhulgun/khankhulgun/lambda/modules/notify/handlers"
	notifyModels "github.com/khankhulgun/khankhulgun/lambda/modules/notify/models"
	agentUtils "github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	"github.com/khankhulgun/khankhulgun/tools"
	"github.com/khankhulgun/khankhulgun/DB"
	agentModels "github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
	krudModels "github.com/khankhulgun/khankhulgun/lambda/modules/krud/models"
	puzzleModels "github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
)


type Permissions struct {
	DefaultMenu string `json:"default_menu"`
	Extra       interface{} `json:"extra"`
	MenuID      int `json:"menu_id"`
	Permissions interface{} `json:"permissions"`
}



func TestFCM(c echo.Context) error {

	users := []int{2, 1}

	FCMData := notifyModels.FCMData{
		Title:"Нэр",
		Body:"Баталгаажуулна уу",
		Sound:"/lambda2/notification.mp3",
		Icon:"http://localhost/asc/logo.png",
		Link:"/p/db4172e3-25ba-807f-1c2b-da6a11d10f3b/d7fb539c-8813-5b66-e893-b4d0b1dd971b/9ac627de-77fe-055f-d347-4bdf63513e90",
		ClickAction:"http://localhost/control#/p/db4172e3-25ba-807f-1c2b-da6a11d10f3b/d7fb539c-8813-5b66-e893-b4d0b1dd971b/9ac627de-77fe-055f-d347-4bdf63513e90",
	}

	FCMNotification := notifyModels.FCMNotification{
		Title:"Нэр",
		Body:"Баталгаажуулна уу",
		Icon:"http://localhost/asc/logo.png",
		ClickAction:"http://localhost/control#/p/db4172e3-25ba-807f-1c2b-da6a11d10f3b/d7fb539c-8813-5b66-e893-b4d0b1dd971b/9ac627de-77fe-055f-d347-4bdf63513e90",
	}


	data := notifyModels.NotificationData{
		Users:users,
		Data:FCMData,
		Notification:FCMNotification,
	}
	notify.CreateNotification(c, data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}
func Index(c echo.Context) error {


	User := agentUtils.AuthUserObject(c)


	Role := agentModels.Role{}

	DB.DB.Where("id = ?", User["role"]).Find(&Role)

	Permissions_ := Permissions{}

	json.Unmarshal([]byte(Role.Permissions), &Permissions_)


	Menu := puzzleModels.VBSchema{}
	DB.DB.Where("id = ?",Permissions_.MenuID).Find(&Menu)
	MenuSchema := new(interface{})
	json.Unmarshal([]byte(Menu.Schema), &MenuSchema)
	Kruds := []krudModels.Krud{}
	DB.DB.Where("deleted_at IS NULL").Find(&Kruds)


	FirebaseConfig := config.Config.Notify.FirebaseConfig

	return c.Render(http.StatusOK, "control.html", map[string]interface{}{
		"title":       config.Config.Title,
		"favicon":     config.Config.Favicon,
		"logo":     config.Config.Logo,
		"permissions": Permissions_,
		"menu":        MenuSchema,
		"cruds":       Kruds,
		"User":        User,
		"data_form_custom_elements": config.Config.DataFormCustomElements,
		"firebase_config":           FirebaseConfig,
		"mix":                       tools.Mix,

	})

}

func Form(c echo.Context) error {



	schema_id := c.Param("schema_id")
	id := c.Param("id")

	return c.Render(http.StatusOK, "form.html", map[string]interface{}{
		"title":       config.Config.Title,
		"favicon":     config.Config.Favicon,
		"mix":                       tools.Mix,
		"schema_id":                       schema_id,
		"data_form_custom_elements": config.Config.DataFormCustomElements,
		"id":                       id,
	})

}