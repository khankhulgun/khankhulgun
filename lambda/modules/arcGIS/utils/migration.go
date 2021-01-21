package utils

import (

	"github.com/khankhulgun/khankhulgun/DB"
	arcGISModels "github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS/models"

)

func AutoMigrateSeed() {
	db := DB.DB

	db.AutoMigrate(
		&arcGISModels.ArcgisConnection{},
	)

}

