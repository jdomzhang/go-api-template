package primitive

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// Criteria contains the condition for building sql query
type Criteria struct {
	Name   string      `json:"name" example:"id"`
	Op     string      `json:"op" example:"<>"`
	Value  interface{} `json:"value" example:"0"`
	Value2 interface{} `json:"value2"`
}

func (obj *Criteria) getWhereClause() string {
	if obj.Name == "" || obj.Value == nil {
		return ""
	}

	op := obj.Op

	if strings.Index(op, "(") != -1 || strings.Index(op, ")") != -1 || strings.Index(op, "'") != -1 || strings.Index(op, " ") != -1 {
		op = "="
	}

	// avoid sql injection
	value := fmt.Sprintf("%v", obj.Value)
	value = strings.Replace(value, "'", "", -1)

	// avoid sql injection
	value2 := fmt.Sprintf("%v", obj.Value2)
	value2 = strings.Replace(value2, "'", "", -1)

	switch op {
	case "contains":
		return fmt.Sprintf(`%s like '%v'`, gorm.ToDBName(obj.Name), "%%"+value+"%%")
	case "between":
		return fmt.Sprintf(`%s between '%v' and '%v'`, gorm.ToDBName(obj.Name), value, value2)
	case "in":
		if strings.Index(value, "(") == -1 {
			value = fmt.Sprintf(`(%v)`, value)
		}
		return fmt.Sprintf(`%s in %v`, gorm.ToDBName(obj.Name), value)
	case "":
		op = "="
		fallthrough
	case "=":
		fallthrough
	case ">":
		fallthrough
	case ">=":
		fallthrough
	case "<":
		fallthrough
	case "<=":
		fallthrough
	default:
		return fmt.Sprintf(`%s %s '%v'`, gorm.ToDBName(obj.Name), op, value)
	}
}
