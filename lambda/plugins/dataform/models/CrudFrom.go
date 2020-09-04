package form

import "time"

type CrudFrom struct {
	Actions   string     `gorm:"column:actions" json:"actions"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Form      int        `gorm:"column:form" json:"form"`
	Grid      int        `gorm:"column:grid" json:"grid"`
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	Template  string     `gorm:"column:template" json:"template"`
	Title     string     `gorm:"column:title" json:"title"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (k *CrudFrom) TableName() string {
	return "krud"
}
func (a *CrudFrom) GetSubForms() []map[string]interface{} {
	subForms := []map[string]interface{}{}
	return subForms
}
