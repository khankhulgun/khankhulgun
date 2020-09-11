package handlers

import (
	"encoding/json"
	"github.com/PaesslerAG/gval"
	"github.com/khankhulgun/khankhulgun/DB"
	models2 "github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify/models"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	"bytes"
	"text/template"
	"regexp"
)

func BuildNotification (rawData []byte, schemaId int64, action string, userId int64){

	target := models.NotificationTarget{}

	DB.DB.Where("schema_id = ? AND target_actions LIKE ?", schemaId, "%"+action+"%").Find(&target)

	if(target.ID >= 1){

		user := models2.User{}

		DB.DB.Where("id = ?", userId).Find(&user)

		dataJson := new(map[string]interface{})
		json.Unmarshal(rawData, dataJson)

		var re1 = regexp.MustCompile(`'{`)
		template := re1.ReplaceAllString(target.Condition, ``)
		var re2 = regexp.MustCompile(`}'`)
		template = re2.ReplaceAllString(template, ``)
		var re3 = regexp.MustCompile(`'`)
		template = re3.ReplaceAllString(template, `"`)

		value, _ := gval.Evaluate(template, *dataJson)

		Body := Execute(dataJson, target.Body)

		if value == true {

			FCMData := models.FCMData{
				Title:target.Title,
				Body:Body,
				FirstName: user.FirstName,
				Sound:"/notification.mp3",
				Icon:config.Config.Favicon,
				Link:target.Link,
				ClickAction:config.Config.Domain+"/control#"+target.Link,
			}

			FCMNotification := models.FCMNotification{
				Title:target.Title,
				Body:Body,
				Icon:config.Config.Domain+"/"+config.Config.Favicon,
				ClickAction:config.Config.Domain+"/control#"+target.Link,
			}

			data := models.NotificationData{
				Roles:[]int{target.TargetRole},
				Data:FCMData,
				Notification:FCMNotification,
			}
			CreateNotification(data)

		}
	}

}

func Execute(data interface{}, TBody string) string {
	t := template.Must(template.New("").Parse(TBody))
	buf := bytes.Buffer{}
	t.Execute(&buf, data)
	return buf.String()
}