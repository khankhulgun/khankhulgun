package DBSchema

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/dbToStruct"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"github.com/volatiletech/sqlboiler/strmangle"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//import "github.com/khankhulgun/khankhulgun/models"
type dbTable struct {
	tableName string
	tableType string
}
type TableMeta struct {
	Model  string `json:"model"`
	Title  string `json:"title"`
	DbType string `json:"dbType"`
	Table  string `json:"table"`
	Key    string `json:"key"`
	Extra  string `json:"extra"`
}
type VBSCHEMA struct {
	TableList []string               `json:"tableList"`
	ViewList  []string               `json:"viewList"`
	//TableMeta map[string][]TableMeta `json:"tableMeta"`
}
type FormItem struct {
	Model       string      `json:"model"`
	Title       string      `json:"title"`
	DbType      string      `json:"dbType"`
	Table       string      `json:"table,omitempty"`
	Key         string      `json:"key"`
	Extra       string      `json:"extra,omitempty"`
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Identity    string      `json:"identity"`
	Label       string      `json:"label"`
	PlaceHolder string      `json:"placeHolder"`
	Hidden      bool        `json:"hidden"`
	Disabled    bool        `json:"disabled"`
	Default     interface{} `json:"default"`
	Prefix      string      `json:"prefix"`
	Ifshowhide  string      `json:"ifshowhide"`
	Rules       []struct {
		Type string `json:"type"`
		Msg  string `json:"msg"`
	} `json:"rules"`
	HasTranslation bool   `json:"hasTranslation"`
	HasUserID      bool   `json:"hasUserId"`
	HasEquation    bool   `json:"hasEquation"`
	Equations      string `json:"equations"`
	IsGridSearch   bool   `json:"isGridSearch"`
	GridSearch     struct {
		Grid     interface{} `json:"grid"`
		Key      interface{} `json:"key"`
		Labels   interface{} `json:"labels"`
		Multiple bool        `json:"multiple"`
	} `json:"gridSearch"`
	IsFkey   bool `json:"isFkey"`
	Relation struct {
		Table              interface{}   `json:"table"`
		Key                interface{}   `json:"key"`
		Fields             []interface{} `json:"fields"`
		FilterWithUser             []interface{} `json:"filterWithUser"`
		SortField          interface{}   `json:"sortField"`
		SortOrder          string        `json:"sortOrder"`
		Multiple           bool          `json:"multiple"`
		Filter             string        `json:"filter"`
		ParentFieldOfForm  string        `json:"parentFieldOfForm"`
		ParentFieldOfTable string        `json:"parentFieldOfTable"`
	} `json:"relation,omitempty"`
	Span struct {
		Xs int `json:"xs"`
		Sm int `json:"sm"`
		Md int `json:"md"`
		Lg int `json:"lg"`
	} `json:"span"`
	Trigger        string `json:"trigger"`
	TriggerTimeout int    `json:"triggerTimeout"`
	File           struct {
		IsMultiple bool   `json:"isMultiple"`
		Count      int    `json:"count"`
		MaxSize    int    `json:"maxSize"`
		Type       string `json:"type"`
	} `json:"file,omitempty"`
	Options []interface{} `json:"options"`
	PasswordOption interface{} `json:"passwordOption"`
	GeographicOption interface{} `json:"GeographicOption"`
	EditorType interface{} `json:"editorType"`
	SchemaID string `json:"schemaID,omitempty"`

	//subForm data
	Name     string      `json:"name"`
	SubType     string      `json:"subtype"`
	Parent      string      `json:"parent"`
	FormId      uint64      `json:"formId"`
	FormType    string `json:"formType"`
	MinHeight    string `json:"min_height"`
	DisableDelete    bool `json:"disableDelete"`
	DisableCreate    bool `json:"disableCreate"`
	ShowRowNumber    bool `json:"showRowNumber"`
	UseTableType    bool `json:"useTableType"`
	TableTypeColumn    string `json:"tableTypeColumn"`
	TableTypeValue    string `json:"tableTypeValue"`
	Schema        []FormItem `json:"schema"`

}
type SCHEMA struct {
	Model         string      `json:"model"`
	Identity      string      `json:"identity"`
	Timestamp     bool        `json:"timestamp"`
	LabelPosition string      `json:"labelPosition"`
	LabelWidth    interface{} `json:"labelWidth"`
	Width         string      `json:"width"`
	Padding       int         `json:"padding"`
	Schema        []FormItem `json:"schema"`
	UI interface{} `json:"ui"`
	Formula []Formula `json:"formula"`
	Triggers struct {
		Namespace string `json:"namespace"`
		Insert    struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"insert"`
		Update struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"update"`
	} `json:"triggers"`
	SortField string `json:"sortField"`
	SordOrder string `json:"sordOrder"`
}
type Formula struct {
	Targets []struct {
		Field string `json:"field"`
		Prop  string `json:"prop"`
	} `json:"targets"`
	Template string `json:"template"`
	Form     string `json:"form"`
	Model     string `json:"model"`
}
type SCHEMAGRID struct {
	Model          string   `json:"model"`
	IsView         bool     `json:"isView"`
	Identity       string   `json:"identity"`
	Actions        []string `json:"actions"`
	ActionPosition int      `json:"actionPosition"`
	IsContextMenu  bool     `json:"isContextMenu"`
	StaticWidth    bool     `json:"staticWidth"`
	FullWidth      bool     `json:"fullWidth"`
	HasCheckbox    bool     `json:"hasCheckbox"`
	IsClient       bool     `json:"isClient"`
	Width          int      `json:"width"`
	Sort           string   `json:"sort"`
	SortOrder      string   `json:"sortOrder"`
	SoftDelete     bool     `json:"softDelete"`
	Paging         int      `json:"paging"`
	Template       int      `json:"template"`
	Schema         []struct {
		Model       string `json:"model"`
		Title       string `json:"title"`
		DbType      string `json:"dbType"`
		Table       string `json:"table"`
		Key         string `json:"key"`
		Extra       string `json:"extra"`
		Label       string `json:"label"`
		GridType    string `json:"gridType"`
		Width       int    `json:"width"`
		Hide        bool   `json:"hide"`
		Sortable    bool   `json:"sortable"`
		Printable   bool   `json:"printable"`
		Pinned      bool   `json:"pinned"`
		PinPosition string `json:"pinPosition"`
		Link        string `json:"link"`
		LinkTarget  string `json:"linkTarget"`
		Relation    struct {
			Table              interface{}   `json:"table"`
			Key                interface{}   `json:"key"`
			Fields             []interface{} `json:"fields"`
			SortField          interface{}   `json:"sortField"`
			SortOrder          string        `json:"sortOrder"`
			Multiple           bool          `json:"multiple"`
			Filter             string        `json:"filter"`
			ParentFieldOfForm  string        `json:"parentFieldOfForm"`
			ParentFieldOfTable string        `json:"parentFieldOfTable"`
		} `json:"relation"`
		Filterable bool `json:"filterable"`
		Filter     struct {
			Type             string      `json:"type"`
			Param            interface{} `json:"param"`
			ParamCompareType string      `json:"paramCompareType"`
			Default          interface{} `json:"default"`
			Relation         struct {
				Table     interface{}   `json:"table"`
				Key       interface{}   `json:"key"`
				Fields    []interface{} `json:"fields"`
				SortField interface{}   `json:"sortField"`
				SortOrder string        `json:"sortOrder"`
			} `json:"relation"`
		} `json:"filter"`
		Editable struct {
			Status       bool   `json:"status"`
			Type         string `json:"type"`
			ShouldUpdate bool   `json:"shouldUpdate"`
			ShouldPost   bool   `json:"shouldPost"`
		} `json:"editable"`
		Searchable     bool          `json:"searchable"`
		HasTranslation bool          `json:"hasTranslation"`
		Options        []interface{} `json:"options"`
	} `json:"schema"`
	Filter                    []interface{} `json:"filter"`
	Formula                   []interface{} `json:"formula"`
	Condition                 string        `json:"condition"`
	ColumnAggregations        []map[string]string `json:"columnAggregations"`
	ColumnAggregationsFormula []interface{} `json:"columnAggregationsFormula"`
	Header                    struct {
		Render    bool          `json:"render"`
		Preview   bool          `json:"preview"`
		Structure []interface{} `json:"structure"`
	} `json:"header"`
	Triggers struct {
		Namespace    string `json:"namespace"`
		BeforeFetch  string `json:"beforeFetch"`
		AfterFetch   string `json:"afterFetch"`
		BeforeDelete string `json:"beforeDelete"`
		AfterDelete  string `json:"afterDelete"`
		BeforePrint  string `json:"beforePrint"`
	} `json:"triggers"`
	Theme                string      `json:"theme"`
	FullText             bool        `json:"fullText"`
	EditableAction       interface{} `json:"editableAction"`
	EditFullRow          bool        `json:"editFullRow"`
	EditableShouldSubmit bool        `json:"editableShouldSubmit"`
	SingleClickEdit      bool        `json:"singleClickEdit"`
	FlashChanges         bool        `json:"flashChanges"`
	ColMenu              bool        `json:"colMenu"`
	ColFilterButton      bool        `json:"colFilterButton"`
	ShowGrid             bool        `json:"showGrid"`
	SordOrder            string      `json:"sordOrder"`
	MainTable            string      `json:"mainTable"`
	IsPivot              bool        `json:"isPivot"`
	IsPrint              bool        `json:"isPrint"`
	PrintSize            string      `json:"printSize"`
	IsExcel              bool        `json:"isExcel"`
	IsRefresh            bool        `json:"isRefresh"`
	IsNumbered           bool        `json:"isNumbered"`
}
type vb_schema struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

