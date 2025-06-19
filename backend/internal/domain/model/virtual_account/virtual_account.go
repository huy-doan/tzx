package model

import (
	"errors"
	"regexp"
	"strings"

	paymentProviderModel "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

type VirtualAccount struct {
	ID                int
	PaymentProviderID int
	VaAccountNumber   string
	VaAccountName     string
	BranchCode        string
	RawJsonData       *RawJsonData `json:"raw_json_data"`
	PaymentProvider   *paymentProviderModel.PaymentProvider
	util.BaseColumnTimestamp
}

// Business logic methods
func (va *VirtualAccount) IsValid() error {
	if va.VaAccountNumber == "" {
		return errors.New("virtual account number is required")
	}

	if va.VaAccountName == "" {
		return errors.New("virtual account name is required")
	}

	if va.BranchCode == "" {
		return errors.New("branch code is required")
	}

	if va.PaymentProviderID <= 0 {
		return errors.New("valid payment provider ID is required")
	}

	return nil
}

func (va *VirtualAccount) ValidateAccountNumber() error {
	// Remove spaces and convert to uppercase
	cleaned := strings.ToUpper(strings.ReplaceAll(va.VaAccountNumber, " ", ""))

	// Basic format validation - flexible for different providers
	// Allow alphanumeric, 6-30 characters to accommodate various formats
	matched, err := regexp.MatchString(`^[A-Z0-9]{6,30}$`, cleaned)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid account number format: must be 6-30 alphanumeric characters")
	}

	va.VaAccountNumber = cleaned
	return nil
}

func (va *VirtualAccount) ValidateBranchCode() error {
	// Remove spaces and convert to uppercase
	cleaned := strings.ToUpper(strings.ReplaceAll(va.BranchCode, " ", ""))

	// Branch code can be 3-6 characters (digits or alphanumeric)
	matched, err := regexp.MatchString(`^[A-Z0-9]{3,6}$`, cleaned)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("branch code must be 3-6 alphanumeric characters")
	}

	va.BranchCode = cleaned
	return nil
}

func (va *VirtualAccount) CanBeUpdated() bool {
	// Add business rules for when a virtual account can be updated
	return va.ID > 0
}

type PaginatedVirtualAccountResult struct {
	Items []*VirtualAccount
	util.Pagination
}
