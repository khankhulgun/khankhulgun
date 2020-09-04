package dbToStruct

import (
	"database/sql"
	"fmt"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/volatiletech/sqlboiler/strmangle"
)
func GetStruct(table string) {

	if(table != ""){
		var DB_ *sql.DB
		DB_ = DB.DB.DB()
		hiddenColumns := []string{}
		columnDataTypes, err := GetColumnsFromSQLlTable(DB_, table, hiddenColumns)
		//fmt.Println(columnDataTypes)

		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())
		}

		struc_, _ := GenerateOnlyStruct(*columnDataTypes, table, strmangle.TitleCase(strmangle.Singular(table)), "models", true, true, true, "", "")
		fmt.Println(string(struc_))
	}


}
