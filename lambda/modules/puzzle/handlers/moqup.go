package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	"net/http"
	"github.com/khankhulgun/khankhulgun/tools"
)

func Moqup(c echo.Context) error  {
	id := c.Param("id")
	return c.Render(http.StatusOK, "moqup.html", map[string]interface{}{
		"title":                     config.Config.Title,
		"favicon":                   config.Config.Favicon,
		"id":                   id,
		"mix":                       tools.Mix,
	})
}
