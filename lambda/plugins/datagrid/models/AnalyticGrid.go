package grid

import "time"

var _ = time.Time{}

type AnalyticGrid struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	Source string `gorm:"column:source" json:"source"`
	Title  string `gorm:"column:title" json:"title"`
}

type AnalyticMainTable514 struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	Source string `gorm:"column:source" json:"source"`
	Title  string `gorm:"column:title" json:"title"`
}

func (a *AnalyticMainTable514) TableName() string {
	return "analytic"
}

//  TableName sets the insert table name for this struct type
func (a *AnalyticGrid) TableName() string {
	return "analytic"
}

			func (v *AnalyticGrid) GetCondition() string {
				return ""
			}
			func (v *AnalyticGrid) GetFilters() map[string]string {


				filters := map[string]string{

			
			}

			return filters
		}
			func (v *AnalyticGrid) GetColumns() map[int]map[string]string{


			
				columns := make(map[int]map[string]string)

			
					columns[0] = map[string]string{"column":"id","label":""}

					columns[1] = map[string]string{"column":"title","label":"Нэр"}

					columns[2] = map[string]string{"column":"source","label":"Хүснэгт"}

			

			return columns
		}
			func (v *AnalyticGrid) GetAggergations() string {


				aggergations := ""

			return aggergations
		}