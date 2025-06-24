package utils

import (
	"fmt"
	"time"
)

var (
	JapanLocation = time.FixedZone("Asia/Tokyo", 9*60*60)
)

// format datetime yyyy-mm-dd hh:mm:ss
func FormatDateTime(date time.Time) string {
	return date.In(JapanLocation).Format(time.DateTime)
}

// format datetime yyyy-mm-dd
func FormatDate(date time.Time) string {
	return date.In(JapanLocation).Format(time.DateOnly)
}

// format datetime hh:mm:ss
func FormatTime(date time.Time) string {
	return date.In(JapanLocation).Format(time.TimeOnly)
}

// parseFlexibleDate parses dates in multiple formats (supports both "-" and "/" separators)
func ParseFlexibleDate(dateStr string) (time.Time, error) {
	// List of supported date formats
	formats := []string{
		"2006/01/02", // Format with slashes (e.g., "2021/04/06")
		"2006-01-02", // Format with dashes (e.g., "2021-04-06")
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date '%s' in any supported format", dateStr)
}

func BeginningOfNextDayJST(t *time.Time) (time.Time, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}

	tInJST := t.In(jst)
	return time.Date(tInJST.Year(), tInJST.Month(), tInJST.Day()+1, 0, 0, 0, 0, jst), nil
}

func ParseToTimePtr(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}

	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, dateStr)
		if err == nil {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("invalid date format: %s", dateStr)
}
