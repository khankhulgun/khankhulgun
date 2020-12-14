package grid

import "time"

var _ = time.Time{}

type CrudLog struct {
	Action    string     `gorm:"column:action" json:"action"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	FirstName string     `gorm:"column:first_name" json:"first_name"`
	ID        int        `gorm:"column:id" json:"id"`
	Input     string     `gorm:"column:input" json:"input"`
	IP        string     `gorm:"column:ip" json:"ip"`
	LastName  string     `gorm:"column:last_name" json:"last_name"`
	Name      string     `gorm:"column:name" json:"name"`
	RowID     string     `gorm:"column:row_id" json:"row_id"`
	UserAgent string     `gorm:"column:user_agent" json:"user_agent"`
}

type MainTableCrudLog struct{
	Action    string     `gorm:"column:action" json:"action"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	Input     string     `gorm:"column:input" json:"input"`
	IP        string     `gorm:"column:ip" json:"ip"`
	RowID     string     `gorm:"column:row_id" json:"row_id"`
	SchemaID  int        `gorm:"column:schemaId" json:"schemaId"`
	UserAgent string     `gorm:"column:user_agent" json:"user_agent"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
}

func (m *MainTableCrudLog) TableName() string {
	return "crud_log"
}

//  TableName sets the insert table name for this struct type
func (d *CrudLog) TableName() string {
	return "ds_crud_log"
}

func (v *CrudLog) GetCondition() string {
	return ""
}
func (v *CrudLog) GetFilters() map[string]string {


	filters := map[string]string{


		"action":"Select",

		"created_at":"DateRange",

		"user_id":"Select",

		"schemaId":"Select",

		"row_id":"Select",

	}

	return filters
}
func (v *CrudLog) GetColumns() map[int]map[string]string{



	columns := make(map[int]map[string]string)


	columns[1] = map[string]string{"column":"name","label":"Маягт"}

	columns[2] = map[string]string{"column":"last_name","label":"Хэрэглэгчийн овог"}

	columns[3] = map[string]string{"column":"first_name","label":"Хэрэглэгчийн нэр"}

	columns[4] = map[string]string{"column":"ip","label":"IP хаяг"}

	columns[5] = map[string]string{"column":"user_agent","label":"хэрэглэгчийн агент"}

	columns[6] = map[string]string{"column":"action","label":"Үйлдэл"}

	columns[7] = map[string]string{"column":"row_id","label":"бүртгэлийн ID"}

	columns[8] = map[string]string{"column":"input","label":"Оролтын мэдээлэл"}

	columns[9] = map[string]string{"column":"created_at","label":"Огноо"}



	return columns
}
func (v *CrudLog) GetAggergations() string {


	aggergations := ""

	return aggergations
}