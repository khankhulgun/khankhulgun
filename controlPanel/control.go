package controlPanel
import (
	"html/template"
	"github.com/khankhulgun/khankhulgun/controlPanel/handlers"
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/agentMW"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/tools"

)

func Set(e *echo.Echo, UseNotify bool) {

	if config.Config.App.Migrate == "true"{
		//utils.AutoMigrateSeed()
	}

	/* REGISTER VIEWS */
	AbsolutePath := config.AbsolutePath()
	templates := tools.GetTemplates(e)
	templates["control.html"] = template.Must(template.ParseFiles(AbsolutePath+"controlPanel/templates/control.html"))
	templates["form.html"] = template.Must(template.ParseFiles(AbsolutePath+"controlPanel/templates/form.html"))

	/* ROUTES */
	e.GET("/control", handlers.Index(UseNotify), agentMW.IsLoggedInCookie)
	e.GET("/form/:schema_id/:id", handlers.Form)
	e.GET("/test-fcm", handlers.TestFCM, agentMW.IsLoggedInCookie)



}
