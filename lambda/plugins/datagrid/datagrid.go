package datagrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/datagrid/model"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
	"reflect"
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/tools"
	"sort"
	"strings"
	//"time"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
	agentUtils "github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
)

func Exec(c echo.Context, schemaId string, action string, id string, GetGridMODEL func(schema_id string) (interface{}, interface{}, string, string, interface{}, string)) error {

	//fmt.Println(schemaId)
	//fmt.Println(action)
	//fmt.Println(id)

	GridModel, GridModelArray, table, name, MainTableModel, Identity := GetGridMODEL(schemaId)

	switch action {
	case "data":
		return fetchData(c, GridModel, GridModelArray, table)
	case "aggergation":
		return Aggergation(c, GridModel, GridModelArray, table)
	case "delete":
		return DeleteData(c, GridModel, MainTableModel, table, id, Identity)
	case "excel":
		return exportExcel(c, GridModel, GridModelArray, table, name)
	case "update-row":
		return UpdateRow(c, GridModel, MainTableModel, table, id, Identity)
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"status": "false",
	})
}
func DeleteData(c echo.Context, GridModel interface{},  MainTableModel interface{}, table string, id string, Identity string) error {

	//fmt.Println(Identity, id, "Identity, id")
	err := DB.DB.Where(Identity+" = ?", id).Delete(MainTableModel).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	} else {
		callTrigger("afterDelete", GridModel, []map[string]interface{}{}, id)

		return c.JSON(http.StatusOK, map[string]string{
			"status": "true",
		})
	}
}

func UpdateRow(c echo.Context, GridModel interface{},  MainTableModel interface{}, table string, id string, Identity string) error {

	RowUpdateData := new(model.RowUpdateData)

	if err := c.Bind(RowUpdateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
			"error": err.Error(),
		})
	}
	if(len(RowUpdateData.Ids) >= 1 && RowUpdateData.Model != "" && RowUpdateData.Value >= 0){
		for _, id_ := range RowUpdateData.Ids{

			DB.DB.Model(MainTableModel).Where(Identity+" = ?", id_).Update(RowUpdateData.Model, RowUpdateData.Value)
		}

	}
		return c.JSON(http.StatusOK, map[string]string{
			"status": "true",
		})
}
func trim(s string, length int) string {
	var size, x int

	for i := 0; i < length && x < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[x:])
		x += size
	}

	return s[:x]
}
func exportExcel(c echo.Context, GridModel interface{}, GridModelArray interface{}, table string, namePre string) error {

	name := trim(namePre, 21)

	GetCondition := reflect.ValueOf(GridModel).MethodByName("GetCondition")
	conditionRes := GetCondition.Call([]reflect.Value{})
	condition := conditionRes[0].String()

	query := DB.DB.Table(table)

	query = Filter(c, GridModel, query)

	if len(condition) > 0 {
		query = query.Where(condition)
	}

	query.Find(GridModelArray)

	//return c.JSON(http.StatusOK, GridModelArray)

	var file *xlsx.File
	var sheet *xlsx.Sheet

	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet(name)
	if err != nil {
		fmt.Printf(err.Error())
	}
//	fmt.Println(sheet)
	/*HEADER*/

	GetColumns := reflect.ValueOf(GridModel).MethodByName("GetColumns")
	columnsPre := GetColumns.Call([]reflect.Value{})
	columns := columnsPre[0].Interface().(map[int]map[string]string)

	fmt.Println(columns)

	keys := make([]int, 0, len(columns))
	for k := range columns {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//fmt.Println(keys)
	//fmt.Println("==========================================================================================================================")
	//fmt.Println(sheet)

	headerRow := sheet.AddRow()
	for _, k := range keys {

		headerCell := headerRow.AddCell()
		headerCell.Value = columns[k]["label"]
	}
	/*HEADER*/

	rows_json, _ := json.Marshal(GridModelArray)

	var rows []map[string]interface{}
	err = json.Unmarshal(rows_json,&rows)

	for i := range rows{

		row_ := rows[i]
		
		dataRow := sheet.AddRow()

		for _, k := range keys {

			dataCell := dataRow.AddCell()
			value := fmt.Sprintf("%v", row_[columns[k]["column"]])
			if value == "<nil>"{
				value = ""
			}
			dataCell.Value =value
		}

	}



	var b bytes.Buffer
	if err := file.Write(&b); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"name": name + ".xlsx",
		"file": b.Bytes(),
	})

}

