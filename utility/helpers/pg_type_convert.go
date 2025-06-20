package helpers

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: s != ""}
}


func ToPgDate(t time.Time) pgtype.Date {
	return pgtype.Date{Time: t, Valid: !t.IsZero()}
}

func PgDateToString(date pgtype.Date) string {
	if !date.Valid {
		return ""
	}
	return date.Time.Format("2006-01-02") // or use "02 Jan 2006"
}

func PgTextToString(t pgtype.Text) string {
	if !t.Valid {
		return ""
	}
	return t.String
}

func PgTimeToString(ts pgtype.Timestamp) string {
	if !ts.Valid {
		return ""
	}
	return ts.Time.Format(time.RFC3339)
}

func SqlNullTime(t time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{Time: t, Valid: true}
}

// ToPgTime converts a Go time.Time to pgtype.Timestamptz (PostgreSQL compatible)
func ToPgTime(t time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  t,
		Valid: true,
	}
}