func DBConnection() *sql.DB {
	var DB_ *sql.DB
	DB_ = DB.DB.DB()

	return DB_
}
func GetDBSchema() VBSCHEMA {
	tables := Tables()

	//table_metas := make(map[string][]TableMeta, 0)



	//for _, table := range tables["tables"] {
	//	table_metas_ := TableMetas(table)
	//	table_metas[table] = table_metas_
	//}

	//for _, table := range tables["views"] {
	//	table_metas_ := TableMetas(table)
	//	table_metas[table] = table_metas_
	//}

	vb_schemas := VBSCHEMA{
		tables["tables"],
		tables["views"],
		//table_metas,
	}

	return vb_schemas
}
func Tables() map[string][]string {
	tables := make([]string, 0)
	views := make([]string, 0)

	//var dbTables []dbTable
	//DB.Raw("SHOW FULL TABLES").Scan(&dbTables)

	DB_ := DBConnection()

	if config.Config.Database.Connection == "mssql"{
		rows, _ := DB_.Query("SELECT TABLE_NAME, TABLE_TYPE FROM INFORMATION_SCHEMA.TABLES ORDER BY TABLE_NAME")

		for rows.Next() {
			var TABLE_NAME, TABLE_TYPE string
			rows.Scan(&TABLE_NAME, &TABLE_TYPE)

			if TABLE_TYPE != "VIEW" {
				tables = append(tables, TABLE_NAME)
			} else {
				views = append(views, TABLE_NAME)
			}
		}
		result := map[string][]string{}

		result["tables"] = tables
		result["views"] = views

		return result
	} else {
		rows, _ := DB_.Query("SHOW FULL TABLES")

		for rows.Next() {
			var tableName, tableType string
			rows.Scan(&tableName, &tableType)

			if tableType == "BASE TABLE" {
				tables = append(tables, tableName)
			} else {
				views = append(views, tableName)
			}
		}
		result := map[string][]string{}

		result["tables"] = tables
		result["views"] = views

		return result
	}


}
func TableMetas(tableName string) []TableMeta {
	table_metas := make([]TableMeta, 0)
	DB_ := DBConnection()

	if config.Config.Database.Connection == "mssql"{


		var pkColumn models.PKColumn
		DB.DB.Raw("SELECT COLUMN_NAME FROM "+config.Config.Database.Database+".INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE TABLE_NAME LIKE '"+tableName+"' AND CONSTRAINT_NAME LIKE '%PK%'").Scan(&pkColumn)

		table_metas_ms := []models.MSTableMata{}
		DB.DB.Raw("SELECT * FROM "+config.Config.Database.Database+".INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '" + tableName+"'").Scan(&table_metas_ms)


		for _, column := range table_metas_ms {
			key := ""
			extra := ""

			if column.ColumnName == pkColumn.ColumnName{
				key = "PRI"
				extra = "auto_increment"
			}

			dataType := column.DataType

			if column.DataType == "nvarchar" {
				dataType = "varchar(255)"
			} else if column.DataType == "ntext" {
				dataType = "text"
			}


			table_metas = append(table_metas, TableMeta{
				Model:  column.ColumnName,
				Title:  column.ColumnName,
				DbType: dataType,
				Table:  tableName,
				Key:    key,
				Extra:  extra,
			})
		}


	} else {
		columns, db_error := DB_.Query("show fields from " + tableName)

		if db_error == nil{
			for columns.Next() {
				var Field, Type, Null, Key, Default, Extra string
				columns.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)

				table_metas = append(table_metas, TableMeta{
					Model:  Field,
					Title:  Field,
					DbType: Type,
					Table:  tableName,
					Key:    Key,
					Extra:  Extra,
				})
			}
		}
	}


	return table_metas

}

