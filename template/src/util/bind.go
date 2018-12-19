package util

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// BindWithFilter will bind data to target, but will be filtered by filter
func BindWithFilter(c *gin.Context, filter, target interface{}) error {
	// set it first
	if err := setA2B(target, filter); err != nil {
		return err
	}

	// bind to filter obj
	if err := c.ShouldBind(filter); err != nil {
		return err
	}

	// copy back
	return setA2B(filter, target)
}

func setA2B(objA, objB interface{}) error {
	// convert objA to json
	if jsonA, err := json.Marshal(objA); err != nil {
		return err
	} else {
		// convert json to objB
		if err := json.Unmarshal(jsonA, objB); err != nil {
			return err
		}
	}

	return nil
}
