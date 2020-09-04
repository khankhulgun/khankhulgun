package models

type LineRequest struct {
	Axis []struct {
		Name       string `json:"name"`
		Title      string `json:"title"`
		Type       string `json:"type"`
		Table      string `json:"table"`
		Alias      string `json:"alias"`
		Output     bool   `json:"output"`
		SortType   string `json:"sortType"`
		SortOrder  int    `json:"sortOrder"`
		GroupBy    bool   `json:"groupBy"`
		GroupOrder int    `json:"groupOrder"`
		Aggregate  string `json:"aggregate"`
	} `json:"axis"`
	Lines []struct {
		Name       string `json:"name"`
		Title      string `json:"title"`
		Type       string `json:"type"`
		Table      string `json:"table"`
		Alias      string `json:"alias"`
		Output     bool   `json:"output"`
		SortType   string `json:"sortType"`
		SortOrder  int    `json:"sortOrder"`
		GroupBy    bool   `json:"groupBy"`
		GroupOrder int    `json:"groupOrder"`
		Aggregate  string `json:"aggregate"`
	} `json:"lines"`
}
type CountRequest struct {
	CountFields []struct {
		Name       string `json:"name"`
		Title      string `json:"title"`
		Type       string `json:"type"`
		Table      string `json:"table"`
		Alias      string `json:"alias"`
		Output     bool   `json:"output"`
		SortType   string `json:"sortType"`
		SortOrder  int    `json:"sortOrder"`
		GroupBy    bool   `json:"groupBy"`
		GroupOrder int    `json:"groupOrder"`
		Aggregate  string `json:"aggregate"`
		Editing    bool   `json:"editing"`
	} `json:"countFields"`
}
type PieRequest struct {
	Value []struct {
		Name       string `json:"name"`
		Title      string `json:"title"`
		Type       string `json:"type"`
		Table      string `json:"table"`
		Alias      string `json:"alias"`
		Output     bool   `json:"output"`
		SortType   string `json:"sortType"`
		SortOrder  int    `json:"sortOrder"`
		GroupBy    bool   `json:"groupBy"`
		GroupOrder int    `json:"groupOrder"`
		Aggregate  string `json:"aggregate"`
	} `json:"value"`
	Title []struct {
		Name       string `json:"name"`
		Title      string `json:"title"`
		Type       string `json:"type"`
		Table      string `json:"table"`
		Alias      string `json:"alias"`
		Output     bool   `json:"output"`
		SortType   string `json:"sortType"`
		SortOrder  int    `json:"sortOrder"`
		GroupBy    bool   `json:"groupBy"`
		GroupOrder int    `json:"groupOrder"`
		Aggregate  string `json:"aggregate"`
	} `json:"title"`
	Filter []interface{} `json:"filter"`
}
type TableRequest struct {
	Values []struct {
		Name       string `json:"name"`
		Title      string `json:"title"`
		Type       string `json:"type"`
		Table      string `json:"table"`
		Alias      string `json:"alias"`
		Output     bool   `json:"output"`
		SortType   string `json:"sortType"`
		SortOrder  int    `json:"sortOrder"`
		GroupBy    bool   `json:"groupBy"`
		GroupOrder int    `json:"groupOrder"`
		Aggregate  string `json:"aggregate"`
	} `json:"values"`
}