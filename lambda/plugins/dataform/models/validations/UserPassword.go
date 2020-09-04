package validations

import "github.com/thedevsaddam/govalidator"


func GetUserPasswordRules() map[string][]string{
	return  govalidator.MapData{
		"password": []string{"required",},
	}
}
func GetUserPasswordMessages() map[string][]string{
	return govalidator.MapData{
		"password": []string{"required:Талбарыг бөглөнө үү!",},
	}
}
