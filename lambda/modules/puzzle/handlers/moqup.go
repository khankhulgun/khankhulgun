package handlers

import (
	"github.com/khankhulgun/khankhulgun/lambda/config"
	"github.com/khankhulgun/khankhulgun/tools"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Moqup(c echo.Context) error  {
	id := c.Param("id")
	//csrfToken := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	csrfToken := ""
	return c.Render(http.StatusOK, "moqup.html", map[string]interface{}{
		"title":                     config.Config.Title,
		"favicon":                   config.Config.Favicon,
		"id":                   id,
		"csrfToken":                   csrfToken,
		"mix":                       tools.Mix,
	})
}
