package baize

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type List struct {
	Data []string
}

// MarshalJSON implements json.Marshaler.
func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.Data)
}

func (l *List) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.Data)
}

func (l *List) Value() (driver.Value, error) {
	join := strings.Join(l.Data, ",")
	return join, nil

}

func (l *List) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if ok {
		if len(value) == 0 {
			l.Data = make([]string, 0)
		} else {
			l.Data = strings.Split(string(value), ",")
		}

		return nil
	}
	return fmt.Errorf("can not convert %v to string", v)
}
