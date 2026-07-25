package store

import (
	"database/sql"
	"time"
)

func pointerTimeToNullTime(timePointer *time.Time) sql.NullTime {
	var t sql.NullTime
	if timePointer != nil {
		t = sql.NullTime{
			Valid: true,
			Time:  *timePointer,
		}
	} else {
		t.Valid = false
	}

	return t
}

func stringToNullString(str string) sql.NullString {
	var s sql.NullString
	if str != "" {
		s = sql.NullString{
			Valid:  true,
			String: str,
		}
	} else {
		s = sql.NullString{
			Valid: false,
		}
	}

	return s
}
