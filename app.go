package khankhulgun

import (
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/controlPanel"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent"
	"github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Echo         *echo.Echo
	ModuleName   string
	GetGridMODEL func(schemaId string) (interface{}, interface{}, string, string, interface{}, string)
	GetMODEL     func(schemaId string) (string, interface{})
	GetMessages  func(schemaId string) map[string][]string
	GetRules     func(schemaId string) map[string][]string
	echoWrapHandler     echo.HandlerFunc
}

func (app *App) Start() {
	app.Echo.Logger.Fatal(app.Echo.Start(config.Config.App.Port))
	defer DB.DB.Close()
}

type Settings struct {
	UseControlPanel bool
	UseNotify       bool
	UseCrudLogger   bool
	UseArcGISConnection   bool
}

var UseControlPanel = true
var UseNotify = false
var UseCrudLogger = false
var UseArcGISConnection = false

func New(moduleName string, GetGridMODEL func(schemaId string) (interface{}, interface{}, string, string, interface{}, string), GetMODEL func(schemaId string) (string, interface{}), GetMessages func(schemaId string) map[string][]string, GetRules func(schemaId string) map[string][]string, controlPanelSettings ...*Settings) *App {

	if len(controlPanelSettings) >= 1 {
		UseControlPanel = controlPanelSettings[0].UseControlPanel
		UseNotify = controlPanelSettings[0].UseNotify
		UseCrudLogger = controlPanelSettings[0].UseCrudLogger
		UseArcGISConnection = controlPanelSettings[0].UseArcGISConnection
	}
	app := &App{
		Echo:         echo.New(),
		ModuleName:   moduleName,
		GetGridMODEL: GetGridMODEL,
		GetMODEL:     GetMODEL,
		GetMessages:  GetMessages,
		GetRules:     GetRules,
		}
	app.Echo.Use(middleware.Secure())
	//app.Echo.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	TokenLookup: "header:X-XSRF-TOKEN",
	//}))
	//CORS
	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*", "http://localhost:*","http://127.0.0.1:*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Requested-With", "x-requested-with"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	krud.Set(app.Echo, app.GetGridMODEL, app.GetMODEL, app.GetMessages, app.GetRules, UseCrudLogger, UseNotify, UseArcGISConnection)
	agent.Set(app.Echo)
	puzzle.Set(app.Echo, app.ModuleName, app.GetGridMODEL)

	if UseControlPanel {
		controlPanel.Set(app.Echo, UseNotify)
	}
	if UseArcGISConnection {
		arcGIS.Set(app.Echo)
	}
	if UseNotify {
		notify.Set(app.Echo)
	}



	app.Echo.Static("/", "public")

	return app
}
