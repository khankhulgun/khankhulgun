package models


type LoginData struct {
	Token      string    `json:"token"`
	Path      string    `json:"path"`
	Status      bool    `json:"status"`
	Data      interface{}    `json:"data"`
}


type Unauthorized struct {
	Error      string    `json:"error"`
	Status      bool    `json:"status"`
}


