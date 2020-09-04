package validations

import "github.com/thedevsaddam/govalidator"


func GetAnalyticFormRules() map[string][]string{
	return  govalidator.MapData{
		
"title": []string{"required",},
"source": []string{"required",},
	}
}
func GetAnalyticFormMessages() map[string][]string{
	return govalidator.MapData{

"title": []string{"required:Талбарыг бөглөнө үү!",},
"source": []string{"required:Талбарыг бөглөнө үү!",},
	}
}
