package validationCaller

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/khankhulgun/khankhulgun/models/form/validations"
)

func GetMessages(schema_id string) map[string][]string {

	switch schema_id {
	case "crud_form":
		return validations.GetCrudFromMessages()
	case "analytic_form":
		return validations.AnalyticFormMessages()

	case "menu_form":
		return validations.GetMenuFormMessages()

	case "user_form":
		return validations.GetUserFormMessages()

	case "user_profile":
		return validations.GetUserProfileMessages()

	case "user_password":
		return validations.GetUserPasswordMessages()

	}
	return govalidator.MapData{}

}
