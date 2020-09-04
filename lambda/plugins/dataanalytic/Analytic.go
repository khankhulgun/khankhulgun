package dataanalytic

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"fmt"
	"strings"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/jinzhu/gorm"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/dataanalytic/models"
)

func AnalyticsData(c echo.Context) error {

	var analytics []models.Analytic
	var AnalyticFilter []models.AnalyticFilter
	var AnalyticRangeFilter []models.AnalyticRangeFilter
	var analyticRowsColumns []models.AnalyticRowsColumn
	var AnalyticRangeRowColumn []models.AnalyticRangeRowColumn

	fmt.Println(analytics)

	DB.DB.Find(&analytics)
	DB.DB.Find(&AnalyticFilter)
	DB.DB.Find(&AnalyticRangeFilter)
	DB.DB.Find(&analyticRowsColumns)
	DB.DB.Find(&AnalyticRangeRowColumn)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"analytics":          analytics,
		"filters":            AnalyticFilter,
		"range_filters":      AnalyticRangeFilter,
		"rows_columns":       analyticRowsColumns,
		"range_rows_columns": AnalyticRangeRowColumn,
	})
}

func Pivot(c echo.Context) error {

	r := new(models.Reguest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "false",
		})
	}

	var analytic models.Analytic
	DB.DB.Where("id = ?", r.AnalyticType).Find(&analytic)

	if analytic.ID >= 1 {

		ColHeader := []string{}
		Data := []interface{}{}

		RowCond := SetPivotCondition(r.Row, *r)

		rowData := []map[string]interface{}{}

		if RowCond["type"] == "table" || RowCond["type"] == "table_range" {
			rowData = GetTableData(RowCond["table"], RowCond["dataCondition"])
		}

		Table := analytic.Source + " as b"

		rowNameField := RowCond["nameField"]
		tableTitle := RowCond["title"]
		ColHeader = append(ColHeader, tableTitle)

		ColData := []map[string]interface{}{}
		colCond := SetPivotCondition(r.Col, *r)

		if r.Col != "" {

			if colCond["type"] == "table" {
				ColData = GetTableData(colCond["table"], colCond["dataCondition"])

				colNameField := colCond["nameField"]
				//	tableTitle := RowCond["title"] + "/" + cowData["title"]

				for _, c := range ColData {

					ColHeader = append(ColHeader, fmt.Sprintf("%v", c[colNameField]))
				}
			} else if colCond["type"] == "table_range" {
				ColData = GetTableData(colCond["table"], colCond["dataCondition"])

				colStartField := colCond["start_field"]
				colEndField := colCond["end_field"]
				//	tableTitle := RowCond["title"] + "/" + cowData["title"]

				for _, c := range ColData {

					ColHeader = append(ColHeader, fmt.Sprintf("%v-%v", c[colStartField], c[colEndField]))
				}
			}
		}

		rin := 0
		for _, row := range rowData {
			arr := []interface{}{}

			if RowCond["type"] == "table_range" {
				startField := RowCond["start_field"]
				endField := RowCond["end_field"]
				arr = append(arr, fmt.Sprintf("%v-%v", row[startField], row[endField]))
			} else {
				arr = append(arr, row[rowNameField])
			}

			if r.Col != nil {
				for _, co := range ColData {

					qrc := DB.DB.Table(Table)
					qrc = AppendFilter(qrc, *r)
					compareAliasRow := "b"
					compareAliasCol := "b"

					switch comparison := RowCond["comparison"]; comparison {
					case "equal":
						rowCompare := RowCond["bCompareField"]
						rowCompareVal := RowCond["compareField"]
						qrc = qrc.Where(compareAliasRow+"."+rowCompare+" =  ?", row[rowCompareVal])

					case "between":
						rowCompare := RowCond["bCompareField"]
						rowStartVal := RowCond["start_field"]
						rowEndVal := RowCond["end_field"]
						qrc = qrc.Where(compareAliasRow+"."+rowCompare+" >=  ?", row[rowStartVal])
						qrc = qrc.Where(compareAliasRow+"."+rowCompare+" <=  ?", row[rowEndVal])
					}

					switch comparison := RowCond["comparison"]; comparison {
					case "equal":
						colCompare := colCond["bCompareField"]
						colCompareVal := colCond["compareField"]
						qrc = qrc.Where(compareAliasCol+"."+colCompare+" =  ?", co[colCompareVal])
					case "between":
						colCompare := colCond["bCompareField"]
						colStartVal := colCond["start_field"]
						colEndVal := colCond["end_field"]
						qrc = qrc.Where(compareAliasCol+"."+colCompare+" >=  ?", co[colStartVal])
						qrc = qrc.Where(compareAliasCol+"."+colCompare+" <=  ?", co[colEndVal])
					}

					if RowCond["comparison"] != "multiSum" {
						var count int
						qrc.Count(&count)
						arr = append(arr, count)
					} else {

						type Result struct {
							Value int64 `gorm:"column:value" json:"value"`
						}

						resultValue := Result{}

						qrc.Select("SUM(" + RowCond["bCompareField"] + ") as value").Group(RowCond["bCompareField"]).Scan(&resultValue)
						arr = append(arr, resultValue.Value)
					}
				}
			} else {
				qrc := DB.DB.Table(Table)
				qrc = AppendFilter(qrc, *r)
				compareAlias := "b"
				switch comparison := RowCond["comparison"]; comparison {
				case "equal":
					rowCompare := RowCond["bCompareField"]
					rowCompareVal := RowCond["compareField"]

					qrc = qrc.Where(compareAlias+"."+rowCompare+" =  ?", row[rowCompareVal])

				case "between":
					rowCompare := RowCond["bCompareField"]
					rowStartVal := RowCond["start_field"]
					rowEndVal := RowCond["end_field"]
					qrc = qrc.Where(compareAlias+"."+rowCompare+" >=  ?", row[rowStartVal])
					qrc = qrc.Where(compareAlias+"."+rowCompare+" <=  ?", row[rowEndVal])
				}

				if RowCond["comparison"] != "multiSum" {
					var count int
					qrc.Count(&count)

					arr = append(arr, count)
				} else {

					type Result struct {
						Value int64 `gorm:"column:value" json:"value"`
					}

					resultValue := Result{}

					qrc.Select("SUM(" + RowCond["bCompareField"] + ") as value").Group(RowCond["bCompareField"]).Scan(&resultValue)
					arr = append(arr, resultValue.Value)
				}
			}

			rin++
			Data = append(Data, arr)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{

			"data":   Data,
			"header": ColHeader,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "false",
		})
	}

}

