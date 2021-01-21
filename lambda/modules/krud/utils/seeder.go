package utils

import (
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	krudModels "github.com/khankhulgun/khankhulgun/lambda/modules/krud/models"
	arcGISModels "github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS/models"
)

func AutoMigrateSeed(UseArcGISConnection bool) {

	if config.Config.Database.Connection == "mssql"{
		DB.DB.AutoMigrate(
			&krudModels.Krud{},
			&krudModels.KrudTemplate{},
			&krudModels.CrudLogMSSQL{},
			&arcGISModels.ArcgisConnection{},
		)
	} else {
		DB.DB.AutoMigrate(
			&krudModels.Krud{},
			&krudModels.KrudTemplate{},
			&krudModels.CrudLog{},
			&arcGISModels.ArcgisConnection{},
		)
	}


	if config.Config.App.Seed == "true" {
		var vbs []krudModels.KrudTemplate
		DB.DB.Find(&vbs)

		if len(vbs) <= 0 {
			seedData()
		}
	}
}
func seedData() {
	/*KRUD TEMPLATES*/
	templates := [4]string{"canvas", "spa", "default", "default"}

	for _, template := range templates {
		newTemplate := krudModels.KrudTemplate{
			TemplateName: template,
		}

		DB.DB.Create(&newTemplate)
	}



}

