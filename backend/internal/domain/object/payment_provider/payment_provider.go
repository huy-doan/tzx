package object

type PaymentProvider string

const (
	PaymentProviderPayPay PaymentProvider = "PAYPAY"
)

// String representation of PaymentProvider
func (p PaymentProvider) String() string {
	return string(p)
}

func (p PaymentProvider) IsValid() bool {
	switch p {
	case PaymentProviderPayPay:
		return true
	default:
		return false
	}
}

func (p PaymentProvider) IsImplemented() bool {
	switch p {
	case PaymentProviderPayPay:
		return true
	default:
		return false
	}
}

func GetSupportedProviders() []PaymentProvider {
	return []PaymentProvider{
		PaymentProviderPayPay,
	}
}

func GetSupportedProviderNames() []string {
	providers := GetSupportedProviders()
	names := make([]string, len(providers))
	for i, provider := range providers {
		names[i] = provider.String()
	}
	return names
}

func IsValidProvider(provider string) bool {
	return PaymentProvider(provider).IsValid()
}

// FromString creates a PaymentProvider from a string
func FromString(provider string) PaymentProvider {
	return PaymentProvider(provider)
}