func SetPivotCondition(rowColId interface{}, r models.Reguest) map[string]string {

	conData := map[string]string{}
	if rowColId != nil {
		switch r.Row.(type) {
		case string:

			var rowColumn models.AnalyticRangeRowColumn

			rowColIdNew := strings.ReplaceAll(rowColId.(string), "range_", "")

			DB.DB.Where("id = ?", rowColIdNew).Find(&rowColumn)

			if rowColumn.ID >= 1 {
				conData = map[string]string{
					"title":         rowColumn.Title,
					"type":          rowColumn.Type,
					"table":         rowColumn.SourceTable,
					"bCompareField": rowColumn.SourceCompareField,
					"comparison":    rowColumn.Comparison,
					"dataCondition": rowColumn.DataCondition,
					"start_field":   rowColumn.StartField,
					"end_field":     rowColumn.EndField,
				}
			}

		default:
			var rowColumn models.AnalyticRowsColumn

			DB.DB.Where("id = ?", rowColId).Find(&rowColumn)

			if rowColumn.ID >= 1 {
				conData = map[string]string{
					"title":         rowColumn.Title,
					"type":          rowColumn.Type,
					"table":         rowColumn.SourceTable,
					"nameField":     rowColumn.NameFeild,
					"bCompareField": rowColumn.SourceCompareField,
					"compareField":  rowColumn.CompareFeild,
					"comparison":    rowColumn.Comparison,
					"dataCondition": rowColumn.DataCondition,
				}
			}
		}

	}

	return conData
}

func AppendFilter(query *gorm.DB, r models.Reguest) *gorm.DB {

	for _, filter := range r.Filters {
		if filter.Value != "" && filter.Value != nil {
			query = query.Where(filter.FilterField+" = ?", filter.Value)
		}
	}
	for _, filter := range r.RangeFilters {
		if filter.Value[0] >= 1 || filter.Value[1] >= 1 {
			query = query.Where("b."+filter.FilterField+" >= ? and b."+filter.FilterField+" <= ?", filter.Value[0], filter.Value[1])
		}
	}

	return query
}

func GetTableData(Table string, Condition string) []map[string]interface{} {
	data := []map[string]interface{}{}

	filter := ""
	if Condition != "" {
		filter = " WHERE " + Condition
	}

	rows, _ := DB.DB.DB().Query("SELECT *  FROM " + Table + filter)

	/*start*/

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

			if ok {

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

	return data

}
