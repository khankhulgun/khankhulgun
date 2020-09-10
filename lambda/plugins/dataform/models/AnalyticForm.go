package form

import "time"

var _ = time.Time{}

type AnalyticForm struct {
	ID     int    `gorm:"column:id" json:"id"`
	Source string `gorm:"column:source" json:"source"`
	Title  string `gorm:"column:title" json:"title"`
}

type AnalyticFilterAnalyticForm struct {
	AnalyticID        int    `gorm:"column:analytic_id" json:"analytic_id"`
	FilterField       string `gorm:"column:filter_field" json:"filter_field"`
	ID                int    `gorm:"column:id" json:"id"`
	NameFeild         string `gorm:"column:name_feild" json:"name_feild"`
	ParentFeild       string `gorm:"column:parent_feild" json:"parent_feild"`
	Source            string `gorm:"column:source" json:"source"`
	SourceParentFelld string `gorm:"column:source_parent_felld" json:"source_parent_felld"`
	Title             string `gorm:"column:title" json:"title"`
	ValueFeild        string `gorm:"column:value_feild" json:"value_feild"`
}

func (a *AnalyticFilterAnalyticForm) TableName() string {
	return "analytic_filter"
}

type AnalyticRowsColumnAnalyticForm struct {
	AnalyticID         int    `gorm:"column:analytic_id" json:"analytic_id"`
	ColOrRow           string `gorm:"column:col_or_row" json:"col_or_row"`
	CompareFeild       string `gorm:"column:compare_feild" json:"compare_feild"`
	Comparison         string `gorm:"column:comparison" json:"comparison"`
	DataCondition      string `gorm:"column:data_condition" json:"data_condition"`
	ID                 int    `gorm:"column:id" json:"id"`
	NameFeild          string `gorm:"column:name_feild" json:"name_feild"`
	SourceCompareField string `gorm:"column:source_compare_field" json:"source_compare_field"`
	SourceTable          string `gorm:"column:table_name" json:"table_name"`
	Title              string `gorm:"column:title" json:"title"`
	Type               string `gorm:"column:type" json:"type"`
}

func (a *AnalyticRowsColumnAnalyticForm) TableName() string {
	return "analytic_rows_columns"
}

type AnalyticRangeFilterAnalyticForm struct {
	AnalyticID  int    `gorm:"column:analytic_id" json:"analytic_id"`
	EndValue    int    `gorm:"column:end_value" json:"end_value"`
	FilterField string `gorm:"column:filter_field" json:"filter_field"`
	ID          int    `gorm:"column:id" json:"id"`
	StartValue  int    `gorm:"column:start_value" json:"start_value"`
	Title       string `gorm:"column:title" json:"title"`
}

func (a *AnalyticRangeFilterAnalyticForm) TableName() string {
	return "analytic_range_filter"
}

type AnalyticRangeRowColumnAnalyticForm struct {
	AnalyticID         int    `gorm:"column:analytic_id" json:"analytic_id"`
	ColOrRow           string `gorm:"column:col_or_row" json:"col_or_row"`
	Comparison         string `gorm:"column:comparison" json:"comparison"`
	DataCondition      string `gorm:"column:data_condition" json:"data_condition"`
	EndField           string `gorm:"column:end_field" json:"end_field"`
	ID                 int    `gorm:"column:id" json:"id"`
	SourceCompareField string `gorm:"column:source_compare_field" json:"source_compare_field"`
	StartField         string `gorm:"column:start_field" json:"start_field"`
	SourceTable          string `gorm:"column:table_name" json:"table_name"`
	Title              string `gorm:"column:title" json:"title"`
	Type               string `gorm:"column:type" json:"type"`
}

func (a *AnalyticRangeRowColumnAnalyticForm) TableName() string {
	return "analytic_range_row_columns"
}

type AnalyticDateFilterAnalyticForm struct {
	AnalyticID int    `gorm:"column:analytic_id" json:"analytic_id"`
	DateField  string `gorm:"column:date_field" json:"date_field"`
	ID         int    `gorm:"column:id;primary_key" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
}

func (a *AnalyticDateFilterAnalyticForm) TableName() string {
	return "analytic_date_filter"
}

//  TableName sets the insert table name for this struct type
func (a *AnalyticForm) TableName() string {
	return "analytic"
}
func (a *AnalyticForm) GetSubForms() []map[string]interface{} {
	subForms := []map[string]interface{}{
		map[string]interface{}{
			"connection_field":"analytic_id",
			"table":"analytic_filter",
			"parentIdentity":"id",
			"subIdentity":"id",
			"subForm":new([]AnalyticFilterAnalyticForm),
			"subFormModel":new(AnalyticFilterAnalyticForm),

		},
		map[string]interface{}{
			"connection_field":"analytic_id",
			"table":"analytic_rows_columns",
			"parentIdentity":"id",
			"subIdentity":"id",
			"subForm":new([]AnalyticRowsColumnAnalyticForm),
			"subFormModel":new(AnalyticRowsColumnAnalyticForm),

		},
		map[string]interface{}{
			"connection_field":"analytic_id",
			"table":"analytic_range_filter",
			"parentIdentity":"id",
			"subIdentity":"id",
			"subForm":new([]AnalyticRangeFilterAnalyticForm),
			"subFormModel":new(AnalyticRangeFilterAnalyticForm),

		},
		map[string]interface{}{
			"connection_field":"analytic_id",
			"table":"analytic_range_row_columns",
			"parentIdentity":"id",
			"subIdentity":"id",
			"subForm":new([]AnalyticRangeRowColumnAnalyticForm),
			"subFormModel":new(AnalyticRangeRowColumnAnalyticForm),

		},
		map[string]interface{}{
			"connection_field":"analytic_id",
			"table":"analytic_date_filter",
			"parentIdentity":"id",
			"subIdentity":"id",
			"subForm":new([]AnalyticDateFilterAnalyticForm),
			"subFormModel":new(AnalyticDateFilterAnalyticForm),

		},}
	return subForms }
func (v *AnalyticForm) GetFromTypes() map[string]string{


	fields := map[string]string{


		"id":"Text",

		"title":"Text",

		"source":"Text",

		"analytic_filter":"SubForm",

		"analytic_rows_columns":"SubForm",

		"analytic_range_filter":"SubForm",

		"analytic_range_row_columns":"SubForm",

		"analytic_date_filter":"SubForm",

	}

	return fields
}
func (v *AnalyticForm) GetFormula() string{

	return ""
}