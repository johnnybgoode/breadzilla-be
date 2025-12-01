package main

import (
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
)

type Text[T interface {
	encoding.TextMarshaler
}, Tp interface {
	*T
	encoding.TextUnmarshaler
}] struct {
	Val T
}
type (
	StringInterfaceMap  map[string]interface{}
	StringInterfaceMaps []StringInterfaceMap
)

func (t Text[T, Tp]) Value() (driver.Value, error) {
	return t.Val.MarshalText()
}

func (t *Text[T, Tp]) Scan(value any) error {
	switch x := value.(type) {
	case string:
		v := Tp(&t.Val)
		return v.UnmarshalText([]byte(x))
	case []byte:
		v := Tp(&t.Val)
		return v.UnmarshalText(x)
	case nil:
		return nil
	default:
		return fmt.Errorf("cannot scan type %T: %v", value, value)
	}
}

func (m *StringInterfaceMaps) Scan(src interface{}) error {
	var source []byte

	switch src := src.(type) {
	case []uint8:
		source = []byte(src)
	case nil:
		return nil
	default:
		return errors.New("incompatible type for StringInterfaceMap")
	}

	_m := make(StringInterfaceMaps, len(source))

	err := json.Unmarshal(source, &_m)
	if err != nil {
		return err
	}

	*m = _m
	return nil
}

func (m StringInterfaceMaps) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}


