package models

import "time"

type Role struct {
	ID        int     `gorm:"column:id;primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Name      string    `gorm:"column:name;not null;unique" json:"name"`
	DisplayName    string    `gorm:"column:display_name" json:"display_name"`
	Description    string    `gorm:"column:description" json:"description"`
	Permissions    string    `gorm:"column:permissions;type:TEXT" json:"permissions"`
	Extra    string    `gorm:"column:extra;type:TEXT" json:"extra"`
	Menu    string    `gorm:"column:menu" json:"menu"`

}

//  TableName sets the insert table name for this struct type
func (v *Role) TableName() string {
	return "roles"
}

type Permissions struct {
	DefaultMenu string `json:"default_menu"`
	Extra       struct {
		Chart       bool `json:"chart"`
		Datasourcce bool `json:"datasourcce"`
		Datasource  bool `json:"datasource"`
		Moqup       bool `json:"moqup"`
	} `json:"extra"`
	MenuID      int                       `json:"menu_id"`
	Permissions map[string]PermissionData `json:"permissions"`
}
type PermissionData struct {
	C      bool   `json:"c"`
	D      bool   `json:"d"`
	MenuID string `json:"menu_id"`
	R      bool   `json:"r"`
	Show   bool   `json:"show"`
	Title  string `json:"title"`
	U      bool   `json:"u"`
}