/*GRID*/
func WriteGridModel(grids []models.VBSchema) {

	for _, vb := range grids {
		var schema SCHEMAGRID

		json.Unmarshal([]byte(vb.Schema), &schema)

		modelAlias := getModelAlias(schema.Model)
		MainTableAlias := getModelAlias(schema.MainTable)+"MainTable"

		modelAliasWithID := modelAlias + strconv.FormatUint(vb.ID, 10)
		MainTableAliasWithID := MainTableAlias + strconv.FormatUint(vb.ID, 10)

		DB_ := DBConnection()

		hiddenColumns := []string{}

		for _, column := range schema.Schema {
			if column.Hide == true && column.Model != schema.Identity && column.Model != "deleted_at" && column.Model != "created_at" && column.Model != "updated_at" {
				hiddenColumns = append(hiddenColumns, column.Model)
			}
		}

		columnDataTypes, err := dbToStruct.GetColumnsFromSQLlTable(DB_, schema.Model, hiddenColumns)

		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())

		}
		triggerPackage := ""

		if schema.Triggers.Namespace != ""{

			triggerPackage = "\n import \""+schema.Triggers.Namespace+"\" \n"

		}

		MainTableColumnDataTypes, _ := dbToStruct.GetOnlyOneField(DB_, schema.MainTable, schema.Identity)


		MainTableStructs, _ := dbToStruct.GenerateOnlyStruct(*MainTableColumnDataTypes, schema.MainTable, MainTableAliasWithID, "", true, true, true, "", "")


		struc, err := dbToStruct.GenerateWithImports(triggerPackage, *columnDataTypes, schema.Model, modelAliasWithID, "grid", true, true, true, "", string(MainTableStructs))

		f, err := os.Create("models/grid/" + modelAlias + strconv.FormatUint(vb.ID, 10) + ".go")
		if err != nil {
			fmt.Println(err, f)
		}

		l2, err := f.WriteString(string(struc))
		if err != nil {
			fmt.Println(err, l2)
			f.Close()

		}
		/*GRID DEFAULT CONDITION*/
		//
		//if len(schema.Condition) > 0 {
		gridCondition := `
			func (v *` + modelAliasWithID + `) GetCondition() string {
				return "` + schema.Condition + `"
			}`
		l3, err := f.WriteString(gridCondition)
		if err != nil {
			fmt.Println(err, l3)
			f.Close()

		}
		gridFilter := `
			func (v *` + modelAliasWithID + `) GetFilters() map[string]string {


				filters := map[string]string{

			`
		for i := range schema.Schema {
			if schema.Schema[i].Filterable == true {

				gridFilter = gridFilter + `
					"` + schema.Schema[i].Model + `":"` + schema.Schema[i].Filter.Type + `",
`
			}
		}

		gridFilter = gridFilter + `
			}

			return filters
		}`
		l4, err := f.WriteString(gridFilter)
		if err != nil {
			fmt.Println(err, l4)
			f.Close()

		}

		gridColumns := `
			func (v *` + modelAliasWithID + `) GetColumns() map[int]map[string]string{


			
				columns := make(map[int]map[string]string)

			`
		for i := range schema.Schema {
			if schema.Schema[i].Hide == false {

				gridColumns = gridColumns + `
					columns[`+fmt.Sprintf("%v", i)+`] = map[string]string{"column":"`  + schema.Schema[i].Model + `","label":"` + schema.Schema[i].Label + `"}
`

			}
		}

		gridColumns = gridColumns + `
			

			return columns
		}`
		l5, err := f.WriteString(gridColumns)
		if err != nil {
			fmt.Println(err, l5)
			f.Close()

		}


		if schema.Triggers.Namespace != ""{


			packageSplited := strings.Split(schema.Triggers.Namespace, "/")

			triggerPackageName := packageSplited[len(packageSplited) - 1]


			beforeFetchMethods := strings.Split(schema.Triggers.BeforeFetch, "@")

			beforeFetchMethod := `""`
			beforeFetchStruct := "new(interface{})"
			if len(beforeFetchMethods) >= 2 {
				beforeFetchMethod = `"`+beforeFetchMethods[1]+`"`
				beforeFetchStruct = `new(`+triggerPackageName+"."+beforeFetchMethods[0]+`)`
			}
			afterFetchMethods := strings.Split(schema.Triggers.AfterFetch, "@")

			afterFetchMethod := `""`
			afterFetchStruct := "new(interface{})"
			if len(afterFetchMethods) >= 2 {
				afterFetchMethod = `"`+afterFetchMethods[1]+`"`
				afterFetchStruct = `new(`+triggerPackageName+"."+afterFetchMethods[0]+`)`
			}

			beforeDeleteMethods := strings.Split(schema.Triggers.BeforeDelete, "@")

			beforeDeleteMethod := `""`
			beforeDeleteStruct := "new(interface{})"
			if len(beforeDeleteMethods) >= 2 {
				beforeDeleteMethod = `"`+beforeDeleteMethods[1]+`"`
				beforeDeleteStruct = `new(`+triggerPackageName+"."+beforeDeleteMethods[0]+`)`
			}

			afterDeleteMethods := strings.Split(schema.Triggers.AfterDelete, "@")

			afterDeleteMethod := `""`
			afterDeleteStruct := "new(interface{})"
			if len(afterDeleteMethods) >= 2 {
				afterDeleteMethod = `"`+afterDeleteMethods[1]+`"`
				afterDeleteStruct = `new(`+triggerPackageName+"."+afterDeleteMethods[0]+`)`
			}


			beforePrintMethods := strings.Split(schema.Triggers.BeforePrint, "@")

			beforePrintMethod := `""`
			beforePrintStruct := "new(interface{})"
			if len(beforePrintMethods) >= 2 {
				beforePrintMethod = `"`+beforePrintMethods[1]+`"`
				beforePrintStruct = `new(`+triggerPackageName+"."+beforePrintMethods[0]+`)`
			}



			triggerText, err := f.WriteString(`
func (a *` + modelAlias + strconv.FormatUint(vb.ID, 10) + `) GetTriggers() (map[string]interface{}, string) {

triggers :=map[string]interface{}{
				"beforeFetch":`+beforeFetchMethod+`,
				"beforeFetchStruct":`+beforeFetchStruct+`,
				"afterFetch":`+afterFetchMethod+`,
				"afterFetchStruct":`+afterFetchStruct+`,
				"beforeDelete":`+beforeDeleteMethod+`,
				"beforeDeleteStruct":`+beforeDeleteStruct+`,
				"afterDelete":`+afterDeleteMethod+`,
				"afterDeleteStruct":`+afterDeleteStruct+`,
				"beforePrint":`+beforePrintMethod+`,
				"beforePrintStruct":`+beforePrintStruct+`,
		}
		
return triggers, "`+schema.Triggers.Namespace+`"

}`)
			if err != nil {
				fmt.Println(err, triggerText)
				f.Close()

			}

		}


		/*GRID Aggergation*/

		gridAggergation := `
			func (v *` + modelAliasWithID + `) GetAggergations() string {


				aggergations := "`
		for i, aggergation := range schema.ColumnAggregations {

			if(i <= 0){
				gridAggergation = gridAggergation + `` + aggergation["aggregation"] + `(` +aggergation["column"]+ `) as `+aggergation["aggregation"] + `_` +aggergation["column"]
			} else {
				gridAggergation = gridAggergation + `, ` + aggergation["aggregation"] + `(` +aggergation["column"]+ `) as `+aggergation["aggregation"] + `_` +aggergation["column"]
			}
			
			
		}

		gridAggergation = gridAggergation + `"

			return aggergations
		}`
		lag, err := f.WriteString(gridAggergation)
		if err != nil {
			fmt.Println(err, lag)
			f.Close()

		}

		err = f.Close()
		if err != nil {
			fmt.Println(err)

		}



		//fmt.Println(vb.Name)

	}

}
func WriteGridDataCaller(forms []models.VBSchema, moduleName string) {
	//return new(models.Naiz)

	f, err := os.Create("models/grid/caller/modelCaller.go")
	if err != nil {
		fmt.Println(err)

	}
	l1, err := f.WriteString("package caller\n")
	if err != nil {
		fmt.Println(err, l1)
		f.Close()

	}
	l2, err := f.WriteString("import \""+moduleName+"/models/grid\"\n")
	if err != nil {
		fmt.Println(err, l2)
		f.Close()

	}
	l3, err := f.WriteString("func GetMODEL(schema_id string) (interface{}, interface{}, string, string, interface{}, string) {\n\nswitch schema_id {\n" + ` 

		case "crud_grid":
			return new(grid.KrudGrid), new([]grid.KrudGrid), "krud", "Крүд тохиргоо",new(grid.KrudGrid), "id"

		case "crud_log":
			return new(grid.CrudLog), new([]grid.CrudLog), "ds_crud_log", "Систем лог", new(grid.MainTableCrudLog), "id"

		case "analytic_grid":
			return new(grid.AnalyticGrid), new([]grid.AnalyticGrid), "analytic", "АНАЛИЗ", new(grid.AnalyticGrid), "id"

 		case "menu_grid":
			return new(grid.MenuGrid), new([]grid.MenuGrid), "vb_schemas", "Цэсний тохиргоо",new(grid.KrudGrid), "id"

 		case "notification_target_grid":
			return new(grid.NotificationTarget), new([]grid.NotificationTarget), "notification_targets", "Зорилтод мэдэгдэл",new(grid.NotificationTarget), "id"
 		
`)
	if err != nil {
		fmt.Println(err, l3)
		f.Close()

	}
	for _, vb := range forms {
		var schema SCHEMAGRID

		json.Unmarshal([]byte(vb.Schema), &schema)

		modelAlias := getModelAlias(schema.Model)
		mainTableAlias := getModelAlias(schema.MainTable)+"MainTable"

		l5, err := f.WriteString("\n case \"" + strconv.FormatUint(vb.ID, 10) + "\": \nreturn new(grid." + modelAlias + strconv.FormatUint(vb.ID, 10) + "), new([]grid." + modelAlias + strconv.FormatUint(vb.ID, 10) + "), \"" + schema.Model + "\", \"" + vb.Name + "\", new(grid." + mainTableAlias + strconv.FormatUint(vb.ID, 10) + "), \"" + schema.Identity + "\"\n")
		if err != nil {
			fmt.Println(err, l5)
			f.Close()

		}

	}

	//fmt.Println(schema.Model, "schema.Model")

	l4, err := f.WriteString("\n} \nreturn new([]interface{}), new([]interface{}),  \"\", \"\", new([]interface{}), \"id\"\n\n}")
	if err != nil {
		fmt.Println(err, l4)
		f.Close()

	}
	//fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)

	}
}

