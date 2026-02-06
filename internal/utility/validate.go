package utility

type PaymentType string

const (
	PaymentCreditCard PaymentType = "credit_card"
	PaymentQR         PaymentType = "qr"
	PaymentCOD        PaymentType = "cod"
)

func IsValidPaymentType(p PaymentType) bool {
	switch p {
	case PaymentCreditCard, PaymentQR, PaymentCOD:
		return true
	default:
		return false
	}
}
