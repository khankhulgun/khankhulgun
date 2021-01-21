package handlers

import (
	"fmt"

	"net/http"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
	"github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS/models"
	"github.com/khankhulgun/khankhulgun/DB"
	"encoding/json"
	"reflect"
	"regexp"
	"net/url"
	"io/ioutil"
	"bytes"
	"strconv"
	"crypto/tls"
	"github.com/grokify/html-strip-tags-go"
)
func SAVEGIS (rawData []byte, schemaId int64, action string, rowId string, GetMODEL func(schema_id string) (string, interface{})){


	dataJson := new(map[string]interface{})

	json.Unmarshal(rawData, dataJson)
	Identity, Model := GetMODEL(strconv.Itoa(int(schemaId)))


	TableName := reflect.ValueOf(Model).MethodByName("TableName")

	if TableName.IsValid() {
		TableNameData := TableName.Call([]reflect.Value{})
		table := TableNameData[0].Interface().(string)

		if(action == "store" || action == "update"){
			(*dataJson)[Identity] = rowId

			ArcGISAfterSave(*dataJson, schemaId, Model, Identity, rowId, table)

		}


	}



}
func DELTEGIS (rawData []byte, schemaId int64, action string, rowId string, GetGridMODEL func(schema_id string) (interface{}, interface{}, string, string, interface{}, string)){


	dataJson := new(map[string]interface{})

	json.Unmarshal(rawData, dataJson)

	_, _, _, _, MainTableModel, Identity := GetGridMODEL(strconv.Itoa(int(schemaId)))



	if(action == "delete"){
		ArcGISAfterDelete(MainTableModel, nil, rowId, strconv.Itoa(int(schemaId)), Identity)
	}



}
func FormFields(c echo.Context) error {

	data := new(Data)
	if err := c.Bind(data); err != nil {


		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": false,
		})
	}


	form_id := data.Model.LocalForm

	var schema Schema

	DB.DB.Table("vb_schemas").Where("id = ?", form_id).Find(&schema)

	var schemaData SchemaData

	json.Unmarshal([]byte(schema.Schema), &schemaData)

	fields := []interface{}{}

	for _, field := range schemaData.Schema{


		fields = append(fields, field["model"])
	}

	Schemas := []interface{}{}

	connection := Field{
		Field: "connection",
		//Value: 1,
		Props: map[string]interface{}{
			"options": fields,
		},
	}
	Schemas = append(Schemas, connection)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"schema":  Schemas,

	})

}
func ArcGISAfterSave(data map[string]interface{}, schemaId int64,  Model interface{}, Identity string, rowId string, table string) map[string]interface{} {


	if schemaId >= 1 {

		formID := schemaId


		var connectionPre models.ArcgisConnection
		DB.DB.Where("local_form = ?", formID).Find(&connectionPre)


		if connectionPre.Connection != ""{

			var connection Connection

			json.Unmarshal([]byte(connectionPre.Connection), &connection)

			objectIDSTRING := fmt.Sprintf("%v", data[connection.ObjectIdField])

			objectID, _ := strconv.ParseInt(objectIDSTRING, 10, 64)

			attributes := ""

			for _, attribute := range connection.Connections {

				//fmt.Println(attribute["attribute"])


				if data[attribute["field"]] != nil {

					fieldValue := ""
					if reflect.TypeOf(data[attribute["field"]]).String() == "time.Time"{
						fieldValue = "\""+reflect.ValueOf(data[attribute["field"]]).Interface().(time.Time).Format("2006-01-02")+"\""

					} else if reflect.TypeOf(data[attribute["field"]]).String() == "float64" {

						fieldValue = fmt.Sprintf("%.0f",data[attribute["field"]])

					} else if reflect.TypeOf(data[attribute["field"]]).String() == "string"{

						var re = regexp.MustCompile(`"`)
						fieldValue = re.ReplaceAllString(reflect.ValueOf(data[attribute["field"]]).Interface().(string), `\"`)

						fieldValue = strip.StripTags(fieldValue)
						fieldValue = "\""+fieldValue+"\""


					} else {
						fieldValue = "\""+fmt.Sprintf("%v",data[attribute["field"]])+"\""
					}
					var re = regexp.MustCompile(`\r?\n`)
					fieldValue = re.ReplaceAllString(fieldValue, `\n`)

					if fieldValue == "" {
						fieldValue = "\""+"\""
					}

					if attributes == "" {

						attributes = "\""+attribute["attribute"]+"\""+":"+fieldValue
					} else {

						attributes = attributes+","+"\""+attribute["attribute"]+"\""+":"+fieldValue
					}
				}


			}

			if objectID >= 1{
				if attributes == "" {
					attributes = "\"OBJECTID\":"+fmt.Sprintf("%v",objectID)
				} else {
					attributes = attributes+",\"OBJECTID\":"+fmt.Sprintf("%v",objectID)
				}
			}



			geoData := data[connection.GeoJsonField].(string)


			if connection.LayerType == "Point"{

				var coordinate Coordinate

				json.Unmarshal([]byte(geoData), &coordinate)


				if coordinate.Lat == ""{

					var coordinate CoordinateFloat
					json.Unmarshal([]byte(geoData), &coordinate)


					features := ""
					if(data["coord_z"] != nil){
						features = fmt.Sprintf("[{\"geometry\":{\"x\":%v,\"y\":%v,\"z\":%v,\"spatialReference\":{\"wkid\":4326}},\"attributes\":{%v}}]", fmt.Sprintf("%v",coordinate.Lng), fmt.Sprintf("%v",coordinate.Lat), fmt.Sprintf("%v",data["coord_z"]), attributes)
					} else {
						features = fmt.Sprintf("[{\"geometry\":{\"x\":%v,\"y\":%v,\"z\":%v,\"spatialReference\":{\"wkid\":4326}},\"attributes\":{%v}}]", fmt.Sprintf("%v",coordinate.Lng), fmt.Sprintf("%v",coordinate.Lat), attributes)
					}


					layers := []string{}
					json.Unmarshal([]byte(connectionPre.Layer), &layers)


					if len(layers) >= 1 {

						layer_url := layers[len(layers)-1]
						objectIDNew := SaveArcGIS(layer_url, features, objectID)

						if objectIDNew >= 1{
							//data[connection.ObjectIdField] = objectIDNew
							//data_, _ := json.Marshal(data)
							//json.Unmarshal(data_, Model)
							dataGIS := map[string]interface{}{}
							dataGIS[connection.ObjectIdField] = objectIDNew
							DB.DB.Table(table).Where(Identity+" = ?", rowId).Update(dataGIS)

						}
					}
				}  else {

					features := ""
					if(data["coord_z"] != nil){
						features = fmt.Sprintf("[{\"geometry\":{\"x\":%v,\"y\":%v,\"z\":%v,\"spatialReference\":{\"wkid\":4326}},\"attributes\":{%v}}]", fmt.Sprintf("%v",coordinate.Lng), fmt.Sprintf("%v",coordinate.Lat), fmt.Sprintf("%v",data["coord_z"]), attributes)
					} else {
						features = fmt.Sprintf("[{\"geometry\":{\"x\":%v,\"y\":%v,\"z\":%v,\"spatialReference\":{\"wkid\":4326}},\"attributes\":{%v}}]", fmt.Sprintf("%v",coordinate.Lng), fmt.Sprintf("%v",coordinate.Lat), attributes)
					}

					layers := []string{}
					json.Unmarshal([]byte(connectionPre.Layer), &layers)


					if len(layers) >= 1 {

						layer_url := layers[len(layers)-1]
						objectIDNew := SaveArcGIS(layer_url, features, objectID)

						if objectIDNew >= 1{
							dataGIS := map[string]interface{}{}
							dataGIS[connection.ObjectIdField] = objectIDNew
							DB.DB.Table(table).Where(Identity+" = ?", rowId).Update(dataGIS)

						}
					}
				}


			} else if connection.LayerType == "Polygon" || connection.LayerType == "Line"{
				var geoJson GeoJSon

				json.Unmarshal([]byte(geoData), &geoJson)

				if len(geoJson.Features) >= 1{
					Coordinates, _ := json.Marshal(geoJson.Features[0].Geometry.Coordinates)



					features := fmt.Sprintf("[{\"geometry\":{\"rings\":%v,\"spatialReference\":{\"wkid\":4326}},\"attributes\":{%v}}]", string(Coordinates), attributes)


					layers := []string{}
					json.Unmarshal([]byte(connectionPre.Layer), &layers)


					if len(layers) >= 1 {

						layer_url := layers[len(layers)-1]
						objectIDNew := SaveArcGIS(layer_url, features, objectID)

						if objectIDNew >= 1{
							dataGIS := map[string]interface{}{}
							dataGIS[connection.ObjectIdField] = objectIDNew
							DB.DB.Table(table).Where(Identity+" = ?", rowId).Update(dataGIS)

						}
					}
				}

			}



		}
	}



	return data

}
func ArcGISAfterDelete(Model interface{}, data []map[string]interface{}, id string, gridID string, Identity string) []map[string]interface{} {


		var connectionPre models.ArcgisConnection
		DB.DB.Where("local_grid = ?", gridID).Find(&connectionPre)

		if connectionPre.Connection != "" {

			DB.DB.Unscoped().Where("id = ?", id).Find(Model)
			deletedData := map[string]interface{}{}

			data__, _ := json.Marshal(Model)
			json.Unmarshal(data__, &deletedData)

			var connection Connection

			json.Unmarshal([]byte(connectionPre.Connection), &connection)

			objectIDSTRING := fmt.Sprintf("%v", deletedData[connection.ObjectIdField])

			objectID, _ := strconv.ParseInt(objectIDSTRING, 10, 64)


			if objectID >= 1{

				layers := []string{}
				json.Unmarshal([]byte(connectionPre.Layer), &layers)


				if len(layers) >= 1 {
					layer_url := layers[len(layers)-1]
					DeleteArcGIS(layer_url, objectID)
					//fmt.Println("objectID", objectID, "objectID")
				}

			}



		}




	return data

}
func GetArcGISTokenForBackEnd() models.ArcGisResponse {



	url := "https://gismap.mris.mn/arcgis/tokens/generateToken"

	payload := strings.NewReader("username=geosystem&password=geosystem2020&client=requestip&f=json&expiration=120")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)




	response := models.ArcGisResponse{}
	json.Unmarshal(body, &response)

	return response
}
func SaveArcGIS(reQuestUrl string, features string, objectID int64) int64{



	reQuestUrl = reQuestUrl+"/applyEdits"

	data := url.Values{}
	data.Set("f", "pjson")
	token := GetArcGISTokenForBackEnd()
	data.Set("token", token.Token)

	//fmt.Println(objectID, "objectID")
	//fmt.Println(reQuestUrl, "reQuestUrl")

	if objectID >= 1{

		data.Set("updates", features)
	} else {
		data.Set("adds", features)
	}

	//
	//fmt.Println(features)
	//fmt.Println(data.Encode())

	req, _ := http.NewRequest("POST", reQuestUrl,  bytes.NewBufferString(data.Encode()))

	req.Header.Add("content-type", "application/x-www-form-urlencoded; param=value")
	req.Header.Add("cache-control", "no-cache")

	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	res, error_error := client.Do(req)
	//fmt.Println("====================================")
	//fmt.Println("CLIENT====================================")
	fmt.Println(error_error)
	//fmt.Println("====================================")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)




	var result ArcGISResult
	json.Unmarshal([]byte(body), &result)



	if len(result.AddResults) >= 1 {
		if result.AddResults[0].Success == true{
			return  result.AddResults[0].ObjectId
		} else {
			fmt.Println(string(body))
			return 0
		}

	} else if len(result.UpdateResults) >= 1 {
		if result.UpdateResults[0].Success == true{
			return  result.UpdateResults[0].ObjectId
		} else {
			fmt.Println(string(body))
			return 0
		}
	} else {
		fmt.Println(string(body))
		return 0
	}
}
func DeleteArcGIS(reQuestUrl string, objectID int64){

	reQuestUrl = reQuestUrl+"/deleteFeatures"


	data := url.Values{}


	if objectID >= 1{
		token := GetArcGISTokenForBackEnd()
		data.Set("token", token.Token)
		data.Add("f", "pjson")
		data.Add("objectIds", fmt.Sprintf("%v", objectID))

		req, _ := http.NewRequest("POST", reQuestUrl,  bytes.NewBufferString(data.Encode()))

		req.Header.Add("content-type", "application/x-www-form-urlencoded; param=value")
		req.Header.Add("cache-control", "no-cache")


		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		//body, _ := ioutil.ReadAll(res.Body)


		//fmt.Println(string(body))

	}




}
type ArcGISResult struct {
	AddResults []ObjectResult `json:"addResults"`
	UpdateResults []ObjectResult `json:"updateResults"`
}
type ObjectResult struct {
	Success bool   `json:"success"`
	ObjectId  int64 `json:"objectId"`
}
type Connection struct {

	Connections []map[string]string `json:"connections"`
	ObjectIdField string  `json:"objectIdField"`
	GeoJsonField string  `json:"geoJsonField"`
	LayerType string  `json:"layerType"`
}
type Coordinate struct {

	Lat string  `json:"lat"`
	Lng string  `json:"lng"`
}
type CoordinateFloat struct {
	Lat float64  `json:"lat"`
	Lng float64  `json:"lng"`
}

type GeoJSon struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			string `json:""`
		} `json:"properties"`
		Geometry struct {
			Type        string       `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}
type ArcGIS struct {
	Data interface{}  `json:"data"`
}
type Schema struct {
	Schema string  `json:"schema"`
}
type SchemaData struct {
	Schema []map[string]interface{}  `json:"schema"`
}
type Field struct {
	Field string                 `json:"field"`
	Props map[string]interface{} `json:"props"`
}
type Data struct {
	EditMode bool             `json:"editMode" form:"editMode"`
	Model    FormModel `json:"model" form:"model"`
}
type FormModel struct {
	LocalForm   int    `json:"local_form"`
}