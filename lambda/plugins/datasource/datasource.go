package datasource

import (
	"encoding/json"
	"github.com/khankhulgun/khankhulgun/DB"
)

func DeleteView(viewName string) {
	DB.DB.Exec("DROP VIEW IF EXISTS "+viewName)
}

func CreateView(viewName string, schema string) error{

	schemaData := make(map[string]interface{})
	json.Unmarshal([]byte(schema), &schemaData)
	DeleteView("ds_"+viewName)
	query := "CREATE VIEW ds_"+viewName+" as "+schemaData["query"].(string)
	err := DB.DB.Exec(query).Error
	return err
}