func fetchData(c echo.Context, GridModel interface{}, GridModelArray interface{}, table string) error {

	pageLimit := c.QueryParam("paginate")
	page := c.QueryParam("page")
	sort := c.QueryParam("sort")
	order := c.QueryParam("order")

	GetCondition := reflect.ValueOf(GridModel).MethodByName("GetCondition")
	conditionRes := GetCondition.Call([]reflect.Value{})
	condition := conditionRes[0].String()

	query := DB.DB.Table(table).Order(sort + " " + order)
	//DB.DB.LogMode(true)
	query = Filter(c, GridModel, query)

	if len(condition) > 0 {
		query = query.Where(condition)
	}
	query = Search(c, GridModel, query)
	var Page_ int = 1
	if page != "" {
		Page_, _ = strconv.Atoi(page)
	}
	Limit_, _ := strconv.Atoi(pageLimit)

	data := tools.Paging(&tools.Param{
		DB:    query,
		Page:  Page_,
		Limit: Limit_,
	}, GridModelArray)

	return c.JSON(http.StatusOK, data)

}

func Aggergation(c echo.Context, GridModel interface{}, GridModelArray interface{}, table string) error {

	//GetAggregations := reflect.ValueOf(GridModel).MethodByName("GetAggregations")
	//aggregationsRes := GetAggregations.Call([]reflect.Value{})
	//aggregations := aggregationsRes[0].Interface().([]map[string]string)
	//
	//
	//func (v *DSIrtsiinBurtgel524) GetAggregations() []map[string]string {
	//	//[{"column":"tetgeleg_dun","aggregation":"SUM","symbol":"₮"},{"column":"id","aggregation":"COUNT","symbol":"Нийт "}]
	//
	//	aggregations := []map[string]string{
	//	map[string]string{
	//	"column": "tetgeleg_dun",
	//	"aggregation": "SUM",
	//	"symbol": "₮",
	//},
	//	map[string]string{
	//	"column": "id",
	//	"aggregation": "COUNT",
	//	"symbol": "Нийт ",
	//},
	//}
	//	return aggregations
	//}


	GetCondition := reflect.ValueOf(GridModel).MethodByName("GetCondition")
	conditionRes := GetCondition.Call([]reflect.Value{})
	condition := conditionRes[0].String()

	query := DB.DB.Table(table)


	query = Filter(c, GridModel, query)

	if len(condition) > 0 {
		query = query.Where(condition)
	}

	query = Search(c, GridModel, query)


	GetAggergations := reflect.ValueOf(GridModel).MethodByName("GetAggergations")
	GetAggergationsRes := GetAggergations.Call([]reflect.Value{})
	aggergations := GetAggergationsRes[0].String()


	query = query.Select(aggergations)

	rows, _  := query.Rows()

	data := []interface{}{}
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	/*end*/

	for rows.Next() {

		/*start */

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		var myMap = make(map[string]interface{})
		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)

			if (ok) {

				v, error := strconv.ParseInt(string(b), 10, 64)
				if error != nil {
					stringValue := string(b)

					myMap[col] = stringValue
				} else {
					myMap[col] = v
				}

			} else {
				myMap[col] = val
			}

		}
		/*end*/


		data = append(data, myMap)

	}




	return c.JSON(http.StatusOK, data)
}


