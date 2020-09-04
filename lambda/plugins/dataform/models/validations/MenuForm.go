package validations

import "github.com/thedevsaddam/govalidator"


func GetMenuFormRules() map[string][]string{
	return  govalidator.MapData{
		
"name": []string{"required",},
"schema": []string{"required",},
	}
}
func GetMenuFormMessages() map[string][]string{
	return govalidator.MapData{

"name": []string{"required:Талбарыг бөглөнө үү!",},
"schema": []string{"required:Талбарыг бөглөнө үү!",},
	}
}
