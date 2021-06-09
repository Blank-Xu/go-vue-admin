package db

import (
	"xorm.io/builder"
)

type Paging struct {
	Page int `json:"page" validate:"gt=0" db:"-"`
	Size int `json:"size" validate:"lt=50" db:"-"`
}

func (p *Paging) Limit() int {
	if p.Size <= 0 {
		p.Size = 10
	} else if p.Size > 50 {
		p.Size = 50
	}
	return p.Size
}

func (p *Paging) Offset() int {
	if p.Page <= 1 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

func (p *Paging) LimitOffset() (int, int) {
	return p.Limit(), p.Offset()
}

type TimeLimit struct {
	Start int64 `json:"start" db:"-"`
	End   int64 `json:"end" db:"-"`
}

func (p *TimeLimit) TimeLimitCond(col string) builder.Cond {
	cond := builder.NewCond()
	if p.Start > 0 {
		cond = cond.And(builder.Gte{col: p.Start})
	}
	if p.End > 0 {
		cond = cond.And(builder.Lte{col: p.End})
	}
	return cond
}

type RecordsResponse struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
}
