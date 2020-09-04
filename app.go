package khankhulgun

import (
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


// App KhanKhulgun application.
type App struct {
	Echo *echo.Echo
	ModuleName string
	GetGridMODEL func(schema_id string) (interface{}, interface{}, string, string, interface{}, string)
	GetMODEL func(schema_id string) (string, interface{})
	GetMessages func(schema_id string) map[string][]string
	GetRules func(schema_id string) map[string][]string
}

func (app *App) Start() {
	app.Echo.Logger.Fatal(app.Echo.Start(config.Config.App.Port))
	defer DB.DB.Close()
}

func New(moduleName string, GetGridMODEL func(schema_id string) (interface{}, interface{}, string, string, interface{}, string), GetMODEL func(schema_id string) (string, interface{}), GetMessages func(schema_id string) map[string][]string, GetRules func(schema_id string) map[string][]string) *App {

	app := &App{
		Echo:echo.New(),
		ModuleName:moduleName,
		GetGridMODEL:GetGridMODEL,
		GetMODEL:GetMODEL,
		GetMessages:GetMessages,
		GetRules:GetRules,
	}

	agent.Set(app.Echo)
	puzzle.Set(app.Echo, app.ModuleName, app.GetGridMODEL)
	krud.Set(app.Echo, app.GetGridMODEL, app.GetMODEL, app.GetMessages, app.GetRules)
	notify.Set(app.Echo)


	app.Echo.Use(middleware.Secure())

	//CORS
	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*","http://localhost:8080"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Requested-With", "x-requested-with"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))


	app.Echo.Static("/", "public")


	return app
}