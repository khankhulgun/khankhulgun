package krud

import (
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/agentMW"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud/handlers"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud/krudMW"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud/utils"
	"github.com/labstack/echo/v4"
)

func Set(e *echo.Echo) {
	if config.Config.App.Migrate == "true"{
		utils.AutoMigrateSeed()
	}

	g :=e.Group("/github.com/khankhulgun/khankhulgun/lambda/krud")
	/* ROUTES */
	g.POST("/upload", handlers.Upload)
	g.OPTIONS("/upload", handlers.Upload)
	g.POST("/unique", handlers.CheckUnique)
	g.POST("/check_current_password", handlers.CheckCurrentPassword,  agentMW.IsLoggedInCookie)
	g.POST("/excel/:schemaId", handlers.ExportExcel,  agentMW.IsLoggedInCookie)
	g.POST("/:schemaId/:action", handlers.Crud, agentMW.IsLoggedInCookie, krudMW.PermissionCreate)
	//g.POST("/:schemaId/:action", handlers.Crud, krudMW.PermissionCreate)
	g.POST("/:schemaId/:action/:id", handlers.Crud, agentMW.IsLoggedInCookie, krudMW.PermissionEdit)
	g.DELETE("/delete/:schemaId/:id", handlers.Delete, agentMW.IsLoggedInCookie, krudMW.PermissionDelete)

	p :=e.Group("/github.com/khankhulgun/khankhulgun/lambda/krud-public")
	p.POST("/:schemaId/:action", handlers.Crud)
}
