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

func TableToStruct(table string, hiddenColumns []string, pkgName string, Subs []string) string{

	if(table != ""){
		var DB_ *sql.DB
		DB_ = DB.DB.DB()

		columnDataTypes, err := GetColumnsFromSQLlTable(DB_, table, hiddenColumns)
		//fmt.Println(columnDataTypes)

		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())
		}



		subStchemas := ""

		for _, sub := range Subs{
			subStchemas = subStchemas+"\n    "+strmangle.TitleCase(strmangle.Singular(sub))+" []*"+strmangle.TitleCase(strmangle.Singular(sub))+""
		}

		struc_, _ := GenerateWithImports("", *columnDataTypes, table, strmangle.TitleCase(strmangle.Singular(table)), pkgName, true, true, true, subStchemas, "")


		return string(struc_)
	}

	return ""

}

func TableToStructNoTime(table string, hiddenColumns []string, pkgName string) string{

	if(table != ""){
		var DB_ *sql.DB
		DB_ = DB.DB.DB()

		columnDataTypes, err := GetColumnsFromSQLlTable(DB_, table, hiddenColumns)
		//fmt.Println(columnDataTypes)

		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())
		}

		struc_, _ := GenerateWithImportsNoTime("", *columnDataTypes, table, strmangle.TitleCase(strmangle.Singular(table)), pkgName, true, true, true, "", "")


		return string(struc_)
	}

	return ""

}

func TableToGraphqlOrderBy(table string, hiddenColumns []string) string{

	if(table != ""){
		var DB_ *sql.DB
		DB_ = DB.DB.DB()

		columnDataTypes, err := GetColumnsFromSQLlTable(DB_, table, hiddenColumns)
		//fmt.Println(columnDataTypes)

		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())
		}

		struc_, _ := GenerateGrapqlOrder(*columnDataTypes, table, strmangle.TitleCase(strmangle.Singular(table))+"OrderBy", "", false, false, true, "", "")

		return string(struc_)
	}

	return ""

}


func TableToGraphql(table string, hiddenColumns []string, Subs []string, isInpute bool) string{

	if(table != ""){
		var DB_ *sql.DB
		DB_ = DB.DB.DB()

		columnDataTypes, err := GetColumnsFromSQLlTable(DB_, table, hiddenColumns)
		//fmt.Println(columnDataTypes)

		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())
		}

		struc_, _ := GenerateGrapql(*columnDataTypes, table, strmangle.TitleCase(strmangle.Singular(table)), "", false, false, true, "", "", Subs, isInpute)

		return string(struc_)
	}

	return ""

}

func TableColumns(table string, hiddenColumns []string) string{

	if(table != ""){
		var DB_ *sql.DB
		DB_ = DB.DB.DB()

		columns, err := GetColumns(DB_, table, hiddenColumns)
		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())
		}
		return columns
	}

	return ""

}
