package validations

import "github.com/thedevsaddam/govalidator"


func GetNotificationTargetRules() map[string][]string{
	return  govalidator.MapData{
		
"link": []string{"required",},
"title": []string{"required",},
"body": []string{"required",},
"schema_id": []string{"required",},
"target_role": []string{"required",},
"condition": []string{"required",},
	}
}
func GetNotificationTargetMessages() map[string][]string{
	return govalidator.MapData{

"link": []string{"required:Талбарыг бөглөнө үү!",},
"title": []string{"required:Талбарыг бөглөнө үү!",},
"body": []string{"required:Талбарыг бөглөнө үү!",},
"schema_id": []string{"required:Талбарыг бөглөнө үү!",},
"target_role": []string{"required:Талбарыг бөглөнө үү!",},
"condition": []string{"required:Талбарыг бөглөнө үү!",},
	}
}
