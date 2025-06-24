package utils

import (
	"testing"
	"time"
)

func TestFormatDateTime(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "UTC time conversion to JST format",
			input:    time.Date(2023, 5, 15, 10, 30, 45, 0, time.UTC),
			expected: "2023-05-15 19:30:45", // UTC+9 for JST
		},
		{
			name:     "JST time format",
			input:    time.Date(2023, 5, 15, 10, 30, 45, 0, JapanLocation),
			expected: "2023-05-15 10:30:45",
		},
		{
			name:     "Zero time",
			input:    time.Time{},
			expected: "0001-01-01 09:00:00", // Zero time in JST is 0001-01-01 09:00:00
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDateTime(tt.input)
			if result != tt.expected {
				t.Errorf("FormatDateTime() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "UTC date conversion to JST format",
			input:    time.Date(2023, 5, 15, 10, 30, 45, 0, time.UTC),
			expected: "2023-05-15", // Date portion remains the same with timezone shift
		},
		{
			name:     "JST date format",
			input:    time.Date(2023, 5, 15, 10, 30, 45, 0, JapanLocation),
			expected: "2023-05-15",
		},
		{
			name:     "UTC date near midnight conversion to JST",
			input:    time.Date(2023, 5, 15, 23, 30, 45, 0, time.UTC),
			expected: "2023-05-16", // Next day in JST
		},
		{
			name:     "Zero date",
			input:    time.Time{},
			expected: "0001-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDate(tt.input)
			if result != tt.expected {
				t.Errorf("FormatDate() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatTime(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "UTC time conversion to JST format",
			input:    time.Date(2023, 5, 15, 10, 30, 45, 0, time.UTC),
			expected: "19:30:45", // UTC+9 for JST
		},
		{
			name:     "JST time format",
			input:    time.Date(2023, 5, 15, 10, 30, 45, 0, JapanLocation),
			expected: "10:30:45",
		},
		{
			name:     "UTC time at day boundary to JST",
			input:    time.Date(2023, 5, 15, 15, 0, 0, 0, time.UTC),
			expected: "00:00:00", // 15:00 UTC = 00:00 JST (next day)
		},
		{
			name:     "Zero time",
			input:    time.Time{},
			expected: "09:00:00", // Zero time in JST
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatTime(tt.input)
			if result != tt.expected {
				t.Errorf("FormatTime() = %v, want %v", result, tt.expected)
			}
		})
	}
}