/*FROM*/
func WriteFormModel(grids []models.VBSchema) {

	for _, vb := range grids {
		var schema SCHEMA

		json.Unmarshal([]byte(vb.Schema), &schema)

		modelAlias := getModelAlias(schema.Model)
		DB_ := DBConnection()

		hiddenColumns := []string{}

		for _, column := range schema.Schema {
			if (column.Hidden == true && column.Default == nil && column.Label == "") || (column.Hidden == true && column.Default == "" && column.Label == "") {
				hiddenColumns = append(hiddenColumns, column.Model)
			}
		}

		columnDataTypes, err := dbToStruct.GetColumnsFromSQLlTable(DB_, schema.Model, hiddenColumns)

		//fmt.Println(columnDataTypes)

		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())

		}

		gormSubItem := ""
		gormStructs := ""
		for _, field := range schema.Schema {
			if field.FormType == "SubForm" {
				if field.SubType == "Form" {

					//Parent := getModelAlias(field.Parent)
					//subAlis := getModelAlias(field.Model)
					//subForm := subAlis+modelAlias+strconv.FormatUint(vb.ID, 10)
					////gormSubItem = gormSubItem+"\n"+subForm+"     []"+subForm+" `gorm:\"foreignkey:"+Parent+";\" json:\""+field.Model+"\"`"
					//
					//
					//subColumnDataTypes, _ :=dbToStruct.GetColumnsFromSQLlTable(DB_, field.Model, hiddenColumns)
					//
					//subStructs, _ := dbToStruct.GenerateOnlyStruct(*subColumnDataTypes, field.Model, subForm, "", true,true, true, "", "")
					//
					//gormStructs = gormStructs + string(subStructs)
				} else {
					//Parent := getModelAlias(field.Parent)
					subAlis := getModelAlias(field.Model)
					subForm := subAlis + modelAlias + strconv.FormatUint(vb.ID, 10)
					//gormSubItem = gormSubItem+"\n"+subForm+"     []"+subForm+" `gorm:\"foreignkey:"+Parent+";\" json:\""+field.Model+"\"`"

					subColumnDataTypes, _ := dbToStruct.GetColumnsFromSQLlTable(DB_, field.Model, hiddenColumns)

					subStructs, _ := dbToStruct.GenerateOnlyStruct(*subColumnDataTypes, field.Model, subForm, "", true, true, true, "", "")

					gormStructs = gormStructs + string(subStructs)
				}
			}
		}


		triggerPackage := ""

		if schema.Triggers.Namespace != ""{

			triggerPackage = "\n import \""+schema.Triggers.Namespace+"\" \n"

		}

		struc, err := dbToStruct.GenerateWithImports(triggerPackage, *columnDataTypes, schema.Model, modelAlias+strconv.FormatUint(vb.ID, 10), "form", true, true, true, gormSubItem, gormStructs)

		f, err := os.Create("models/form/" + modelAlias + strconv.FormatUint(vb.ID, 10) + ".go")
		if err != nil {
			fmt.Println(err, f)
		}

		l2, err := f.WriteString(string(struc))
		if err != nil {
			fmt.Println(err, l2)
			f.Close()

		}

		getSubText, err := f.WriteString(`func (a *` + modelAlias + strconv.FormatUint(vb.ID, 10) + `) GetSubForms() []map[string]interface{} {
	subForms := []map[string]interface{}{`)
		if err != nil {
			fmt.Println(err, getSubText)
			f.Close()

		}
		for _, field := range schema.Schema {
			if field.FormType == "SubForm" {
				if field.SubType == "Grid" {
					subAlis := getModelAlias(field.Model)


					subFormTex, err := f.WriteString("\nmap[string]interface{}{")
					if err != nil {
						fmt.Println(err, subFormTex)
						f.Close()

					}
					//subForm := subAlis+strconv.FormatUint(field.FormId, 10)
					subForm := subAlis+modelAlias+strconv.FormatUint(vb.ID, 10)
					subFormBody, err := f.WriteString(`
							"connection_field":"`+field.Parent+`",
							"tableTypeColumn":"`+field.TableTypeColumn+`",
							"tableTypeValue":"`+field.TableTypeValue+`",
							"table":"`+field.Model+`",
							"parentIdentity":"`+schema.Identity+`",
							"subIdentity":"`+field.Identity+`",
							"subForm":new([]`+subForm+`),
							"subFormModel":new(`+subForm+`),
`)
					if err != nil {
						fmt.Println(err, subFormBody)
						f.Close()
					}

					subFormTexEnd, err := f.WriteString(`
},`)
					if err != nil {
						fmt.Println(err, subFormTexEnd)
						f.Close()
					}
				} else {
					subAlis := getModelAlias(field.Model)


					subFormTex, err := f.WriteString("\nmap[string]interface{}{")
					if err != nil {
						fmt.Println(err, subFormTex)
						f.Close()

					}
					subForm := subAlis+strconv.FormatUint(field.FormId, 10)
					//subForm := subAlis+modelAlias+strconv.FormatUint(vb.ID, 10)
					subFormBody, err := f.WriteString(`
							"connection_field":"`+field.Parent+`",
							"tableTypeColumn":"`+field.TableTypeColumn+`",
							"tableTypeValue":"`+field.TableTypeValue+`",
							"table":"`+field.Model+`",
							"parentIdentity":"`+schema.Identity+`",
							"subIdentity":"`+field.Identity+`",
							"subForm":new([]`+subForm+`),
							"subFormModel":new(`+subForm+`),
`)
					if err != nil {
						fmt.Println(err, subFormBody)
						f.Close()
					}

					subFormTexEnd, err := f.WriteString(`
},`)
					if err != nil {
						fmt.Println(err, subFormTexEnd)
						f.Close()
					}
				}

			}
		}
		getSubTextEnd, err := f.WriteString(`}
	return subForms }`)
		if err != nil {
			fmt.Println(err, getSubTextEnd)
			f.Close()

		}

		if schema.Triggers.Namespace != ""{


			packageSplited := strings.Split(schema.Triggers.Namespace, "/")

			triggerPackageName := packageSplited[len(packageSplited) - 1]


			insertBeforeMethods := strings.Split(schema.Triggers.Insert.Before, "@")

			insertBeforeMethod := `""`
			insertBeforeStruct := "new(interface{})"
			if len(insertBeforeMethods) >= 2 {
				insertBeforeMethod = `"`+insertBeforeMethods[1]+`"`
				insertBeforeStruct = `new(`+triggerPackageName+"."+insertBeforeMethods[0]+`)`
			}

			insertAfterMethods := strings.Split(schema.Triggers.Insert.After, "@")



			insertAfterMethod := `""`
			insertAfterStruct := "new(interface{})"
			if len(insertAfterMethods) >= 2 {
				insertAfterMethod = `"`+insertAfterMethods[1]+`"`
				insertAfterStruct = `new(`+triggerPackageName+"."+insertAfterMethods[0]+`)`
			}

			updateBeforeMethods := strings.Split(schema.Triggers.Update.Before, "@")

			updateBeforeMethod :=`""`
			updateBeforeStruct := "new(interface{})"
			if len(updateBeforeMethods) >= 2 {
				updateBeforeMethod = `"`+updateBeforeMethods[1]+`"`
				updateBeforeStruct = `new(`+triggerPackageName+"."+updateBeforeMethods[0]+`)`
			}

			updateAfterMethods := strings.Split(schema.Triggers.Update.After, "@")

			updateAfterMethod := `""`
			updateAfterStruct := "new(interface{})"
			if len(updateAfterMethods) >= 2 {
				updateAfterMethod = `"`+updateAfterMethods[1]+`"`
				updateAfterStruct = `new(`+triggerPackageName+"."+updateAfterMethods[0]+`)`
			}


			triggerText, err := f.WriteString(`
func (a *` + modelAlias + strconv.FormatUint(vb.ID, 10) + `) GetTriggers() (map[string]map[string]interface{}, string) {

triggers :=map[string]map[string]interface{}{
			"insert":map[string]interface{}{
				"before":`+insertBeforeMethod+`,
				"beforeStruct":`+insertBeforeStruct+`,
				"after":`+insertAfterMethod+`,
				"afterStruct":`+insertAfterStruct+`,
			},
			"update":map[string]interface{}{
		        "before":`+updateBeforeMethod+`,
				"beforeStruct":`+updateBeforeStruct+`,
				"after":`+updateAfterMethod+`,
				"afterStruct":`+updateAfterStruct+`,
			},
		}
		
return triggers, "`+schema.Triggers.Namespace+`"

}`)
			if err != nil {
				fmt.Println(err, triggerText)
				f.Close()

			}

		}


		formTypes := `
			func (v *` + modelAlias + strconv.FormatUint(vb.ID, 10) + `) GetFromTypes() map[string]string{


				fields := map[string]string{

			`
		for i := range schema.Schema {


			formTypes = formTypes + `
					"` + schema.Schema[i].Model + `":"` + schema.Schema[i].FormType + `",
`

		}

		formTypes = formTypes + `
			}

			return fields
		}`
		l5, err := f.WriteString(formTypes)
		if err != nil {
			fmt.Println(err, l5)
			f.Close()

		}



			formula := `
			func (v *` + modelAlias + strconv.FormatUint(vb.ID, 10) + `) GetFormula() string{

			return `
		if len(schema.Formula) >= 1 {
			stringFormula, _ := json.Marshal(schema.Formula)


			var re = regexp.MustCompile(`"`)
			jsonString := re.ReplaceAllString(string(stringFormula), `\"`)

			formula = formula + `"`+jsonString+`"`
		} else {
			formula = formula + `""`
		}
		formula = formula +  `
		}`


			l6, err := f.WriteString(formula)
			if err != nil {
				fmt.Println(err, l6)
				f.Close()

			}








			err = f.Close()
		if err != nil {
			fmt.Println(err)

		}

	}

}
func WriteModelCaller(forms []models.VBSchema, moduleName string) {
	//return new(models.Naiz)

	f, err := os.Create("models/form/caller/modelCaller.go")
	if err != nil {
		fmt.Println(err)

	}
	l1, err := f.WriteString("package caller\n")
	if err != nil {
		fmt.Println(err, l1)
		f.Close()

	}
	l2, err := f.WriteString("import \""+moduleName+"/models/form\"\n")
	if err != nil {
		fmt.Println(err, l2)
		f.Close()

	}
	l3, err := f.WriteString("func GetMODEL(schema_id string) (string, interface{}) {\n\nswitch schema_id {\n" + `
 case "crud_form":
return "id", new(form.CrudFrom)

 case "analytic_form":
return "id", new(form.AnalyticForm)

 case "notification_target_form":
return "id", new(form.NotificationTarget)

 case "menu_form":
return "id", new(form.MenuForm)

 case "user_form":
return "id", new(form.UserForm)

 case "user_profile":
return "id", new(form.UserProfile)

 case "user_password":
return "id", new(form.UserPassword)

`)
	if err != nil {
		fmt.Println(err, l3)
		f.Close()

	}
	for _, vb := range forms {
		var schema SCHEMA

		json.Unmarshal([]byte(vb.Schema), &schema)

		modelAlias := getModelAlias(schema.Model)


		l5, err := f.WriteString("\n case \"" + strconv.FormatUint(vb.ID, 10) + "\": \nreturn \""+schema.Identity+"\",  new(form." + modelAlias + strconv.FormatUint(vb.ID, 10) + ")\n")
		if err != nil {
			fmt.Println(err, l5)
			f.Close()

		}
	}

	//fmt.Println(schema.Model, "schema.Model")

	l4, err := f.WriteString("\n} \nreturn \"id\", new(interface{})\n\n}")
	if err != nil {
		fmt.Println(err, l4)
		f.Close()

	}
	//fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)

	}
}
func WriteValidationCaller(forms []models.VBSchema, moduleName string) {

	f, err := os.Create("models/form/validationCaller/rulesCaller.go")
	if err != nil {
		fmt.Println(err)

	}
	l1, err := f.WriteString("package validationCaller\n")
	if err != nil {
		fmt.Println(err, l1)
		f.Close()

	}
	l2, err := f.WriteString(`import (
	"github.com/thedevsaddam/govalidator"
	"`+moduleName+`/models/form/validations"
)

`)
	if err != nil {
		fmt.Println(err, l2)
		f.Close()

	}
	l3, err := f.WriteString(`func GetRules(schema_id string) map[string][]string {

	switch schema_id {


	case "crud_form":
		return validations.GetCrudFromRules()

	case "analytic_form":
		return validations.GetAnalyticFormRules()

	case "notification_target_form":
		return validations.GetNotificationTargetRules()

	case "menu_form":
		return validations.GetMenuFormRules()

	case "user_form":
		return validations.GetUserFormRules()

	case "user_profile":
		return validations.GetUserProfileRules()

	case "user_password":
		return validations.GetUserPasswordRules()

	
`)
	if err != nil {
		fmt.Println(err, l3)
		f.Close()

	}
	for _, vb := range forms {
		var schema SCHEMA

		json.Unmarshal([]byte(vb.Schema), &schema)

		WriteModelValidation(schema, vb.ID)
		modelAlias := getModelAlias(schema.Model)

		l5, err := f.WriteString("\n case \"" + strconv.FormatUint(vb.ID, 10) + "\": \nreturn validations.Get" + modelAlias + strconv.FormatUint(vb.ID, 10) + "Rules()\n")
		if err != nil {
			fmt.Println(err, l5)
			f.Close()

		}
	}

	//fmt.Println(schema.Model, "schema.Model")

	l4, err := f.WriteString("\n} \nreturn govalidator.MapData{}\n\n}")
	if err != nil {
		fmt.Println(err, l4)
		f.Close()

	}
	//fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)

	}

	WriteValidationMessageCaller(forms, moduleName)
}
func WriteValidationMessageCaller(forms []models.VBSchema, moduleName string) {

	f, err := os.Create("models/form/validationCaller/messagesCaller.go")
	if err != nil {
		fmt.Println(err)

	}
	l1, err := f.WriteString("package validationCaller\n")
	if err != nil {
		fmt.Println(err, l1)
		f.Close()

	}
	l2, err := f.WriteString(`import (
	"github.com/thedevsaddam/govalidator"
    "`+moduleName+`/models/form/validations"
)

`)
	if err != nil {
		fmt.Println(err, l2)
		f.Close()

	}
	l3, err := f.WriteString(`func GetMessages(schema_id string) map[string][]string {

	switch schema_id {

	case "crud_form":
		return validations.GetCrudFromMessages()

	case "analytic_form":
		return validations.GetAnalyticFormMessages()

	case "notification_target_form":
		return validations.GetNotificationTargetMessages()

	case "menu_form":
		return validations.GetMenuFormMessages()

	case "user_form":
		return validations.GetUserFormMessages()

	case "user_profile":
		return validations.GetUserProfileMessages()

	case "user_password":
		return validations.GetUserPasswordMessages()


`)
	if err != nil {
		fmt.Println(err, l3)
		f.Close()

	}
	for _, vb := range forms {
		var schema SCHEMA

		json.Unmarshal([]byte(vb.Schema), &schema)

		modelAlias := getModelAlias(schema.Model)

		l5, err := f.WriteString("\n case \"" + strconv.FormatUint(vb.ID, 10) + "\": \nreturn validations.Get" + modelAlias + strconv.FormatUint(vb.ID, 10) + "Messages()\n")
		if err != nil {
			fmt.Println(err, l5)
			f.Close()

		}
	}

	//fmt.Println(schema.Model, "schema.Model")

	l4, err := f.WriteString("\n} \nreturn govalidator.MapData{}\n\n}")
	if err != nil {
		fmt.Println(err, l4)
		f.Close()

	}
	//fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)

	}

}
func WriteModelValidation(vb SCHEMA, ID uint64) {

	id_ := strconv.FormatUint(ID, 10)

	modelAlias := getModelAlias(vb.Model)
	f, err := os.Create("models/form/validations/" + vb.Model + id_ + ".go")
	if err != nil {
		fmt.Println(err, f)
	}

	l1, err := f.WriteString(`package validations

import "github.com/thedevsaddam/govalidator"


func Get` + modelAlias + id_ + `Rules() map[string][]string{
	return  govalidator.MapData{
		`)
	if err != nil {
		fmt.Println(err, l1)
		f.Close()

	}

	for _, field := range vb.Schema {

		if len(field.Rules) >= 1 && vb.Identity != field.Model && field.Model != "created_at" && field.Model != "updated_at" && field.Model != "deleted_at" {
			rules := ""
			for _, rule := range field.Rules {

				if rule.Type != "unique"{
					if rule.Type == "number"{
						rules = rules + "\"" + "numeric" + "\","
					} else {
						rules = rules + "\"" + rule.Type + "\","
					}

				}


			}

			rules = "\n\"" + field.Model + "\": []string{" + rules + "},"

			l5, err := f.WriteString(rules)
			if err != nil {
				fmt.Println(err, l5)
				f.Close()

			}
		}
	}

	l3, err := f.WriteString(`
	}
}
func Get` + modelAlias + id_ + `Messages() map[string][]string{
	return govalidator.MapData{
`)
	if err != nil {
		fmt.Println(err, l3)
		f.Close()

	}

	for _, field := range vb.Schema {

		if len(field.Rules) >= 1 && vb.Identity != field.Model && field.Model != "created_at" && field.Model != "updated_at" && field.Model != "deleted_at" {
			rules := ""
			for _, rule := range field.Rules {

				rules = rules + "\"" + rule.Type + ":" + rule.Msg + "\","

			}

			rules = "\n\"" + field.Model + "\": []string{" + rules + "},"

			l5, err := f.WriteString(rules)
			if err != nil {
				fmt.Println(err, l5)
				f.Close()

			}
		}
	}

	l2, err := f.WriteString(`
	}
}
`)
	if err != nil {
		fmt.Println(err, l2)
		f.Close()

	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)

	}

}

func getModelAlias(modelName string) string {
	return strmangle.TitleCase(strmangle.Singular(modelName))
}
