package validations

import "github.com/thedevsaddam/govalidator"


func GetUserFormRules() map[string][]string{
	return  govalidator.MapData{

		"login": []string{"required"},
		"password": []string{"required",},
		"role": []string{"required",},
	}
}
func GetUserFormMessages() map[string][]string{
	return govalidator.MapData{

		"login": []string{"required:Талбарыг бөглөнө үү!"},
		"password": []string{"required:Талбарыг бөглөнө үү!",},
		"role": []string{"required:Хандах эрхээ сонгоно уу",},
	}
}
