package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"github.com/khankhulgun/khankhulgun/tools"
)

func Moqup(c echo.Context) error  {
	id := c.Param("id")
	csrfToken := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	return c.Render(http.StatusOK, "moqup.html", map[string]interface{}{
		"title":                     config.Config.Title,
		"favicon":                   config.Config.Favicon,
		"id":                   id,
		"csrfToken":                   csrfToken,
		"mix":                       tools.Mix,
	})
}
