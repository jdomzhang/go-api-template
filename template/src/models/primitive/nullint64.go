package primitive

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// NullInt64 is alias of sql.NullInt64
type NullInt64 sql.NullInt64

var nullBytes = []byte("null")

// Scan implements the driver Scanner interface
func (n *NullInt64) Scan(value interface{}) error {
	if value == nil {
		n.Int64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	n.Int64, _ = value.(int64)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

// MarshalJSON marshals NullInt64 without valid field
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullBytes, nil
	}
	return []byte(strconv.FormatInt(int64(n.Int64), 10)), nil
}

// UnmarshalJSON unmashals NullInt64
func (n *NullInt64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		n.Valid = false
		n.Int64 = 0
		return nil
	}
	var x int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	n.Int64 = x
	n.Valid = true
	return nil
}
