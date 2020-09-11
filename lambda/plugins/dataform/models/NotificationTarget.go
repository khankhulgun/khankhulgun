package form

import "time"

var _ = time.Time{}

type NotificationTarget struct {
	Body          string `gorm:"column:body" json:"body"`
	Condition     string `gorm:"column:condition" json:"condition"`
	ID            int64  `gorm:"column:id;primary_key" json:"id"`
	Link          string `gorm:"column:link" json:"link"`
	SchemaID      int    `gorm:"column:schema_id" json:"schema_id"`
	TargetActions string `gorm:"column:target_actions" json:"target_actions"`
	TargetRole    int    `gorm:"column:target_role" json:"target_role"`
	Title         string `gorm:"column:title" json:"title"`
}

//  TableName sets the insert table name for this struct type
func (n *NotificationTarget) TableName() string {
	return "notification_targets"
}
func (a *NotificationTarget) GetSubForms() []map[string]interface{} {
	subForms := []map[string]interface{}{}
	return subForms }
			func (v *NotificationTarget) GetFromTypes() map[string]string{


				fields := map[string]string{

			
					"id":"Text",

					"link":"Text",

					"title":"Text",

					"body":"Textarea",

					"schema_id":"Select",

					"target_role":"Select",

					"condition":"Text",

					"target_actions":"Select",

			}

			return fields
		}
			func (v *NotificationTarget) GetFormula() string{

			return ""
		}