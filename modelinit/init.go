package modelinit

import (
	"fmt"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/DBSchema"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"github.com/otiai10/copy"
	"os"
)

func ModelInit(projectPath string, moduleName string) {

	dir := projectPath
	AbsolutePath := config.AbsolutePath()

	formPatch := dir+"/models/form/"
	gridPatch := dir+"/models/grid/"
	if _, err := os.Stat(formPatch); os.IsNotExist(err) {
		os.MkdirAll("models/", 0755)
		os.MkdirAll(formPatch, 0755)
		os.MkdirAll("models/form/validationCaller/", 0755)
		os.MkdirAll("models/form/caller/", 0755)
	} else {
		os.MkdirAll("models/", 0755)
		os.RemoveAll(formPatch)
		os.MkdirAll(formPatch, 0755)
		os.MkdirAll("models/form/validationCaller/", 0755)
		os.MkdirAll("models/form/caller/", 0755)
	}
	if _, err := os.Stat(gridPatch); os.IsNotExist(err) {
		os.MkdirAll("models/", 0755)
		os.MkdirAll(gridPatch, 0755)
		os.MkdirAll("models/grid/caller", 0755)
	} else {
		os.MkdirAll("models/", 0755)
		os.RemoveAll(gridPatch)
		os.MkdirAll(gridPatch, 0755)
		os.MkdirAll("models/grid/caller", 0755)
	}

	copy.Copy(AbsolutePath+"lambda/plugins/dataform/models/", dir+"/models/form/")
	copy.Copy(AbsolutePath+"lambda/plugins/datagrid/models/", dir+"/models/grid/")

	WriteGridModel(moduleName)
	WriteModelData(moduleName)
	fmt.Println("MODEL INIT DONE")
}

/*GRID*/
func WriteGridModel(moduleName string) {

	VBSchemas := []models.VBSchema{}
	DB.DB.Where("type = ?", "grid").Find(&VBSchemas)
	DBSchema.WriteGridModel(VBSchemas)
	DBSchema.WriteGridDataCaller(VBSchemas, moduleName)
}
/*FROM*/
func WriteModelData(moduleName string) {
	VBSchemas := []models.VBSchema{}
	DB.DB.Where("type = ?", "form").Find(&VBSchemas)
	DBSchema.WriteFormModel(VBSchemas)
	DBSchema.WriteModelCaller(VBSchemas, moduleName)
	DBSchema.WriteValidationCaller(VBSchemas, moduleName)
}