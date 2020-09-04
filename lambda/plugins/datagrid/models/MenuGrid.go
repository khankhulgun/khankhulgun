package grid

import "time"

type MenuGrid struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Schema    string    `gorm:"column:schema" json:"schema"`
	Type      string    `gorm:"column:type" json:"type"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (v *MenuGrid) TableName() string {
	return "vb_schemas"
}
func (v *MenuGrid) GetCondition() string {
	return "type = 'menu'"
}