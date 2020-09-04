package models

type Analytic struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	Source string `gorm:"column:source" json:"source"`
	Title  string `gorm:"column:title" json:"title"`
}
func (a *Analytic) TableName() string {
	return "analytic"
}

type AnalyticFilter struct {
	AnalyticID  int    `gorm:"column:analytic_id" json:"analytic_id"`
	FilterField string `gorm:"column:filter_field" json:"filter_field"`
	ID          int    `gorm:"column:id;primary_key" json:"id"`
	NameFeild   string `gorm:"column:name_feild" json:"name_feild"`
	Source      string `gorm:"column:source" json:"source"`
	Title       string `gorm:"column:title" json:"title"`
	ValueFeild  string `gorm:"column:value_feild" json:"value_feild"`
	SourceParentFelld string `gorm:"column:source_parent_felld" json:"source_parent_felld"`
	ParentFeild       string `gorm:"column:parent_feild" json:"parent_feild"`
}

func (a *AnalyticFilter) TableName() string {
	return "analytic_filter"
}

type AnalyticRowsColumn struct {
	AnalyticID         int    `gorm:"column:analytic_id" json:"analytic_id"`
	ColOrRow           string `gorm:"column:col_or_row" json:"col_or_row"`
	CompareFeild       string `gorm:"column:compare_feild" json:"compare_feild"`
	Comparison         string `gorm:"column:comparison" json:"comparison"`
	DataCondition      string `gorm:"column:data_condition" json:"data_condition"`
	ID                 int    `gorm:"column:id;primary_key" json:"id"`
	NameFeild          string `gorm:"column:name_feild" json:"name_feild"`
	SourceCompareField string `gorm:"column:source_compare_field" json:"source_compare_field"`
	SourceTable          string `gorm:"column:table_name" json:"table_name"`
	Title              string `gorm:"column:title" json:"title"`
	Type               string `gorm:"column:type" json:"type"`
}

func (a *AnalyticRowsColumn) TableName() string {
	return "analytic_rows_columns"
}

type AnalyticRangeFilter struct {
	AnalyticID  int `gorm:"column:analytic_id" json:"analytic_id"`
	EndValue    int    `gorm:"column:end_value" json:"end_value"`
	FilterField string `gorm:"column:filter_field" json:"filter_field"`
	ID          int    `gorm:"column:id;primary_key" json:"id"`
	StartValue  int    `gorm:"column:start_value" json:"start_value"`
	Title       string `gorm:"column:title" json:"title"`
}

func (a *AnalyticRangeFilter) TableName() string {
	return "analytic_range_filter"
}

type AnalyticRangeRowColumn struct {
	AnalyticID         int    `gorm:"column:analytic_id" json:"analytic_id"`
	ColOrRow           string `gorm:"column:col_or_row" json:"col_or_row"`
	Comparison         string `gorm:"column:comparison" json:"comparison"`
	DataCondition      string `gorm:"column:data_condition" json:"data_condition"`
	EndField           string `gorm:"column:end_field" json:"end_field"`
	ID                 int    `gorm:"column:id;primary_key" json:"id"`
	SourceCompareField string `gorm:"column:source_compare_field" json:"source_compare_field"`
	StartField         string `gorm:"column:start_field" json:"start_field"`
	SourceTable          string `gorm:"column:table_name" json:"table_name"`
	Title              string `gorm:"column:title" json:"title"`
	Type               string `gorm:"column:type" json:"type"`
}

func (a *AnalyticRangeRowColumn) TableName() string {
	return "analytic_range_row_columns"
}



type Reguest struct {
	AnalyticType int         `json:"analytic_type"`
	Col          interface{} `json:"col"`
	Row          interface{}         `json:"row"`
	Aggregation  string      `json:"aggregation"`
	Filters      []struct {
		AnalyticID  int         `json:"analytic_id"`
		FilterField string      `json:"filter_field"`
		ID          int         `json:"id"`
		NameFeild   string      `json:"name_feild"`
		Source      string      `json:"source"`
		Title       string      `json:"title"`
		ValueFeild  string      `json:"value_feild"`
		Value       interface{} `json:"value"`
	} `json:"filters"`
	RangeFilters []struct {
		AnalyticID  int    `json:"analytic_id"`
		EndValue    int    `json:"end_value"`
		FilterField string `json:"filter_field"`
		ID          int    `json:"id"`
		StartValue  int    `json:"start_value"`
		Title       string `json:"title"`
		Value       []int  `json:"value"`
	} `json:"range_filters"`
}