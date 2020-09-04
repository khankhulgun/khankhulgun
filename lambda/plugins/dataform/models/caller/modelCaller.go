package caller
import "github.com/khankhulgun/khankhulgun/models/form"
func GetMODEL(schema_id string) (string, interface{}) {

switch schema_id {

 case "crud_form":
return "id", new(form.CrudFrom)

 case "analytic_form":
return "id", new(form.AnalyticForm)

 case "menu_form":
return "id", new(form.MenuForm)

 case "user_form":
return "id", new(form.UserForm)

 case "user_profile":
return "id", new(form.UserProfile)

 case "user_password":
return "id", new(form.UserPassword)

} 
return "id", new(interface{})

}