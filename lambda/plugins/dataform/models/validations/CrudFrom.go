package validations

import "github.com/thedevsaddam/govalidator"


func GetCrudFromRules() map[string][]string{
	return  govalidator.MapData{
		
"title": []string{"required",},
	}
}
func GetCrudFromMessages() map[string][]string{
	return govalidator.MapData{

"title": []string{"required:Талбарыг бөглөнө үү!",},
	}
}
