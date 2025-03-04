package common

import (
	"gorm.io/gorm"
)

type TPaginate struct {
	sql         *gorm.DB
	Data        interface{} `json:"paginate"`
	CurrentPage int         `json:"current"`
	Next        int         `json:"next"`
	Previous    int         `json:"previous"`
	offset      int
}

func Paginate(sql *gorm.DB) *TPaginate {
	return &TPaginate{sql: sql}
}

func (p *TPaginate) Get(data interface{}, page int, limit int) (*TPaginate, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	p.offset = limit
	p.CurrentPage = page
	p.Next = -1
	p.Previous = -1

	var count int64

	p.sql.Offset((page * p.offset) - p.offset).Limit(p.offset).Find(data)
	p.sql.Count(&count)

	if count > int64(page*p.offset) {
		p.CurrentPage = page
		p.Next = p.CurrentPage + 1
	}

	if p.CurrentPage > 1 {
		p.Previous = p.CurrentPage - 1
	}

	p.Data = data

	return p, nil
}
