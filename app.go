package khankhulgun

import (
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/controlPanel"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App KhanKhulgun application.
type App struct {
	Echo         *echo.Echo
	ModuleName   string
	GetGridMODEL func(schema_id string) (interface{}, interface{}, string, string, interface{}, string)
	GetMODEL     func(schema_id string) (string, interface{})
	GetMessages  func(schema_id string) map[string][]string
	GetRules     func(schema_id string) map[string][]string
}

func (app *App) Start() {
	app.Echo.Logger.Fatal(app.Echo.Start(config.Config.App.Port))
	defer DB.DB.Close()
}

type Settings struct {
	UseControlPanel bool
	UseNotify       bool
	UseCrudLogger   bool
}

var UseControlPanel = true
var UseNotify = false
var UseCrudLogger = false

func New(moduleName string, GetGridMODEL func(schema_id string) (interface{}, interface{}, string, string, interface{}, string), GetMODEL func(schema_id string) (string, interface{}), GetMessages func(schema_id string) map[string][]string, GetRules func(schema_id string) map[string][]string, controlPanelSettings ...*Settings) *App {

	if len(controlPanelSettings) >= 1 {
		UseControlPanel = controlPanelSettings[0].UseControlPanel
		UseNotify = controlPanelSettings[0].UseNotify
		UseCrudLogger = controlPanelSettings[0].UseCrudLogger
	}
	app := &App{
		Echo:         echo.New(),
		ModuleName:   moduleName,
		GetGridMODEL: GetGridMODEL,
		GetMODEL:     GetMODEL,
		GetMessages:  GetMessages,
		GetRules:     GetRules,
	}
	krud.Set(app.Echo, app.GetGridMODEL, app.GetMODEL, app.GetMessages, app.GetRules, UseCrudLogger, UseNotify)
	agent.Set(app.Echo)
	puzzle.Set(app.Echo, app.ModuleName, app.GetGridMODEL)

	if UseControlPanel {
		controlPanel.Set(app.Echo, UseNotify)
	}
	if UseNotify {
		notify.Set(app.Echo)
	}

	app.Echo.Use(middleware.Secure())
	app.Echo.Use(middleware.CSRF())

	//CORS
	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*", "http://localhost:*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Requested-With", "x-requested-with"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	app.Echo.Static("/", "public")

	return app
}
