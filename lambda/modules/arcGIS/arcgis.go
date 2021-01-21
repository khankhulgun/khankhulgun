package arcGIS

import (

	"github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS/utils"
	"github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS/handlers"
	//"lambda/modules/agent/agentMW"
	"github.com/labstack/echo/v4"
	vpUtils "github.com/khankhulgun/khankhulgun/config"


)

func Set(e *echo.Echo) {
	if vpUtils.Config.App.Migrate == "true"{
		utils.AutoMigrateSeed()
	}

	g :=e.Group("/arcgis")
	///* ROUTES */
	g.POST("/form-fields", handlers.FormFields)
	//g.GET("/all/:user_id", handlers.GetAllNotifications, agentMW.IsLoggedInCookie)
	//g.GET("/seen/:id", handlers.SetSeen, agentMW.IsLoggedInCookie)
	//g.GET("/token/:user_id/:token", handlers.SetToken, agentMW.IsLoggedInCookie)
	//g.GET("/fcm", handlers.Fcm)

}


