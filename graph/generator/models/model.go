package models


type GqlTable struct {
	Table     string `json:"table"`
	CheckAuth struct {
		IsLoggedIn bool  `json:"isLoggedIn"`
		Roles      []int `json:"roles"`
	} `json:"checkAuth"`
	HiddenColumns []string `json:"hidden_columns"`
}