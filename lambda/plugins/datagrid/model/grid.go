package model

type RowUpdateData struct {
	Ids   []int  `json:"ids"`
	Value int    `json:"value"`
	Model string `json:"model"`
}
