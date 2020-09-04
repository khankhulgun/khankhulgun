package tools

import (
	"math"
	"github.com/jinzhu/gorm"
)


type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	ShowSQL bool
}

// Paginator
type Paginator struct {
	//To int         `json:"to"`
	Total int         `json:"total"`
	LastPage   int         `json:"last_page"`
	Data     interface{} `json:"data"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	CurrentPage        int         `json:"current_page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

type gridResult struct {
	CurrentPage string `json:"current_page"`
	Data interface{} `json:"data"`
	//Data map[string]map[string]string `json:"data"`
	FirstPageUrl string `json:"first_page_url"`
	From int `json:"from"`
	LastPage int `json:"last_page"`
	LastPageUrl string `json:"last_page_url"`
	NextPageUrl string `json:"next_page_url"`
	Path string `json:"path"`
	PerPage string `json:"per_page"`
	PrevPageUrl string `json:"prev_page_url"`
	To int `json:"to"`
	Total int `json:"total"`
}

// Paging
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}


	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(result)
	<-done


	paginator.Total = count
	paginator.Data = result
	paginator.CurrentPage = p.Page


	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.LastPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.LastPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}