package grid

import "time"

type KrudGrid struct {
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Form      int        `gorm:"column:form" json:"form"`
	Grid      int        `gorm:"column:grid" json:"grid"`
	Actions   string     `gorm:"column:actions" json:"actions"`
	Template  string     `gorm:"column:template" json:"template"`
	Title     string     `gorm:"column:title" json:"title"`
}

//  TableName sets the insert table name for this struct type
func (k *KrudGrid) TableName() string {
	return "krud"
}
func (k *KrudGrid) GetCondition() string {
	return ""
}
func (k *KrudGrid) GetFilters() map[string]string {

	filters := map[string]string{

	}

	return filters
}
func (k *KrudGrid) GetColumns() map[int]map[string]string {

	columns := make(map[int]map[string]string)

	columns[1] = map[string]string{"column": "title", "label": "Гарчиг"}



	return columns
}
func (k *KrudGrid) GetAggergations() string {

	aggergations := ""

	return aggergations
}
