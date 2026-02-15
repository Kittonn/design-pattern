package factory

import "fmt"

type PaymentType string

const (
	PromptPayType  PaymentType = "promptpay"
	CreditCardType PaymentType = "creditcard"
)

type Payment interface {
	Pay(amount float64) string
}

type PromptPay struct{}

func (p *PromptPay) Pay(amount float64) string {
	return "Paid " + fmt.Sprintf("%.2f", amount) + " with PromptPay"
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return "Paid " + fmt.Sprintf("%.2f", amount) + " with Credit Card"
}

func PaymentFactory(paymentType PaymentType) (Payment, error) {
	switch paymentType {
	case PromptPayType:
		return new(PromptPay), nil

	case CreditCardType:
		return new(CreditCard), nil

	default:
		return nil, fmt.Errorf("Unknown payment type: %s", paymentType)
	}
}
