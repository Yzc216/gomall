package repo

import (
	"gorm.io/gorm"
	"time"
)

type Pagination struct {
	Page     int
	PageSize int
}

func (p Pagination) Offset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	pageSize := p.PageSize
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return (p.Page - 1) * p.PageSize
}

func Paginate(page Pagination) func(db *gorm.DB) *gorm.DB {
	if page.Page < 1 {
		page.Page = 1
	}
	pageSize := page.PageSize
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return func(db *gorm.DB) *gorm.DB {
		offset := (page.Page - 1) * page.PageSize
		return db.Offset(offset).Limit(page.PageSize)
	}
}

type QueryBuilder func(*gorm.DB) *gorm.DB

type SPUFilter struct {
	ShopID      uint64
	Brand       string
	Status      int8
	MinPrice    float64 // 基于SKU最低价
	MaxPrice    float64
	CreateStart time.Time
	CreateEnd   time.Time
	Keyword     string
	OrderBy     string // 排序字段示例："sales_desc"/"price_asc"
}

type Media struct {
	MainImages []string
	Video      string
}
