package business

import (
	echo "github.com/labstack/echo/v4"
	"project/business/routes"
)

func Set(e *echo.Echo) {
	routes.Web(e)
	routes.Api(e)
}