func Filter(c echo.Context, GridModel interface{}, query *gorm.DB) *gorm.DB {

	filterRaw, _ := ioutil.ReadAll(c.Request().Body)
	var filterData map[string]interface{}
	json.Unmarshal([]byte(filterRaw), &filterData)



	if len(filterData) >= 1 {

		GetFilters := reflect.ValueOf(GridModel).MethodByName("GetFilters")
		preFilters := GetFilters.Call([]reflect.Value{})
		filters := preFilters[0].Interface().(map[string]string)

		for k, v := range filterData {
			if k == "user_condition" {

				for _, userCondition := range v.([]interface{}){
					codintion := reflect.ValueOf(userCondition).Interface().(map[string]interface{})
					User := agentUtils.AuthUserObject(c)

					query = query.Where(codintion["grid_field"].(string)+" = ?", User[codintion["user_field"].(string)])
				}

			} else {
				filterType := filters[k]


				if filterType != "" {
					switch filterType {
					case "Select":
						query = query.Where(k+" = ?", v)
					case "Tag":
						query = query.Where(k+" IN (?)", v)
					case "DateRange":
						query = query.Where(k+" BETWEEN ? AND ?", reflect.ValueOf(v).Index(0).Interface().(string), reflect.ValueOf(v).Index(1).Interface().(string))
					case "DateRangeDouble":
						start := reflect.ValueOf(v).Index(0).Interface().(string)
						end := reflect.ValueOf(v).Index(1).Interface().(string)
						if start != "" && end != "" {
							query = query.Where(k+" BETWEEN ? AND ?", start, end)
						} else if start != "" && end == "" {
							query = query.Where(k+" >= ?", start)
						} else if start == "" && end != "" {
							query = query.Where(k+" <= ?", end)
						}

					default:
						switch vtype := v.(type) {
						case map[string]interface{}:
							fmt.Println(vtype)
							vmap := v.(map[string]interface{})
							switch vmap["type"] {
							case "contains":
								query = query.Where("LOWER("+k+") LIKE ?", "%"+strings.ToLower(fmt.Sprintf("%v", v))+"%")
							case "equals":
								query = query.Where(k+" = ?", fmt.Sprintf("%v", vmap["filter"]))
							case "lessThan":
								query = query.Where(k+" <= ?", fmt.Sprintf("%v", vmap["filter"]))
							case "greaterThan":
								query = query.Where(k+" >= ?", fmt.Sprintf("%v", vmap["filter"]))
							case "notContains":
								query = query.Where(k+" != ?", fmt.Sprintf("%v", vmap["filter"]))
							default:
								query = query.Where("LOWER("+k+") LIKE ?", "%"+strings.ToLower(fmt.Sprintf("%v", v))+"%")
							}
						default:
							query = query.Where("LOWER("+k+") LIKE ?", "%"+strings.ToLower(fmt.Sprintf("%v", v))+"%")
						}

					}
				}
			}

		}

	}

	return query
}

func Search(c echo.Context, GridModel interface{}, query *gorm.DB) *gorm.DB {

	search := c.QueryParam("search")

	if search != "" {

		GetColumns := reflect.ValueOf(GridModel).MethodByName("GetColumns")
		columnsPre := GetColumns.Call([]reflect.Value{})
		columns := columnsPre[0].Interface().(map[int]map[string]string)

		i := 0
		for _, c := range columns {
			if i <= 0 {
				query = query.Where(c["column"]+" LIKE ?", "%"+search+"%")
			} else {
				//query = query.Or(c+" LIKE ?", "%"+search+"%")
				//query = query.Where(c+" LIKE ?", "%"+search+"%")
			}
			i++
		}

	}

	return query
}

func callTrigger(action string, Model interface{}, data []map[string]interface{}, id string) []map[string]interface{} {



	GetTriggers := reflect.ValueOf(Model).MethodByName("GetTriggers")

	if GetTriggers.IsValid() {



		GetTriggersRes := GetTriggers.Call([]reflect.Value{})

		triggers := GetTriggersRes[0].Interface().(map[string]interface{})
		namespace := GetTriggersRes[1].Interface().(string)


		if len(triggers) <= 0 {
			return data
		}

		if namespace == "" {
			return data
		}

		switch action {
		case "afterDelete":
			Method := triggers["afterDelete"].(string)
			Struct := triggers["afterDeleteStruct"]
			return execTrigger(Method, Struct, Model, data, id)
		case "beforeFetch":
			Method := triggers["beforeFetch"].(string)
			Struct := triggers["beforeFetchStruct"]
			return execTrigger(Method, Struct, Model, data, id)
		case "beforeDelete":
			Method := triggers["beforeDelete"].(string)
			Struct := triggers["beforeDeleteStruct"]
			return execTrigger(Method, Struct, Model, data, id)
		case "beforePrint":
			Method := triggers["beforePrint"].(string)
			Struct := triggers["beforePrintStruct"]
			return execTrigger(Method, Struct, Model, data, id)

		}
		return data
	} else {

		return  data
	}


}
func execTrigger(triggerMethod string, triggerStruct interface{}, Model interface{}, data []map[string]interface{}, id string) []map[string]interface{} {


	if triggerMethod != ""{
		triggerMethod_ := reflect.ValueOf(triggerStruct).MethodByName(triggerMethod)

		if triggerMethod_.IsValid() {

			input := make([]reflect.Value, 3)
			input[0] = reflect.ValueOf(Model)
			input[1] = reflect.ValueOf(data)
			input[2] = reflect.ValueOf(id)
			triggerMethodRes := triggerMethod_.Call(input)

			return triggerMethodRes[0].Interface().([]map[string]interface{})
		}
		return data
	} else{
		return data
	}

}
