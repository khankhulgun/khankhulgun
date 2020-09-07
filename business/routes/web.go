package routes

import (
	echo "github.com/labstack/echo/v4"
	"project/business/handlers"
	echoview "github.com/foolin/goview/supports/echoview-v4"
	"github.com/foolin/goview"
	"html/template"
	"vp/utils"
)

func Web(e *echo.Echo) {

	mix := utils.FrontMix("catalog/manifest.json")
	viewMiddleware := echoview.NewMiddleware(goview.Config{
		Root:      "_projects/catalog/business/templates", //template root path
		Extension: ".html",
		Funcs: template.FuncMap{
			"data": handlers.HomeData,
			"mix": func(index string) string {
				return utils.CallMix(index, mix)
			},
		},
	})
	//homepage
	e.GET("/", handlers.HomeProduction, viewMiddleware)     //undsen
	e.GET("/dev", handlers.HomeDevelopment, viewMiddleware) // hogjuulelt
}
