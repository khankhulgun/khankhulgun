package grid

import "time"

var _ = time.Time{}

type NotificationTarget struct {
	ID    int64  `gorm:"column:id;primary_key" json:"id"`
	Title string `gorm:"column:title" json:"title"`
}

type NotificationTargetMainTable struct {
	Body       string `gorm:"column:body" json:"body"`
	Condition  string `gorm:"column:condition" json:"condition"`
	ID         int64  `gorm:"column:id;primary_key" json:"id"`
	Link       string `gorm:"column:link" json:"link"`
	SchemaID   int    `gorm:"column:schema_id" json:"schema_id"`
	TargetRole int    `gorm:"column:target_role" json:"target_role"`
	Title      string `gorm:"column:title" json:"title"`
}

func (n *NotificationTargetMainTable) TableName() string {
	return "notification_targets"
}

//  TableName sets the insert table name for this struct type
func (n *NotificationTarget) TableName() string {
	return "notification_targets"
}

			func (v *NotificationTarget) GetCondition() string {
				return ""
			}
			func (v *NotificationTarget) GetFilters() map[string]string {


				filters := map[string]string{

			
			}

			return filters
		}
			func (v *NotificationTarget) GetColumns() map[int]map[string]string{


			
				columns := make(map[int]map[string]string)

			
					columns[2] = map[string]string{"column":"title","label":"Нэр"}

			

			return columns
		}
			func (v *NotificationTarget) GetAggergations() string {


				aggergations := ""

			return aggergations
		}