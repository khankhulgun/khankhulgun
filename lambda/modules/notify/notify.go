package notify

import (
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/agentMW"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify/handlers"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify/utils"
	"github.com/labstack/echo/v4"
)

func Set(e *echo.Echo) {
	if config.Config.App.Migrate == "true"{
		utils.AutoMigrateSeed()
	}

	g :=e.Group("/github.com/khankhulgun/khankhulgun/lambda/notify")
	/* ROUTES */
	g.GET("/new/:user_id", handlers.GetNewNotifications, agentMW.IsLoggedInCookie)
	g.GET("/all/:user_id", handlers.GetAllNotifications, agentMW.IsLoggedInCookie)
	g.GET("/seen/:id", handlers.SetSeen, agentMW.IsLoggedInCookie)
	g.GET("/token/:user_id/:token", handlers.SetToken, agentMW.IsLoggedInCookie)
	g.GET("/token", handlers.SetTokenUrlParam)
	g.GET("/fcm", handlers.Fcm)

}


