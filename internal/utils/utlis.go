package utils

import (
	"fmt"
	"strconv"
)

type NullableInt struct {
	Present bool
	Null    bool
	Value   int64
}

func (m NullableInt) GoString() string {
	if !m.Present {
		return "<missing>"
	}
	if m.Null {
		return "<null>"
	}
	return fmt.Sprint(m.Value)
}

func (m *NullableInt) UnmarshalJSON(data []byte) error {
	s := string(data)
	m.Present = true
	if s == "null" {
		m.Null = true
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	m.Value = v
	return err
}
