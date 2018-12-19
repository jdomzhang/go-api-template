package primitive

import (
	"bytes"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Pagination contains all fields for paging
type Pagination struct {
	Descending  bool       `json:"descending,omitempty" example:"true"`
	Page        uint64     `json:"page,omitempty" example:"1"`
	RowsPerPage int64      `json:"rowsPerPage,omitempty" example:"5"` // -1 means all
	SortBy      string     `json:"sortBy,omitempty" example:"id"`
	TotalItems  uint64     `json:"totalItems,omitempty"`
	Criteria    []Criteria `json:"criteria,omitempty"`
}

// GetOrderClause will get the order sql clause with current pagination value
func (obj *Pagination) GetOrderClause() string {
	result := ""
	if obj.SortBy == "" {
		// by default, sort by id desc
		return "id desc"
	}

	result = gorm.ToDBName(obj.SortBy)
	if obj.Descending {
		result = fmt.Sprintf("%s DESC", result)
	}

	return result
}

// GetWhereClause will get the where sql clause with current pagination value
func (obj *Pagination) GetWhereClause() string {
	var buf bytes.Buffer

	for _, c := range obj.Criteria {
		sql := c.getWhereClause()
		if sql == "" {
			continue
		}

		if buf.Len() > 0 {
			buf.WriteString(" and ")
		}
		buf.WriteString(sql)
	}

	return buf.String()
}
