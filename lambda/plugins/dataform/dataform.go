package dataform

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	agentUtils "github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/DBSchema"
	puzzleModel "github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	"github.com/thedevsaddam/govalidator"
	"io/ioutil"
	"regexp"
	"github.com/PaesslerAG/gval"
	"github.com/araddon/dateparse"
	//"github.com/khankhulgun/khankhulgun/tools"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Exec(c echo.Context, schemaId string, action string, id string, GetMODEL func(schema_id string) (string, interface{}), GetMessages func(schema_id string) map[string][]string, GetRules func(schema_id string) map[string][]string) error {

	Identity, Model := GetMODEL(schemaId)
	switch action {
	case "store":
		return Store(c, Model, schemaId, id, action, Identity, GetMessages, GetRules)
	case "update":
		return Store(c, Model, schemaId, id, action, Identity, GetMessages, GetRules)
	case "edit":
		return Edit(c, Model, schemaId, id, Identity)
	case "options":
		return Options(c)
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status": "false",
	})

}

func Edit(c echo.Context, Model interface{}, schemaId string, id string, Identity string) error {

	GetSubForms := reflect.ValueOf(Model).MethodByName("GetSubForms")

	if GetSubForms.IsValid() {
		GetSubFormsRes := GetSubForms.Call([]reflect.Value{})
		SubForms := GetSubFormsRes[0].Interface().([]map[string]interface{})

		if len(SubForms) >= 1 {

			DB.DB.Where(Identity+" = ?", id).Find(Model)

			data := make(map[string]interface{})
			dataPre, _ := json.Marshal(Model)
			json.Unmarshal(dataPre, &data)

			for _, Sub := range SubForms {

				connectionField := Sub["connection_field"].(string)
				subTable := Sub["table"].(string)

				DB.DB.Where(connectionField+" = ?", id).Find(Sub["subForm"])

				dataSub := []map[string]interface{}{}
				subData, _ := json.Marshal(Sub["subForm"])
				json.Unmarshal(subData, &dataSub)

				dataWitSub := []map[string]interface{}{}
				for _, sData := range dataSub {

					subIdentity := Sub["subIdentity"].(string)

					parentId := fmt.Sprintf("%g", sData[subIdentity])
					subFormModel := Sub["subFormModel"]

					GetSubForms2 := reflect.ValueOf(subFormModel).MethodByName("GetSubForms")

					if GetSubForms2.IsValid() {
						GetSubFormsRes2 := GetSubForms2.Call([]reflect.Value{})
						SubForms2 := GetSubFormsRes2[0].Interface().([]map[string]interface{})

						for _, Sub2 := range SubForms2 {

							connectionField2 := Sub2["connection_field"].(string)
							subTable2 := Sub2["table"].(string)
							DB.DB.Where(connectionField2+" = ?", parentId).Find(Sub2["subForm"])
							sData[subTable2] = Sub2["subForm"]

						}
					}

					dataWitSub = append(dataWitSub, sData)
				}

				data[subTable] = dataWitSub

			}

			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": "true",
				"data":   data,
			})

		} else {
			data := DB.DB.Where(Identity+" = ?", id).Find(Model)

			if data.Error == nil {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"status": "true",
					"data":   data.Value,
				})
			}
		}
	} else {
		data := DB.DB.Where(Identity+" = ?", id).Find(Model)

		if data.Error == nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": "true",
				"data":   data.Value,
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status": "false",
	})
}

