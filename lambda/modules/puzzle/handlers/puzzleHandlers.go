package handlers

import (
	"database/sql"
	"fmt"
	"strconv"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	agentUtils "github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/DBSchema"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/dataform"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/datagrid"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/datasource"
	"net/http"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/tools"
)
import "regexp"

type vb_schema struct {
	ID         int        `gorm:"column:id;primary_key" json:"id"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

func Index(c echo.Context) error {
	dbSchema := DBSchema.GetDBSchema()

	gridList := []models.VBSchema{}

	DB.DB.Where("type = ?", "grid").Find(&gridList)

	//gridList, err := models.VBSchemas(qm.Where("type = ?", "grid")).All(context.Background(), DB)
	//dieIF(err)

	User := agentUtils.AuthUserObject(c)

	return c.Render(http.StatusOK, "puzzle.html", map[string]interface{}{
		"title":                     config.Config.Title,
		"favicon":                     config.Config.Favicon,
		"app_logo":                     config.Config.Logo,
		"app_text":                     "СИСТЕМИЙН УДИРДЛАГА",
		"dbSchema":                  dbSchema,
		"gridList":                  gridList,
		"User":                      User,
		"user_fields":               config.Config.UserDataFields,
		"data_form_custom_elements": config.Config.DataFormCustomElements,
		"mix":                       utils.Mix,
	})

}

func GetVB(c echo.Context) error {

	type_ := c.Param("type")
	id := c.Param("id")
	condition := c.Param("condition")

	if id != "" {

		match, _ := regexp.MatchString("_", id)



		if match {
			VBSchema := models.VBSchemaAdmin{}

			DB.DB.Where("id = ?", id).First(&VBSchema)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": "true",
				"data":   VBSchema,
			})
		} else {
			VBSchema := models.VBSchema{}

			DB.DB.Where("id = ?", id).First(&VBSchema)


			if type_ == "form"{

				if condition != ""{
					if condition != "builder"{
						return dataform.SetCondition(condition, c, VBSchema)
					}
				}
			}

			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": "true",
				"data":   VBSchema,
			})
		}

	} else {

		VBSchemas := []models.VBSchema{}

		DB.DB.Where("type = ?", type_).Find(&VBSchemas)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "true",
			"data":   VBSchemas,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status": "false",
	})

}
func SaveVB(c echo.Context) error {

	type_ := c.Param("type")
	id := c.Param("id")
	//condition := c.Param("condition")

	vbs := new(vb_schema)
	if err := c.Bind(vbs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	}

	if id != "" {
		id_, _ := strconv.ParseUint(id, 0, 64)

		vb := models.VBSchema{}

		DB.DB.Where("id = ?", id_).First(&vb)

		vb.Name = vbs.Name
		vb.Schema = vbs.Schema
		//_, err := vb.Update(context.Background(), DB, boil.Infer())

		BeforeSave(id_, type_)

		err := DB.DB.Save(&vb).Error

		if type_ == "form" {
			//WriteModelData(id_)
			WriteModelData()
		} else if type_ == "grid" {
			WriteGridModel()
		}






		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"status": "false",
			})
		} else {
			afterStatus := AfterSave(vb, type_)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": afterStatus,
			})
		}

	} else {
		vb := models.VBSchema{
			Name:   vbs.Name,
			Schema: vbs.Schema,
			Type:   type_,
			ID:0,
		}

		//err := vb.Insert(context.Background(), DB, boil.Infer())

		DB.DB.NewRecord(vb) // => returns `true` as primary key is blank

		err := DB.DB.Create(&vb).Error

		if type_ == "form" {
			//WriteModelData(vb.ID)
			WriteModelData()
		} else if type_ == "grid" {
			WriteGridModel()
		}



		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"status": "false",
			})
		} else {
			afterStatus := AfterSave(vb, type_)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": afterStatus,
			})
		}

	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status": "false",
	})

}
func DeleteVB(c echo.Context) error {

	type_ := c.Param("type")
	id := c.Param("id")
	//condition := c.Param("condition")

	vbs := new(vb_schema)
	id_, _ := strconv.ParseUint(id, 0, 64)

	BeforeDelete(id_, type_)

	err := DB.DB.Where("id = ?", id).Where("type = ?", type_).Delete(&vbs).Error

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

func BeforeDelete(id uint64, type_ string){

	if type_ == "datasource"{
		vb := models.VBSchema{}

		DB.DB.Where("id = ?", id).First(&vb)

		datasource.DeleteView("ds_"+vb.Name)
	}

}
func BeforeSave(id uint64, type_ string){

	if type_ == "datasource"{
		vb := models.VBSchema{}

		DB.DB.Where("id = ?", id).First(&vb)

		datasource.DeleteView("ds_"+vb.Name)
	}

}
func AfterSave(vb models.VBSchema, type_ string) bool{

	if type_ == "datasource"{
		return datasource.CreateView(vb.Name, vb.Schema)
	}

	return true

}


/*GRID*/
func GridVB(c echo.Context) error {
	schemaId := c.Param("schemaId")
	action := c.Param("action")
	id := c.Param("id")

	return datagrid.Exec(c, schemaId, action, id)

}
func WriteGridModel() {

	VBSchemas := []models.VBSchema{}
	DB.DB.Where("type = ?", "grid").Find(&VBSchemas)
	DBSchema.WriteGridModel(VBSchemas)
	DBSchema.WriteGridDataCaller(VBSchemas)

}

/*FROM*/
func WriteModelData() {
	VBSchemas := []models.VBSchema{}
	DB.DB.Where("type = ?", "form").Find(&VBSchemas)
	DBSchema.WriteFormModel(VBSchemas)
	DBSchema.WriteModelCaller(VBSchemas)
	DBSchema.WriteValidationCaller(VBSchemas)
}
func GetOptions(c echo.Context) error {

	r := new(dataform.Relations)
	if err := c.Bind(r); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	}
	optionsData := map[string][]map[string]interface{}{}

	var DB_ *sql.DB
	DB_ = DB.DB.DB()
	for table, relation := range r.Relations {
		data := dataform.OptionsData(DB_, relation, c)
		optionsData[table] = data

	}
	return c.JSON(http.StatusOK, optionsData)

}

