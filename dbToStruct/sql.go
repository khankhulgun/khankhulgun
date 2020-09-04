package dbToStruct

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"github.com/khankhulgun/khankhulgun/config"
	"sort"
	"strings"
)

func GetColumnsFromSQLlTable(db *sql.DB, dbTable string, hiddenColumns []string) (*map[string]map[string]string, error) {

	// Store colum as map of maps
	columnDataTypes := make(map[string]map[string]string)
	// Select columnd data from INFORMATION_SCHEMA

	var pkColumn models.PKColumn

	columnDataTypeQuery := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '" + dbTable+"' AND table_schema = '" + config.Config.Database.Database+"'"

	if config.Config.Database.Connection == "mssql"{

		DB.DB.Raw("SELECT COLUMN_NAME FROM "+config.Config.Database.Database+".INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE TABLE_NAME LIKE '"+dbTable+"' AND CONSTRAINT_NAME LIKE '%PK%'").Scan(&pkColumn)

		columnDataTypeQuery = "SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE FROM "+config.Config.Database.Database+".INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '" + dbTable+"'"
	}

	if Debug {
		fmt.Println("running: " + columnDataTypeQuery)
	}

	rows, err := db.Query(columnDataTypeQuery)

	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return nil, err
	}
	if rows != nil {
		defer rows.Close()
	} else {
		return nil, errors.New("No results returned for table")
	}

	for rows.Next() {
		var column string
		var columnKey string
		var dataType string
		var nullable string
		if config.Config.Database.Connection == "mssql" {
			rows.Scan(&column, &dataType, &nullable)
		} else {
			rows.Scan(&column, &columnKey, &dataType, &nullable)
		}

		var isHidden bool = false

		for _, hiddenColumn := range hiddenColumns{
			if hiddenColumn == column{
				isHidden = true
			}
		}
		if isHidden == false{
			if config.Config.Database.Connection == "mssql" {
				if pkColumn.ColumnName == column{
					columnKey = "PRI"
				}
			}


			columnDataTypes[column] = map[string]string{"value": dataType, "nullable": nullable, "primary": columnKey}
		}


	}

	return &columnDataTypes, err
}

func GetOnlyOneField(db *sql.DB, dbTable string, oneField string) (*map[string]map[string]string, error) {


	// Store colum as map of maps
	columnDataTypes := make(map[string]map[string]string)
	// Select columnd data from INFORMATION_SCHEMA
	var pkColumn models.PKColumn

	columnDataTypeQuery := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '" + dbTable+"' AND table_schema = '" + config.Config.Database.Database+"'"

	if config.Config.Database.Connection == "mssql"{

		DB.DB.Raw("SELECT COLUMN_NAME FROM "+config.Config.Database.Database+".INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE TABLE_NAME LIKE '"+dbTable+"' AND CONSTRAINT_NAME LIKE 'PK%'").Scan(&pkColumn)

		columnDataTypeQuery = "SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE FROM "+config.Config.Database.Database+".INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '" + dbTable+"'"
	}

	if Debug {
		fmt.Println("running: " + columnDataTypeQuery)
	}

	rows, err := db.Query(columnDataTypeQuery)

	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return nil, err
	}
	if rows != nil {
		defer rows.Close()
	} else {
		return nil, errors.New("No results returned for table")
	}

	for rows.Next() {
		var column string
		var columnKey string
		var dataType string
		var nullable string
		if config.Config.Database.Connection == "mssql" {
			rows.Scan(&column, &dataType, &nullable)
		} else {
			rows.Scan(&column, &columnKey, &dataType, &nullable)
		}


		if config.Config.Database.Connection == "mssql" {
			if pkColumn.ColumnName == column{
				columnKey = "PRI"
			}
		}

	//	if oneField == column {
			columnDataTypes[column] = map[string]string{"value": dataType, "nullable": nullable, "primary": columnKey}
	//	}


	}

	return &columnDataTypes, err
}

// Generate go struct entries for a map[string]interface{} structure
func generateMysqlTypes(obj map[string]map[string]string, depth int, jsonAnnotation bool, gormAnnotation bool, gureguTypes bool) (string, bool) {

	structure := "struct {"
	time_found := false
	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		//fmt.Println(key)
		mysqlType := obj[key]
		nullable := false
		if mysqlType["nullable"] == "YES" {
			nullable = true
		}
		if mysqlType["value"] == "timestamp" || mysqlType["value"] == "datetime" || mysqlType["value"] == "date"  || mysqlType["value"] == "year"  || mysqlType["value"] == "time"{

			//if key == "created_at" ||  key == "updated_at" ||  key == "deleted_at"{
				time_found = true
			//}
			//else {
			//	mysqlType["value"] = "text"
			//}

		}

		primary := ""
		if mysqlType["primary"] == "PRI" {
			primary = ";primary_key"
			//primary = ""
		}

		// Get the corresponding go value type for this mysql type
		var valueType string
		// If the guregu (https://github.com/guregu/null) CLI option is passed use its types, otherwise use go's sql.NullX

		valueType = sqlTypeToGoType(mysqlType["value"], nullable, gureguTypes)

		fieldName := fmtFieldName(stringifyFirstChar(key))
		var annotations []string
		if gormAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s%s\"", key, primary))
		}
		if jsonAnnotation == true {
			//annotations = append(annotations, fmt.Sprintf("json:\"%s%s\"", key, primary))
			annotations = append(annotations, fmt.Sprintf("json:\"%s%s\"", key, ""))
		}
		if len(annotations) > 0 {
			structure += fmt.Sprintf("\n%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			structure += fmt.Sprintf("\n%s %s",
				fieldName,
				valueType)
		}
	}

	return structure, time_found
}

// sqlTypeToGoType converts the mysql types to go compatible sql.Nullable (https://golang.org/pkg/database/sql/) types
func sqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt
	case "bigint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt64
	case "char", "enum", "varchar", "nvarchar", "longtext", "mediumtext", "text", "ntext",  "tinytext", "geometry":
		if nullable {
			if gureguTypes {
				return gureguNullString
			}
			return sqlNullString
		}
		return "string"
	case "date", "datetime", "time", "timestamp", "datetimeoffset":
		if nullable && gureguTypes {
			return gureguNullTime
		}
		return golangTime
	case "decimal", "double", "numeric":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat64
	case "float":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}