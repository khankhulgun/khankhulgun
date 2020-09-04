package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/dataform"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/datagrid"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
)

func Crud(c echo.Context) error {
	schemaId := c.Param("schemaId")
	action := c.Param("action")
	id := c.Param("id")

	return dataform.Exec(c, schemaId, action, id)
}
func CheckUnique(c echo.Context) error {
	return dataform.CheckUnique(c)
}
func Upload(c echo.Context) error {

	return dataform.Upload(c)

}
func CheckCurrentPassword(c echo.Context) error {
	return utils.CheckCurrentPassword(c)
}

func Delete(c echo.Context) error {
	schemaId := c.Param("schemaId")
	id := c.Param("id")

	return datagrid.Exec(c, schemaId, "delete", id)

}
func ExportExcel(c echo.Context) error {
	schemaId := c.Param("schemaId")

	return datagrid.Exec(c, schemaId, "excel", "")

}

func dieIF(err error) {
	if err != nil {
		panic(err)
	}
}