func saveNestedSubItem(ParentModel interface{}, data map[string]interface{}) {

	GetSubForms := reflect.ValueOf(ParentModel).MethodByName("GetSubForms")

	if GetSubForms.IsValid() {
		GetSubFormsRes := GetSubForms.Call([]reflect.Value{})

		SubForms := GetSubFormsRes[0].Interface().([]map[string]interface{})

		if len(SubForms) >= 1 {

			for _, Sub := range SubForms {
				table := Sub["table"].(string)
				parentIdentity := Sub["parentIdentity"].(string)
				subIdentity := Sub["subIdentity"].(string)
				connectionField := Sub["connection_field"].(string)
				subForm := Sub["subFormModel"]
				subData := data[table]

				if subData != nil {

					parentData := make(map[string]interface{})
					parentDataPre, _ := json.Marshal(ParentModel)
					json.Unmarshal(parentDataPre, &parentData)
					parentId := parentData[parentIdentity]

					DB.DB.Where(connectionField+" = ?", parentId).Delete(subForm)
					currentData := reflect.ValueOf(subData).Interface().([]interface{})

					//fmt.Println(table)
					//fmt.Println("======= sub table")

					if len(currentData) >= 1 {

						for _, sData := range currentData {

							subD := reflect.ValueOf(sData).Interface().(map[string]interface{})

							subIdentityValue := subD[subIdentity]

							subD[connectionField] = parentId

							if subIdentityValue == nil || config.Config.Database.Connection == "mssql"{
								subD[subIdentity] = 0
							}




							saveData, _ := json.Marshal(subD)
							fmt.Println(string(saveData))
							json.Unmarshal(saveData, &subForm)
							fmt.Println(subForm)

							DB.DB.NewRecord(subForm)
							err := DB.DB.Create(subForm).Error

							//err := DB.DB.Save(subForm).Error

							if err == nil {
								callTrigger("afterUpdate", subForm, subD, "")

								saveNestedSubItem(subForm, subD)
							}

							/*creareNewRow := true


							switch vtype := subIdentityValue.(type) {
							case int:
								if(subIdentityValue.(int) >= 1){
									creareNewRow = false

								}
							case float64:
								if(subIdentityValue.(float64) >= 1){
									creareNewRow = false

								}
							case float32:
								if(subIdentityValue.(float32) >= 1){
									creareNewRow = false

								}
							case int64:

								if(subIdentityValue.(int64) >= 1){
									creareNewRow = false

								}
							default:

								fmt.Println(vtype)
							}


							if (!creareNewRow){


								err := DB.DB.Save(subForm).Error

								if err == nil {
									callTrigger("afterUpdate", subForm, subD, "")

									saveNestedSubItem(subForm, subD)
								}

							} else {

								DB.DB.NewRecord(subForm)
								err := DB.DB.Create(subForm).Error

								if err == nil {

									callTrigger("afterInsert", subForm, subD, "")
									saveNestedSubItem(subForm, subD)

								}

							}*/

						}
					}

				}

			}
		}
	}

}

