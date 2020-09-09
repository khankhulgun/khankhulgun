package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Krud struct {
	ID        int     `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Title    string    `gorm:"column:title" json:"title"`
	Template    string    `gorm:"column:template" json:"template"`
	Grid    int    `gorm:"column:grid" json:"grid"`
	Form    int    `gorm:"column:form" json:"form"`
	Actions    string    `gorm:"column:actions" json:"actions"`

}

//  TableName sets the insert table name for this struct type
func (v *Krud) TableName() string {
	return "krud"
}
type KrudTemplate struct {
	gorm.Model
	TemplateName    string    `gorm:"column:template_name" json:"template_name"`


}

//  TableName sets the insert table name for this struct type
func (v *KrudTemplate) TableName() string {
	return "krud_templates"
}
type CrudLog struct {
	ID        int     `gorm:"column:id;primary_key" json:"id"`
	UserId        int64     `gorm:"column:user_id;" json:"user_id"`
	Ip        string     `gorm:"column:ip;" json:"ip"`
	UserAgent        string     `gorm:"column:user_agent;" json:"user_agent"`
	Action        string     `gorm:"column:action;" json:"action"`
	SchemaId        int64     `gorm:"column:schemaId;" json:"schemaId"`
	RowId        string     `gorm:"column:row_id;" json:"row_id"`
	Input        string     `gorm:"column:input;type:LONGTEXT" json:"input"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
func (v *CrudLog) TableName() string {
	return "crud_log"
}
type CrudLogMSSQL struct {
	ID        int     `gorm:"column:id;primary_key" json:"id"`
	UserId        int64     `gorm:"column:user_id;" json:"user_id"`
	Ip        string     `gorm:"column:ip;" json:"ip"`
	UserAgent        string     `gorm:"column:user_agent;" json:"user_agent"`
	Action        string     `gorm:"column:action;" json:"action"`
	SchemaId        int64     `gorm:"column:schemaId;" json:"schemaId"`
	RowId        string     `gorm:"column:row_id;" json:"row_id"`
	Input        string     `gorm:"column:input;type:NTEXT" json:"input"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
func (v *CrudLogMSSQL) TableName() string {
	return "crud_log"
}
