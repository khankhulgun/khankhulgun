package utils

import (
	"encoding/json"
	"fmt"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	puzzleModels "github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	analyticModels "github.com/khankhulgun/khankhulgun/lambda/plugins/dataanalytic/models"
	"os"
)

func AutoMigrateSeed() {

	if config.Config.Database.Connection == "mssql"{
		DB.DB.AutoMigrate(
			&puzzleModels.VBSchemaMSSQL{},
			&puzzleModels.VBSchemaAdminMSSQL{},
			&analyticModels.Analytic{},
			&analyticModels.AnalyticFilter{},
			&analyticModels.AnalyticRangeFilter{},
			&analyticModels.AnalyticRowsColumn{},
			&analyticModels.AnalyticRangeRowColumn{},
		)
	} else {
		DB.DB.AutoMigrate(
			&puzzleModels.VBSchema{},
			&puzzleModels.VBSchemaAdmin{},
			&analyticModels.Analytic{},
			&analyticModels.AnalyticFilter{},
			&analyticModels.AnalyticRangeFilter{},
			&analyticModels.AnalyticRowsColumn{},
			&analyticModels.AnalyticRangeRowColumn{},
		)
	}

	if config.Config.App.Seed == "true" {
		var vbs []puzzleModels.VBSchemaAdmin
		DB.DB.Find(&vbs)

		if len(vbs) <= 0 {
			seedData()
		}
	}
}
func seedData() {

	var vbs []puzzleModels.VBSchemaAdmin
	AbsolutePath := config.AbsolutePath()
	dataFile, err := os.Open(AbsolutePath+"lambda/modules/puzzle/initialData/vb_schemas_admin.json")
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

	for _, vb := range vbs {

		DB.DB.Create(&vb)
	}

}
