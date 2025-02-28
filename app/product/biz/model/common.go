package model

import "gorm.io/gorm"

type Pagination struct {
	Page     int
	PageSize int
}

func (p Pagination) Offset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return (p.Page - 1) * p.PageSize
}

func Paginate(page Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page.Page - 1) * page.PageSize
		return db.Offset(offset).Limit(page.PageSize)
	}
}
