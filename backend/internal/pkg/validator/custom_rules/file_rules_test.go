package custom_rules

import (
	"reflect"
	"testing"

	"github.com/oapi-codegen/runtime/types"
)

// Mock FieldLevel for testing
type mockFieldLevel struct {
	field reflect.Value
}

func (m *mockFieldLevel) Top() reflect.Value      { return reflect.Value{} }
func (m *mockFieldLevel) Parent() reflect.Value   { return reflect.Value{} }
func (m *mockFieldLevel) Field() reflect.Value    { return m.field }
func (m *mockFieldLevel) FieldName() string       { return "" }
func (m *mockFieldLevel) StructFieldName() string { return "" }
func (m *mockFieldLevel) Param() string           { return "" }
func (m *mockFieldLevel) GetTag() string          { return "" }
func (m *mockFieldLevel) ExtractType(field reflect.Value) (reflect.Value, reflect.Kind, bool) {
	return reflect.Value{}, reflect.Invalid, false
}
func (m *mockFieldLevel) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	return reflect.Value{}, reflect.Invalid, false
}
func (m *mockFieldLevel) GetStructFieldOKAdvanced(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool) {
	return reflect.Value{}, reflect.Invalid, false
}
func (m *mockFieldLevel) GetStructFieldOK2() (reflect.Value, reflect.Kind, bool, bool) {
	return reflect.Value{}, reflect.Invalid, false, false
}
func (m *mockFieldLevel) GetStructFieldOKAdvanced2(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool, bool) {
	return reflect.Value{}, reflect.Invalid, false, false
}

// Mock implementation of types.File for testing
func createMockFile(filename string) types.File {
	var file types.File
	file.InitFromBytes([]byte("test content"), filename)
	return file
}

func TestIsCSVFile(t *testing.T) {
	tests := []struct {
		name        string
		file        types.File
		expected    bool
		description string
	}{{
		name:        "valid_csv_file",
		file:        createMockFile("test.csv"),
		expected:    true,
		description: "Should return true for valid CSV file",
	},
		{
			name:        "invalid_file_extension",
			file:        createMockFile("test.txt"),
			expected:    false,
			description: "Should return false for non-CSV file extension",
		},
		{
			name:        "no_file_extension",
			file:        createMockFile("test"),
			expected:    false,
			description: "Should return false for file with no extension",
		}, {
			name:        "short_filename",
			file:        createMockFile("a.c"),
			expected:    false,
			description: "Should return false for short filename (less than 4 characters)",
		},
		{
			name:        "empty_filename",
			file:        createMockFile(""),
			expected:    false,
			description: "Should return false for empty filename",
		},
		{
			name:        "csv_in_middle_of_filename",
			file:        createMockFile("test.csv.txt"),
			expected:    false,
			description: "Should return false when .csv is not at the end",
		},
		{
			name:        "uppercase_csv_extension",
			file:        createMockFile("test.CSV"),
			expected:    true,
			description: "Should return true for uppercase .CSV extension (case insensitive)",
		}, {
			name:        "long_filename_with_csv",
			file:        createMockFile("very_long_filename_with_data.csv"),
			expected:    true,
			description: "Should return true for long filename ending with .csv",
		},
		{
			name:        "csv_with_multiple_dots",
			file:        createMockFile("file.backup.csv"),
			expected:    true,
			description: "Should return true for filename with multiple dots ending in .csv",
		},
		{
			name:        "csv_substring_in_filename",
			file:        createMockFile("csvfile.csv"),
			expected:    true,
			description: "Should return true for filename containing csv substring",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock field level with the file
			fieldValue := reflect.ValueOf(tt.file)
			mockFL := &mockFieldLevel{field: fieldValue}

			result := IsCSVFile(mockFL)
			if result != tt.expected {
				t.Errorf("IsCSVFile() = %v, expected %v. %s", result, tt.expected, tt.description)
			}
		})
	}
}

func TestIsCSVFile_NilFile(t *testing.T) {
	// Test with nil file
	var nilFile types.File
	fieldValue := reflect.ValueOf(nilFile)
	mockFL := &mockFieldLevel{field: fieldValue}

	result := IsCSVFile(mockFL)
	if result != false {
		t.Errorf("IsCSVFile() with nil file = %v, expected false", result)
	}
}

func TestIsCSVFile_InvalidType(t *testing.T) {
	// Test with non-File type
	invalidValue := "not a file"
	fieldValue := reflect.ValueOf(invalidValue)
	mockFL := &mockFieldLevel{field: fieldValue}

	result := IsCSVFile(mockFL)
	if result != false {
		t.Errorf("IsCSVFile() with invalid type = %v, expected false", result)
	}
}

func TestIsCSVFile_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{
			name:     "exactly_4_chars_valid",
			filename: "a.csv",
			expected: true,
		},
		{
			name:     "exactly_4_chars_invalid_extension",
			filename: "a.txt",
			expected: false,
		}, {
			name:     "three_chars_filename",
			filename: "abc",
			expected: false,
		},
		{
			name:     "csv_case_sensitive",
			filename: "test.Csv",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := createMockFile(tt.filename)
			fieldValue := reflect.ValueOf(file)
			mockFL := &mockFieldLevel{field: fieldValue}

			result := IsCSVFile(mockFL)
			if result != tt.expected {
				t.Errorf("IsCSVFile() for %s = %v, expected %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestIsCSVFileTag(t *testing.T) {
	if IsCSVFileTag != "is_csv" {
		t.Errorf("IsCSVFileTag = %s, expected 'is_csv'", IsCSVFileTag)
	}
}