func DataClear(c echo.Context, Model interface{}, action string, id string, rules map[string][]string) (interface{}, *map[string]interface{}, map[string][]string) {
	filterRaw, _ := ioutil.ReadAll(c.Request().Body)

	dataJson := new(map[string]interface{})
	json.Unmarshal([]byte(filterRaw), dataJson)

	getFromTypes := reflect.ValueOf(Model).MethodByName("GetFromTypes")
	GetFormula := reflect.ValueOf(Model).MethodByName("GetFormula")
	GetTableName := reflect.ValueOf(Model).MethodByName("TableName")

	/*
	ONLY USE MANAIKHOROO
	*/
	if GetTableName.IsValid() {
		getTableRes := GetTableName.Call([]reflect.Value{})
		tableName := getTableRes[0].Interface().(string)

		if(tableName == "zurchil"){

			if (*dataJson)["id"] == nil {
				if (*dataJson)["user_id"] == nil || (*dataJson)["user_id"] == 0 {
					User := agentUtils.AuthUserObject(c)
					(*dataJson)["user_id"] = User["id"]
				}
			}
		}

	}

	if getFromTypes.IsValid() {
		getFromTypesRes := getFromTypes.Call([]reflect.Value{})
		formTypes := getFromTypesRes[0].Interface().(map[string]string)

		if len(formTypes) >= 1 {

			for key, Value := range formTypes {
				if Value == "Date" {
					if (*dataJson)[key] != nil {
						parsedDate, err := time.Parse("2006-01-02", (*dataJson)[key].(string))

						if err != nil {
							parsedDate2, err2 := time.Parse(time.RFC3339, (*dataJson)[key].(string))
							if err2 != nil {
								panic(err2)
							}

							//parsedDate2Converted, err3 := time.Parse("2006-01-02", parsedDate2.Format("2006-01-02"))
							//if err3 != nil {
							//	panic(err3)
							//}
							//parsedDate = parsedDate2Converted

							//fmt.Println(key, parsedDate2.Format("2006-01-02"))
							//(*dataJson)[key] = parsedDate2.Format("2006-01-02")
							//delete(rules, key)

							parsedDate = parsedDate2
						}

						(*dataJson)[key] = parsedDate
						delete(rules, key)

					} else {
						(*dataJson)[key] = nil
					}

				}
				if Value == "DateTime" {
					if (*dataJson)[key] != nil {
						parsedDate, err := dateparse.ParseLocal((*dataJson)[key].(string))

						if err != nil {
							parsedDate2, err2 := time.Parse(time.RFC3339, (*dataJson)[key].(string))
							if err2 != nil {
								panic(err2)
							}

							//parsedDate2Converted, err3 := time.Parse("2006-01-02", parsedDate2.Format("2006-01-02"))
							//if err3 != nil {
							//	panic(err3)
							//}
							//parsedDate = parsedDate2Converted

							//fmt.Println(key, parsedDate2.Format("2006-01-02"))
							//(*dataJson)[key] = parsedDate2.Format("2006-01-02")
							//delete(rules, key)

							parsedDate = parsedDate2
						}

						(*dataJson)[key] = parsedDate
						delete(rules, key)

					} else {
						(*dataJson)[key] = nil
					}

				}
				if Value == "Password" {
					if action == "store" {
						password, _ := agentUtils.Hash((*dataJson)[key].(string))
						(*dataJson)[key] = password
					} else {
						if((*dataJson)[key] == nil){
							delete(rules, key)
						} else {
							if len((*dataJson)[key].(string)) >= 1 {
								password, _ := agentUtils.Hash((*dataJson)[key].(string))
								(*dataJson)[key] = password
							} else {
								delete(rules, key)
							}
						}

					}
				}
			}
		}
	}

	if id != "" {
		*dataJson = callTrigger("beforeUpdate", Model, *dataJson, id)
	}
	data, _ := json.Marshal(dataJson)
	json.Unmarshal(data, Model)

	if GetFormula.IsValid() {
		GetFormulaRes := GetFormula.Call([]reflect.Value{})
		formulaString := GetFormulaRes[0].Interface().(string)

		if formulaString != "" {
			formulas := []DBSchema.Formula{}
			json.Unmarshal([]byte(formulaString), &formulas)

			if len(formulas) >= 1 {
				for _, formula := range formulas {

					for _, target := range formula.Targets {
						if target.Prop == "hidden" {



							var re1 = regexp.MustCompile(`'{`)
							template := re1.ReplaceAllString(formula.Template, ``)
							var re2 = regexp.MustCompile(`}'`)
							template = re2.ReplaceAllString(template, ``)
							var re3 = regexp.MustCompile(`'`)
							template = re3.ReplaceAllString(template, `"`)

							value, _ := gval.Evaluate(template, *dataJson)

							if value == true {

								delete(rules, target.Field)
							}
						}
					}

				}
			}
		}

	}

	return Model, dataJson, rules
}

