package caller

import (
 "github.com/khankhulgun/khankhulgun/models/grid"
)
func GetMODEL(schema_id string) (interface{},interface{}, string, string, interface{}, string) {

switch schema_id {

 case "crud_grid":
return new(grid.KrudGrid), new([]grid.KrudGrid), "krud", "Крүд тохиргоо", new(grid.KrudGrid), "id"

 case "analytic_grid":
return new(grid.AnalyticGrid), new([]grid.AnalyticGrid), "analytic", "АНАЛИЗ", new(grid.AnalyticGrid), "id"

 case "menu_grid":
return new(grid.MenuGrid), new([]grid.MenuGrid), "vb_schemas", "Цэсний тохиргоо", new(grid.KrudGrid), "id"

} 
return new(interface{}), new([]interface{}), "", "", new([]interface{}), "id"

}