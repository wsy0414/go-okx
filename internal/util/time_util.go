package util

import "time"

func GetISOFormat(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
