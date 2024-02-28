package baize

import (
	"database/sql/driver"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func NewTime() *Time {
	return &Time{Time: time.Now()}
}

// MarshalJSON implements json.Marshaler.
func (t *Time) MarshalJSON() ([]byte, error) {
	//do your serializing here
	seconds := t.UnixMilli()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	t.Time = time.Unix(gconv.Int64(string(b)), 0)
	return nil
}
func (t *Time) ToString() string {
	return t.Format("2006-01-02 15:04:05")
}

func (t *Time) Value() (driver.Value, error) {
	return t.Time, nil

}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}
