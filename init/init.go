package main

import (
	"fmt"
	"lambda/modules/puzzle/DBSchema"
	"lambda/modules/puzzle/models"
	"os"
	"github.com/otiai10/copy"
	"vp/DB"
)

func main() {

	dir, _ := os.Getwd()

	formPatch := dir+"/models/form"
	gridPatch := dir+"/models/grid"
	if _, err := os.Stat(formPatch); os.IsNotExist(err) {
		os.MkdirAll(formPatch, 0755)
	} else {
		os.RemoveAll(formPatch)
		os.MkdirAll(formPatch, 0755)
	}
	if _, err := os.Stat(gridPatch); os.IsNotExist(err) {
		os.MkdirAll(gridPatch, 0755)
	} else {
		os.RemoveAll(gridPatch)
		os.MkdirAll(gridPatch, 0755)
	}

	copy.Copy(dir+"/lambda/plugins/dataform/models", dir+"/models/form")
	copy.Copy(dir+"/lambda/plugins/datagrid/models", dir+"/models/grid")

	WriteGridModel()
	WriteModelData()
	fmt.Println("LAMBDA INIT DONE")
}

/*GRID*/
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