func Store(c echo.Context, FromModel interface{}, schemaId string, id string, action string, Identity string, GetMessages func(schema_id string) map[string][]string, GetRules func(schema_id string) map[string][]string) error {

	/*FORM VALIDATION*/
	messages := GetMessages(schemaId)
	Model, dataJson, rules := DataClear(c, FromModel, action, id, GetRules(schemaId))

	if len(rules) >= 1 {
		opts := govalidator.Options{
			Data:     Model,    // request object
			Rules:    rules,    // rules map
			Messages: messages, // custom message map (Optional)
			//RequiredDefault: true,     // all the field to be pass the rules
		}
		v := govalidator.New(opts)
		e := v.ValidateStruct()

		if len(e) >= 1 {
			err := map[string]interface{}{"error": e}
			err["status"] = false
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	if id != "" {
		err := DB.DB.Save(Model).Error
		if err != nil {


			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status": false,
				"error": err.Error(),
			})
		} else {

			saveNestedSubItem(Model, *dataJson)

			data := callTrigger("afterUpdate", Model, *dataJson, id)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": true,
				"data":   data,
			})
		}
	} else {
		DB.DB.NewRecord(&Model)
		err := DB.DB.Create(Model).Error

		if err != nil {



			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status": false,
				"error": err.Error(),
			})
		} else {

			saveNestedSubItem(Model, *dataJson)

			var formInterface map[string]interface{}
			inrec, _ := json.Marshal(Model)
			json.Unmarshal(inrec, &formInterface)
			data := callTrigger("afterInsert", Model, *dataJson, fmt.Sprintf("%v", formInterface[Identity]))
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": true,
				"data":   data,
			})
		}
	}

}

func callTrigger(action string, Model interface{}, data map[string]interface{}, id string) map[string]interface{} {

	GetTriggers := reflect.ValueOf(Model).MethodByName("GetTriggers")

	if GetTriggers.IsValid() {

		GetTriggersRes := GetTriggers.Call([]reflect.Value{})

		triggers := GetTriggersRes[0].Interface().(map[string]map[string]interface{})
		namespace := GetTriggersRes[1].Interface().(string)

		if len(triggers) <= 0 {
			return data
		}

		if namespace == "" {
			return data
		}

		switch action {
		case "beforeInsert":
			Method := triggers["insert"]["before"].(string)
			Struct := triggers["insert"]["beforeStruct"]
			return execTrigger(Method, Struct, Model, data, id)
		case "afterInsert":
			Method := triggers["insert"]["after"].(string)
			Struct := triggers["insert"]["afterStruct"]
			return execTrigger(Method, Struct, Model, data, id)
		case "beforeUpdate":
			Method := triggers["update"]["before"].(string)
			Struct := triggers["update"]["beforeStruct"]
			return execTrigger(Method, Struct, Model, data, id)
		case "afterUpdate":
			Method := triggers["update"]["after"].(string)
			Struct := triggers["update"]["afterStruct"]
			return execTrigger(Method, Struct, Model, data, id)

		}
		return data
	} else {
		return data
	}

}

func execTrigger(triggerMethod string, triggerStruct interface{}, Model interface{}, data map[string]interface{}, id string) map[string]interface{} {

	if triggerMethod != "" {
		triggerMethod_ := reflect.ValueOf(triggerStruct).MethodByName(triggerMethod)

		if triggerMethod_.IsValid() {

			input := make([]reflect.Value, 3)
			input[0] = reflect.ValueOf(data)
			input[1] = reflect.ValueOf(id)
			input[2] = reflect.ValueOf(Model)
			triggerMethodRes := triggerMethod_.Call(input)

			return triggerMethodRes[0].Interface().(map[string]interface{})
		}
		return data
	} else {
		return data
	}

}

type UniquePost struct {
	Table          string `json:"table"`
	IdentityColumn string `json:"identityColumn"`
	Identity       uint64 `json:"identity"`
	Field          string `json:"field"`
	Val            string `json:"val"`
}

