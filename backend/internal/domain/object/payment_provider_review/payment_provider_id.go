package object

type PaymentProviderID int

func (m PaymentProviderID) Value() int {
	return int(m)
}
