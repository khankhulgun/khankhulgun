package models


type GqlTable struct {
	Table     string `json:"table"`
	Identity     string `json:"identity"`
	CheckAuth struct {
		IsLoggedIn bool  `json:"isLoggedIn"`
		Roles      []int `json:"roles"`
	} `json:"checkAuth"`
	HiddenColumns []string `json:"hidden_columns"`
	Subs []SubTable `json:"subs"`
}

type SubTable struct {
	Table           string `json:"table"`
	ConnectionField string `json:"connection_field"`
	ParentIdentity  string `json:"parent_identity"`
}