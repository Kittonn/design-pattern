package abstract_factory

import "testing"

func TestNikeFactory(t *testing.T) {
	factory, err := GetBrandFactory("nike")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	shoe := factory.CreateShoe()
	shirt := factory.CreateShirt()

	if shoe.Wear() != "Wearing Nike shoes" {
		t.Errorf("Unexpected shoe behavior: %s", shoe.Wear())
	}

	if shirt.PutOn() != "Putting on a Nike shirt" {
		t.Errorf("Unexpected shirt behavior: %s", shirt.PutOn())
	}
}

func TestAdidasFactory(t *testing.T) {
	factory, err := GetBrandFactory("adidas")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	shoe := factory.CreateShoe()
	shirt := factory.CreateShirt()

	if shoe.Wear() != "Wearing Adidas shoes" {
		t.Errorf("Unexpected shoe behavior: %s", shoe.Wear())
	}

	if shirt.PutOn() != "Putting on an Adidas shirt" {
		t.Errorf("Unexpected shirt behavior: %s", shirt.PutOn())
	}
}

func TestUnknownFactory(t *testing.T) {
	_, err := GetBrandFactory("unknown")
	if err == nil {
		t.Error("Expected error for unknown brand, got nil")
	}
}
