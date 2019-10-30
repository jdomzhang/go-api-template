package vo

import "{{name}}/src/models/primitive"

// DataByPage contains pagination information and data
type DataByPage struct {
	primitive.Pagination
	Items interface{} `json:"items"`
}