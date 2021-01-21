package utils

import (
	"encoding/json"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	arcGISModels "github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS/models"
	puzzleModels "github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"os"
	"fmt"
)

func AutoMigrateSeed() {
	db := DB.DB
	AbsolutePath := config.AbsolutePath()
	db.AutoMigrate(
		&arcGISModels.ArcgisConnection{},
	)
	var vbs2 []puzzleModels.VBSchema

	dataFile2, err2 := os.Open(AbsolutePath+"lambda/modules/arcGIS/initialData/vb_schemas.json")
	defer dataFile2.Close()
	if err2 != nil {
		fmt.Println("PUZZLE SEED ERROR")
	}
	jsonParser2 := json.NewDecoder(dataFile2)
	err := jsonParser2.Decode(&vbs2)
	if err != nil {
		fmt.Println(err)
		fmt.Println("PUZZLE SEED DATA ERROR")
	}
	//fmt.Println(len(vbs))

	for _, vb := range vbs2 {

		DB.DB.Create(&vb)

	}

}