func CheckUnique(c echo.Context) error {

	post := new(UniquePost)
	if err := c.Bind(post); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false from json",
		})
	}

	DB_ := DBSchema.DBConnection()
	var count int

	if post.IdentityColumn != "" {
		err := DB_.QueryRow("SELECT COUNT(*) FROM " + post.Table + " WHERE " + post.IdentityColumn + " != '" + strconv.FormatUint(post.Identity, 10) + "' AND " + post.Field + " = '" + post.Val + "'").Scan(&count)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status": false,
			})
		} else {

			if count > 0 {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"status": false,
					"msg":    "'" + post.Val + "' утга бүртгэлтэй байна",
				})
			} else {

				return c.JSON(http.StatusOK, map[string]interface{}{
					"status": true,
				})
			}
		}
	} else {
		err := DB_.QueryRow("SELECT COUNT(*) FROM " + post.Table + " WHERE " + post.Field + " = '" + post.Val + "'").Scan(&count)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status": false,
			})
		} else {

			if count > 0 {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"status": false,
					"msg":    "'" + post.Val + "' утга бүртгэлтэй байна",
				})
			} else {

				return c.JSON(http.StatusOK, map[string]interface{}{
					"status": true,
				})
			}
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status": false,
	})
}

type Relations struct {
	Relations map[string]Ralation_ `json:"relations"`
}

type Ralation_ struct {
	Fields             []string            `json:"Fields"`
	FilterWithUser     []map[string]string `json:"filterWithUser"`
	Filter             string              `json:"filter"`
	Key                string              `json:"key"`
	Multiple           bool                `json:"multiple"`
	ParentFieldOfForm  string              `json:"parentFieldOfForm"`
	ParentFieldOfTable string              `json:"parentFieldOfTable"`
	SortField          string              `json:"sortField"`
	SortOrder          string              `json:"sortOrder"`
	Table              string              `json:"table"`
}

type RalationOption struct {
	Fields         []string            `json:"Fields"`
	FilterWithUser []map[string]string `json:"filterWithUser"`
	Filter         string              `json:"filter"`
	Key            string              `json:"key"`
	SortField      string              `json:"sortField"`
	SortOrder      string              `json:"sortOrder"`
	Table          string              `json:"table"`
}

func Options(c echo.Context) error {
	r := new(RalationOption)
	if err := c.Bind(r); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": err.Error(),
		})
	}

	Relation := Ralation_{}

	Relation.Fields = r.Fields
	Relation.Filter = r.Filter
	Relation.Key = r.Key
	Relation.SortField = r.SortField
	Relation.SortOrder = r.SortOrder
	Relation.Table = r.Table
	Relation.FilterWithUser = r.FilterWithUser

	var DB_ *sql.DB
	DB_ = DB.DB.DB()
	data := OptionsData(DB_, Relation, c)
	return c.JSON(http.StatusOK, data)
}

