package bankaccount

import (
	"testing"
)

func TestAccountHolderKana_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    AccountHolderKana
		expected bool
	}{
		{
			name:     "Valid half-width Katakana",
			input:    AccountHolderKana("ｱｲｳｴｵ"),
			expected: true,
		},
		{
			name:     "Valid alphanumeric characters",
			input:    AccountHolderKana("abc123"),
			expected: true,
		},
		{
			name:     "Valid with symbols",
			input:    AccountHolderKana("abc-123()"),
			expected: true,
		},
		{
			name:     "Valid full-width Katakana (converted to half-width)",
			input:    AccountHolderKana("アイウエオ"),
			expected: true,
		},
		{
			name:     "Invalid characters",
			input:    AccountHolderKana("あいうえお"),
			expected: false,
		},
		{
			name:     "Empty string",
			input:    AccountHolderKana(""),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.IsValid(); got != tt.expected {
				t.Errorf("AccountHolderKana.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAccountNumber_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    AccountNumber
		expected bool
	}{
		{
			name:     "Valid 7-digit number",
			input:    AccountNumber("1234567"),
			expected: true,
		},
		{
			name:     "Invalid 6-digit number",
			input:    AccountNumber("123456"),
			expected: false,
		},
		{
			name:     "Invalid 8-digit number",
			input:    AccountNumber("12345678"),
			expected: false,
		},
		{
			name:     "Invalid negative number",
			input:    AccountNumber("-1234567"),
			expected: false,
		},
		{
			name:     "Invalid zero",
			input:    AccountNumber("0"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.IsValid(); got != tt.expected {
				t.Errorf("AccountNumber.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAccountKind_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    AccountKind
		expected bool
	}{
		{
			name:     "Valid OrdinaryAccount",
			input:    OrdinaryAccount,
			expected: true,
		},
		{
			name:     "Valid CurrentAccount",
			input:    CurrentAccount,
			expected: true,
		},
		{
			name:     "Invalid AccountKind (zero value)",
			input:    AccountKind(0),
			expected: false,
		},
		{
			name:     "Invalid AccountKind (out of range)",
			input:    AccountKind(3),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.IsValid(); got != tt.expected {
				t.Errorf("AccountKind.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAccountKind_String(t *testing.T) {
	tests := []struct {
		name     string
		input    AccountKind
		expected string
	}{
		{
			name:     "OrdinaryAccount String",
			input:    OrdinaryAccount,
			expected: "普通口座",
		},
		{
			name:     "CurrentAccount String",
			input:    CurrentAccount,
			expected: "当座口座",
		},
		{
			name:     "Invalid AccountKind String",
			input:    AccountKind(3),
			expected: "不明",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.String(); got != tt.expected {
				t.Errorf("AccountKind.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAccountKind_Value(t *testing.T) {
	tests := []struct {
		name     string
		input    AccountKind
		expected int
	}{
		{
			name:     "OrdinaryAccount Value",
			input:    OrdinaryAccount,
			expected: 1,
		},
		{
			name:     "CurrentAccount Value",
			input:    CurrentAccount,
			expected: 2,
		},
		{
			name:     "Invalid AccountKind Value",
			input:    AccountKind(3),
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Value(); got != tt.expected {
				t.Errorf("AccountKind.Value() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBankBranchCode_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    BankBranchCode
		expected bool
	}{
		{
			name:     "Valid 3-digit branch code",
			input:    BankBranchCode("123"),
			expected: true,
		},
		{
			name:     "Invalid 2-digit branch code",
			input:    BankBranchCode("12"),
			expected: false,
		},
		{
			name:     "Invalid 4-digit branch code",
			input:    BankBranchCode("1234"),
			expected: false,
		},
		{
			name:     "Invalid negative branch code",
			input:    BankBranchCode("-123"),
			expected: false,
		},
		{
			name:     "Invalid zero branch code",
			input:    BankBranchCode("0"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.IsValid(); got != tt.expected {
				t.Errorf("BankBranchCode.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBankCode_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    BankCode
		expected bool
	}{
		{
			name:     "Valid 4-digit bank code",
			input:    BankCode("1234"),
			expected: true,
		},
		{
			name:     "Invalid 3-digit bank code",
			input:    BankCode("123"),
			expected: false,
		},
		{
			name:     "Invalid 5-digit bank code",
			input:    BankCode("12345"),
			expected: false,
		},
		{
			name:     "Invalid negative bank code",
			input:    BankCode("-1234"),
			expected: false,
		},
		{
			name:     "Invalid zero bank code",
			input:    BankCode("0"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.IsValid(); got != tt.expected {
				t.Errorf("BankCode.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}
