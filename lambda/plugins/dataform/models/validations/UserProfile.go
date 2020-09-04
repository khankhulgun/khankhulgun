package validations

import "github.com/thedevsaddam/govalidator"


func GetUserProfileRules() map[string][]string{
	return  govalidator.MapData{

		"email": []string{"required",},
		"register_number": []string{"required",},
		"first_name": []string{"required",},
		"last_name": []string{"required",},
		"birthday": []string{"required",},
		"gender": []string{"required",},
	}
}
func GetUserProfileMessages() map[string][]string{
	return govalidator.MapData{

		"email": []string{"required:Талбарыг бөглөнө үү!","unique:Давхацсан утга оруулсан байна!",},
		"register_number": []string{"required:Талбарыг бөглөнө үү!","unique:Давхацсан утга оруулсан байна!",},
		"first_name": []string{"required:Талбарыг бөглөнө үү!",},
		"last_name": []string{"required:Талбарыг бөглөнө үү!",},
		"birthday": []string{"required:Талбарыг бөглөнө үү!",},
		"gender": []string{"required:Талбарыг бөглөнө үү!",},
	}
}
