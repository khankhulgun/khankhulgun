package handlers

import (
	"fmt"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"net/http"
	"github.com/khankhulgun/khankhulgun/DB"
)

func CountData(c echo.Context) (err error) {

	request := new(models.CountRequest)

	if err = c.Bind(request); err != nil {
		return
	}


	if len(request.CountFields) == 1 {

		var count int

		DB.DB.Table(request.CountFields[0].Table).Count(&count)
		return c.JSON(http.StatusOK, count)
	} else {
		return c.JSON(http.StatusOK, "0")
	}


}
func PieData(c echo.Context) (err error)  {
	request := new(models.PieRequest)

	if err = c.Bind(request); err != nil {
		return
	}


	if len(request.Value) >= 1 && len(request.Title) >= 1 {


		var columns string
		for _, col := range request.Value {
			if columns == ""{
				columns = col.Name
			} else {
				columns = columns + ", "+col.Name
			}
		}
		for _, col := range request.Title {
			if columns == ""{
				columns = col.Name
			} else {
				columns = columns + ", "+col.Name
			}
		}

		data := GetTableData(request.Value[0].Table, 		columns, "")

		return c.JSON(http.StatusOK, data)
	} else {
		return c.JSON(http.StatusOK, "[]")
	}
}
func TableData(c echo.Context) (err error)  {
	request := new(models.TableRequest)

	if err = c.Bind(request); err != nil {
		return
	}


	if len(request.Values) >= 1  {


		var columns string
		for _, col := range request.Values {
			if columns == ""{
				columns = col.Name
			} else {
				columns = columns + ", "+col.Name
			}
		}


		data := GetTableData(request.Values[0].Table, 		columns, "")

		return c.JSON(http.StatusOK, data)
	} else {
		return c.JSON(http.StatusOK, "[]")
	}
}
func LineData(c echo.Context) (err error)  {
	request := new(models.LineRequest)

	if err = c.Bind(request); err != nil {
		return
	}


	if len(request.Axis) >= 1 && len(request.Lines) >= 1 {


		var columns string
		for _, col := range request.Axis {
			if columns == ""{
				columns = col.Name
			} else {
				columns = columns + ", "+col.Name
			}
		}
		for _, col := range request.Lines {
			if columns == ""{
				columns = col.Name
			} else {
				columns = columns + ", "+col.Name
			}
		}

		data := GetTableData(request.Axis[0].Table, 		columns, "")

		return c.JSON(http.StatusOK, data)
	} else {
		return c.JSON(http.StatusOK, "[]")
	}
}



func GetTableData(Table string, Columns string, Condition string)[]map[string]interface{}  {
	data := []map[string]interface{}{}

	filter := ""
	if Condition != ""{
		filter = " WHERE "+Condition
	}

	fmt.Println("SELECT "+Columns+"  FROM " + Table + filter)
	rows, _ := DB.DB.DB().Query("SELECT "+Columns+"  FROM " + Table + filter)

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


			if (ok) {

				v, error := strconv.ParseInt(string(b), 10, 64)
				if error != nil {
					stringValue := string(b)
					fmt.Println(stringValue)

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