func OptionsData(DB *sql.DB, relation Ralation_, c echo.Context) []map[string]interface{} {

	table := relation.Table
	labels := strings.Join(relation.Fields[:], ",', ',")
	key := relation.Key
	sortField := relation.SortField
	sortOrder := relation.SortOrder
	parentFieldOfTable := relation.ParentFieldOfTable
	filter := relation.Filter
	FilterWithUser := relation.FilterWithUser

	//fmt.Println(FilterWithUser)
	data := []map[string]interface{}{}

	if table == "" || len(labels) < 1 || key == "" {
		return data
	}
	var parent_column string
	if parentFieldOfTable != "" {
		parent_column = ", " + parentFieldOfTable + " as parent_value"
	}
	var order_value string
	if sortField != "" && sortOrder != "" {
		order_value = "order by " + sortField + " " + sortOrder
	}
	var where_value string
	if filter != "" {
		where_value = "WHERE " + filter
	}

	if len(FilterWithUser) >= 1 {

		User := agentUtils.AuthUserObject(c)
		for _, userCon := range FilterWithUser {

			tableField := userCon["tableField"]
			userField := User[userCon["userField"]]

			//if userField
			if(userField != nil){
				userFieldValue := strconv.FormatInt(reflect.ValueOf(userField).Int(), 10)

				userDataFilter := tableField + " = '" + userFieldValue + "'"

				if userFieldValue != "" && userFieldValue != "0" {
					if where_value == "" {
						where_value = "WHERE " + userDataFilter
					} else {
						where_value = where_value + " AND " + userDataFilter
					}
				}
			}
		}
	}

		//fmt.Println("SELECT " + key + " as value, concat(" + labels + ") as label " + parent_column + "  FROM " + table + " " + where_value + " " + order_value)


	concatTxt := "CONCAT"
	if config.Config.Database.Connection == "mssql"{
		if len(relation.Fields) <= 1 {
			concatTxt = ""
		}

	}

	//rows, _ := DB.Query("SELECT " + key + " as value, "+concatTxt+"(" + labels + ") as label " + parent_column + "  FROM " + table + " " + where_value + " " + order_value)
	//
	//fmt.Println("SELECT " + key + " as value, "+concatTxt+"(" + labels + ") as label " + parent_column + "  FROM " + table + " " + where_value + " " + order_value)
	//

	return  GetTableData("SELECT " + key + " as value, "+concatTxt+"(" + labels + ") as label " + parent_column + "  FROM " + table + " " + where_value + " " + order_value)

	///*start*/
	//
	//columns, _ := rows.Columns()
	//count := len(columns)
	//values := make([]interface{}, count)
	//valuePtrs := make([]interface{}, count)
	//
	///*end*/
	//
	//for rows.Next() {
	//
	//	/*start */
	//
	//	for i := range columns {
	//		valuePtrs[i] = &values[i]
	//	}
	//
	//	rows.Scan(valuePtrs...)
	//
	//	var myMap = make(map[string]interface{})
	//	for i, col := range columns {
	//		val := values[i]
	//
	//		b, ok := val.([]byte)
	//
	//		if (ok) {
	//
	//			v, error := strconv.ParseInt(string(b), 10, 64)
	//			if error != nil {
	//				stringValue := string(b)
	//
	//				myMap[col] = stringValue
	//			} else {
	//				myMap[col] = v
	//			}
	//
	//		}
	//
	//	}
	//	/*end*/
	//
	//	data = append(data, myMap)
	//
	//}
	return data
}

func GetTableData(query string)[]map[string]interface{}  {
	data := []map[string]interface{}{}

	rows, err := DB.DB.DB().Query(query)

	if(err != nil){
		fmt.Println(err.Error())
	
		return data
	}
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

			if config.Config.Database.Connection == "mssql"{
				myMap[col] = val

			} else {
				b, ok := val.([]byte)



				if (ok) {

					v, error := strconv.ParseInt(string(b), 10, 64)
					if error != nil {
						stringValue := string(b)

						myMap[col] = stringValue
					} else {
						myMap[col] = v
					}

				}
			}



		}
		/*end*/

		data = append(data, myMap)

	}

	return data

}

func SetCondition(condition string, c echo.Context, VBSchema puzzleModel.VBSchema) error {

	con, _ := url.ParseQuery(condition)
	var schema DBSchema.SCHEMA

	json.Unmarshal([]byte(VBSchema.Schema), &schema)

	for uC, _ := range con {

		uString := reflect.ValueOf(uC).Interface().(string)

		var conditionData []map[string]string
		json.Unmarshal([]byte(uString), &conditionData)

		User := agentUtils.AuthUserObject(c)
		for _, userCondition := range conditionData {

			for i := range schema.Schema {

				if schema.Schema[i].Model == userCondition["form_field"] {
					schema.Schema[i].Disabled = true
					schema.Schema[i].Default = User[userCondition["user_field"]]
				}

			}
		}

	}
	schemaString, _ := json.Marshal(schema)
	VBSchema.Schema = string(schemaString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "true",
		"data":   VBSchema,
	})
}