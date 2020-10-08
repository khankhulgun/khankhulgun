package handlers

import (
	"encoding/json"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	agentModels "github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
	krudModels "github.com/khankhulgun/khankhulgun/lambda/modules/krud/models"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/DBSchema"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetRolesMenus(c echo.Context) error {

	roles := []agentModels.Role{}
	menus := []models.VBSchema{}
	kruds := []krudModels.Krud{}

	DB.DB.Where("id != 1").Find(&roles)
	DB.DB.Find(&kruds)
	DB.DB.Where("type = 'menu'").Find(&menus)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "true",
		"roles":  roles,
		"menus":  menus,
		"cruds":  kruds,
	})
}

type Role struct {
	ID          int                    `json:"id"`
	Permissions map[string]interface{} `json:"permissions"`
	Extra       map[string]interface{} `json:"extra"`
}
type RoleNew struct {
	Description string `json:"description"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
}

//  TableName sets the insert table name for this struct type
func (v *Role) TableName() string {
	return "roles"
}

func SaveRole(c echo.Context) error {

	role := new(Role)
	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	}

	role_ := agentModels.Role{}

	DB.DB.Where("id = ?", role.ID).First(&role_)

	Extra, _ := json.Marshal(role.Extra)
	Permissions, _ := json.Marshal(role.Permissions)

	role_.Extra = string(Extra)
	role_.Permissions = string(Permissions)

	err := DB.DB.Save(&role_).Error

	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error":err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	}
}

func CreateRole(c echo.Context) error {

	role_ := new(RoleNew)

	if err := c.Bind(role_); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"errer": err.Error(),
		})
	}

	role := agentModels.Role{}
	role.Description = role_.Description
	role.DisplayName = role_.DisplayName
	role.Name = role_.Name

	DB.DB.NewRecord(role)
	err := DB.DB.Create(&role).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	}
}

func UpdateRole(c echo.Context) error {
	id := c.Param("id")
	role_ := new(RoleNew)

	if err := c.Bind(role_); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": err.Error(),
		})
	}

	role := agentModels.Role{}

	DB.DB.Where("id = ?", id).First(&role)
	role.Description = role_.Description
	role.DisplayName = role_.DisplayName
	role.Name = role_.Name

	err := DB.DB.Save(&role).Error
	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "true",
		})
	}
}

func DeleteRole(c echo.Context) error {
	id := c.Param("id")
	role := new(agentModels.Role)

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

func GetKrudFields(c echo.Context) error {
	id := c.Param("id")
	krud := krudModels.Krud{}
	form := models.VBSchema{}
	grid := models.VBSchema{}


	DB.DB.Where("id = ?", id).Find(&krud)
	DB.DB.Where("id = ?", krud.Form).Find(&form)
	DB.DB.Where("id = ?", krud.Grid).Find(&grid)

	var schema DBSchema.SCHEMA
	var gridSchema DBSchema.SCHEMAGRID

	json.Unmarshal([]byte(form.Schema), &schema)
	json.Unmarshal([]byte(grid.Schema), &gridSchema)

	formFields := []string{}
	gridForm := []string{}


	for _, field := range schema.Schema {
		formFields = append(formFields, field.Model)
	}
	for _, field := range gridSchema.Schema {
		gridForm = append(gridForm, field.Model)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "true",
		"user_fields": config.Config.UserDataFields,
		"form_fields": formFields,
		"grid_fields": gridForm,
	})

}