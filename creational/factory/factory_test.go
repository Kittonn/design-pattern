package factory

import "testing"

func TestPayWithPromptPay(t *testing.T) {
	payment, err := PaymentFactory(PromptPayType)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result := payment.Pay(100.0)
	expected := "Paid 100.00 with PromptPay"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestPayWithCreditCard(t *testing.T) {
	payment, err := PaymentFactory(CreditCardType)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result := payment.Pay(200.0)
	expected := "Paid 200.00 with Credit Card"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestPayWithUnknownType(t *testing.T) {
	_, err := PaymentFactory("unknown")
	if err == nil {
		t.Error("Expected error for unknown payment type, got nil")
	}
}
