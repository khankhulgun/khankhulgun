package models


type GqlTable struct {
	Table     string `json:"table"`
	Identity     string `json:"identity"`
	CheckAuth struct {
		IsLoggedIn bool  `json:"isLoggedIn"`
		Roles      []int `json:"roles"`
	} `json:"checkAuth"`
	HiddenColumns []string `json:"hidden_columns"`
	Actions       struct {
		Create bool `json:"create"`
		Update bool `json:"update"`
		Delete bool `json:"delete"`
	} `json:"actions"`
	Subs []SubTable `json:"subs"`
	Subscription bool `json:"subscription"`
}

type SubTable struct {
	Table           string `json:"table"`
	ConnectionField string `json:"connection_field"`
	ParentIdentity  string `json:"parent_identity"`
}

