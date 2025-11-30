package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type (
	JSON[T any] struct {
		Val T
	}
)

func (v *JSON[T]) Scan(src interface{}) error {
	switch t := src.(type) {
	case string:
	case []uint8:
		return json.NewDecoder(bytes.NewBuffer([]byte(t))).Decode(&v.Val)
	// case []byte:
	// 	return json.NewDecoder(bytes.NewBuffer(t)).Decode(v)
	case nil:
		return nil
	default:
		return fmt.Errorf("can't convert %T to []byte", v)
	}
	return nil
}

func (v JSON[T]) Value() (driver.Value, error) {
	j, err := json.Marshal(v.Val)
	if err != nil {
		return nil, err
	}

	return driver.Value([]byte(j)), nil
}
