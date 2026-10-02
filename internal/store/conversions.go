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

func NullStringToString(ns sql.NullString) string {
	var s string
	if ns.Valid {
		s = ns.String
	}

	return s
}

func NullTimeToPointerTime(nt sql.NullTime) *time.Time {
	var pt *time.Time
	if nt.Valid {
		t := nt.Time
		pt = &t
	}

	return pt
}
