package utils

import (
	"testing"
)

func TestArrayToMap(t *testing.T) {
	type Object struct {
		ID   int
		Name string
	}

	type testCase[T1 any, T2 comparable] struct {
		input     []T1
		transform func(T1) T2
		expected  map[T2]T1
	}

	tests := []testCase[Object, int]{
		{
			input: []Object{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
			transform: func(obj Object) int { return obj.ID },
			expected: map[int]Object{
				1: {ID: 1, Name: "Alice"},
				2: {ID: 2, Name: "Bob"},
				3: {ID: 3, Name: "Charlie"},
			},
		},
		{
			input:     []Object{},
			transform: func(obj Object) int { return obj.ID },
			expected:  map[int]Object{},
		},
		{
			input: []Object{
				{ID: 1, Name: "Alice"},
				{ID: 1, Name: "Duplicate Alice"},
				{ID: 2, Name: "Bob"},
			},
			transform: func(obj Object) int { return obj.ID },
			expected: map[int]Object{
				1: {ID: 1, Name: "Duplicate Alice"}, // Duplicate keys are overwritten.
				2: {ID: 2, Name: "Bob"},
			},
		},
	}

	for _, tc := range tests {
		result := ArrayToMap(tc.input, tc.transform)
		if len(result) != len(tc.expected) {
			t.Errorf("Expected map length %d, got %d", len(tc.expected), len(result))
		}
		for key, value := range tc.expected {
			if result[key] != value {
				t.Errorf("Expected key %v to have value %v, got %v", key, value, result[key])
			}
		}
	}
}
