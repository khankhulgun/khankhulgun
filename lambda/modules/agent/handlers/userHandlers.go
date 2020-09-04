package handlers

import (
	"fmt"

	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/tools"
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"

	agentModels "github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
)

func GetUsers(c echo.Context) error {
	role := c.QueryParam("role")
	sort := c.QueryParam("sort")
	direction := c.QueryParam("direction")


	users := []agentModels.User{}



	query := DB.DB.Table("users").Order(sort+" "+direction)

	fmt.Println(role)
	if role != "all"{
		query = query.Where("role = ?", role)
	}
	query = query.Where("deleted_at IS NULL")

	data := utils.Paging(&utils.Param{
		DB:      query,
		Page:    GetPage(c),
		Limit:   16,
	}, &users)
	return c.JSON(http.StatusOK, data)
}

func SearchUsers(c echo.Context) error {
	q := c.Param("q")



	users := []agentModels.User{}



	query := DB.DB.Table("users")



	if q != "="{
		query = query.Where("deleted_at IS NULL")
		query = query.Where("login LIKE ?", "%"+q+"%")
		query = query.Or("first_name LIKE ?", "%"+q+"%")
		query = query.Or("last_name LIKE ?", "%"+q+"%")
		query = query.Or("register_number LIKE ?", "%"+q+"%")
		query = query.Or("phone LIKE ?", "%"+q+"%")
	}

	data := utils.Paging(&utils.Param{
		DB:      query,
		Page:    GetPage(c),
		Limit:   16,
	}, &users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "true",
		"data": data,
	})
}
func GetDeletedUsers(c echo.Context) error {
	role := c.QueryParam("role")
	sort := c.QueryParam("sort")
	direction := c.QueryParam("direction")


	users := []agentModels.User{}



	query := DB.DB.Table("users").Order(sort+" "+direction)

	fmt.Println(role)
	if role != "all"{
		query = query.Where("role = ?", role)
	}
	query = query.Where("deleted_at IS NOT NULL")



	data := utils.Paging(&utils.Param{
		DB:      query,
		Page:    GetPage(c),
		Limit:   16,
	}, &users)
	return c.JSON(http.StatusOK, data)
}

func GetRoles(c echo.Context) error {

	roles := []agentModels.Role{}
	DB.DB.Where("id != 1").Find(&roles)
	return c.JSON(http.StatusOK, roles)
}


func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	role := new(agentModels.User)

	err := DB.DB.Where("id = ?", id).Delete(&role).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "true",
		})
	}

}
func GetPage(c echo.Context) int  {
	page := c.QueryParam("page")

	var Page_ int = 1
	if page != ""{
		Page_, _ = strconv.Atoi(page)
	}

	return Page_;
}