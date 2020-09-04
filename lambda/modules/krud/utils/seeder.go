package utils

import (

	 "github.com/khankhulgun/khankhulgun/config"
	 "github.com/khankhulgun/khankhulgun/DB"
	krudModels "github.com/khankhulgun/khankhulgun/lambda/modules/krud/models"
	
)

func AutoMigrateSeed() {
	db := DB.DB

	db.AutoMigrate(
		&krudModels.Krud{},
		&krudModels.KrudTemplate{},
	)
	if config.Config.App.Seed == "true" {
		var vbs []krudModels.KrudTemplate
		db.Find(&vbs)

		if len(vbs) <= 0 {
			seedData()
		}
	}
}
func seedData() {
	/*KRUD TEMPLATES*/
	templates := [4]string{"canvas", "spa", "default", "default"}
	db := DB.DB

	for _, template := range templates {
		newTemplate := krudModels.KrudTemplate{
			TemplateName: template,
		}

		db.Create(&newTemplate)
	}

}

