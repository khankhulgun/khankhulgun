package utils

import (
	"fmt"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/DB"
	puzzleModels "github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	analyticModels "github.com/khankhulgun/khankhulgun/lambda/plugins/dataanalytic/models"
	"encoding/json"
	"os"
)

func AutoMigrateSeed() {
	db := DB.DB

	db.AutoMigrate(
		&puzzleModels.VBSchema{},
		&puzzleModels.VBSchemaAdmin{},
		&analyticModels.Analytic{},
		&analyticModels.AnalyticFilter{},
		&analyticModels.AnalyticRangeFilter{},
		&analyticModels.AnalyticRowsColumn{},
		&analyticModels.AnalyticRangeRowColumn{},
	)

	if config.Config.App.Seed == "true" {
		var vbs []puzzleModels.VBSchemaAdmin
		db.Find(&vbs)

		if len(vbs) <= 0 {
			seedData()
		}
	}
}
func seedData() {

	var vbs []puzzleModels.VBSchemaAdmin
	dataFile, err := os.Open("github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/initialData/vb_schemas_admin.json")
	defer dataFile.Close()
	if err != nil {
		fmt.Println("PUZZLE SEED ERROR")
	}
	jsonParser := json.NewDecoder(dataFile)
	err = jsonParser.Decode(&vbs)
	if err != nil {
		fmt.Println(err)
		fmt.Println("PUZZLE SEED DATA ERROR")
	}
	//fmt.Println(len(vbs))
	db := DB.DB
	for _, vb := range vbs {

		db.Create(&vb)
	}

}
