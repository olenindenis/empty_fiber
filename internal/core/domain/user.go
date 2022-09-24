package domain

import (
	"errors"
	"fmt"
	"time"
)

var incompatibleTypeFound = errors.New("error: incompatible type found")

type User struct {
	ID        uint
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	CreatedAt TimeStampUnix `json:"created_at"`
	UpdatedAt TimeStampUnix `json:"updated_at"`
}

type TimeStampUnix int64

func (t *TimeStampUnix) Scan(src interface{}) error {
	switch src := src.(type) {
	case time.Time:
		*t = TimeStampUnix(src.Unix())
		return nil
	case []byte:
		str := string(src)
		var y, m, d, hr, min, s, tzh, tzm int
		var sign rune
		_, e := fmt.Sscanf(str, "%d-%d-%d %d:%d:%d",
			&y, &m, &d, &hr, &min, &s)
		if e != nil {
			return e
		}

		offset := 60 * (tzh*60 + tzm)
		if sign == '-' {
			offset = -1 * offset
		}

		loc := time.FixedZone("local-tz", offset)
		t1 := time.Date(y, time.Month(m), d, hr, min, s, 0, loc)
		*t = TimeStampUnix(t1.Unix())
		return nil
	case int64:
		*t = TimeStampUnix(src)
		return nil
	case int32:
		*t = TimeStampUnix(src)
		return nil
	default:
		fmt.Printf("Value '%s' of incompatible type '%T' found\n", src, src)
		return incompatibleTypeFound
	}
}
